package rebase64

const encodeStd = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
const StdPadding = '='

var StdEncoding = NewEncoding(encodeStd)

type Encoding struct {
	encode  [64]byte
	padChar rune
}

func NewEncoding(encoder string) *Encoding {
	if len(encoder) != 64 {
		panic("encoder strings length is not 64")
	}
	for i := 0; i < len(encoder); i++ {
		if encoder[i] == '\n' || encoder[i] == '\r' {
			panic("encoder strings has newline character")
		}
	}

	e := new(Encoding)
	e.padChar = StdPadding
	copy(e.encode[:], encoder)
	return e
}

func (enc *Encoding) Encode(dst, src []byte) {
	if len(src) == 0 {
		return
	}

	// さきに nil checkすることで、以降のfor内での nil checkが発生せず、
	// パフォーマンスがあがるらしい
	_ = enc.encode

	di, si := 0, 0
	n := (len(src) / 3) * 3

	for si < n {
		// 参考: https://qiita.com/PlanetMeron/items/2905e2d0aa7fe46a36d4
		//    : https://go.dev/src/encoding/base64/base64.go?s=3725:3769#L140
		//
		// もし間違いなどありましたら https://twitter.com/mahiro0x00 までご連絡いただけますと幸いです。
		//
		// 1. 変換したい文字列 (例:abcdefg の場合)
		// →0110 0001, 0110 0010, 0110 0011, 0110 0100, 0110  0101, 0110 0110, 0110 0111(2進数)(1文字あたり8bit)
		//
		// 2. バイナリを6bitずつに変換
		// 011000, 010110, 001001, 100011, 011001, 000110, 010101, 100110, 011001, 11
		// memo: この時点でもともと3文字だった場合、変換後は4文字になる
		//     : 8bit * 3文字分 (変換前) = 6bit * 4文字分 (変換後)
		//
		// 3. 最後の2ビットが余るので,6ビットになるように0を追加する
		// 011000, 010110, 001001, 100011, 011001, 000110, 010101, 100110, 011001, 110000
		//
		// 4. {encodeStd}よりビットを文字に変換する
		//
		// 5. 4文字に分けた時に、足りない分{paddingStr}を追加する

		// "0110 0001", "0110 0010", "0110 0011" の3つの文字について最後の文字以外シフトする
		// シフトし、論理和をとった値(下記のval)は以下のようになる
		// "0110 0001 0110 0010 0110 0011"
		val := uint(src[si+0])<<16 | uint(src[si+1])<<8 | uint(src[si+2])

		// 上記のvalを6bitずつに変換し、4文字に分ける
		// それぞれの具体的な値については以下である

		// "0000 0000 0000 0000 0001 1000"
		// 論理積をとると "011000"
		dst[di+0] = enc.encode[val>>18&0x3F]

		// "0000 0000 0000 0110 0001 0110"
		// 論理積をとると "010110"
		dst[di+1] = enc.encode[val>>12&0x3F]

		// "0000 0001 1000 0101 1000 1001"
		// 論理積をとると "001001"
		dst[di+2] = enc.encode[val>>6&0x3F]

		// "0110 0001 0110 0010 0110 0011"
		// 論理積をとると "100011"
		dst[di+3] = enc.encode[val&0x3F]

		si += 3
		di += 4
	}

	remain := len(src) - si
	if remain == 0 {
		return
	}

	val := uint(src[si+0]) << 16
	if remain == 2 {
		val |= uint(src[si+1]) << 8
	}
	dst[di+0] = enc.encode[val>>18&0x3F]
	dst[di+1] = enc.encode[val>>12&0x3F]

	switch remain {
	case 2:
		dst[di+2] = enc.encode[val>>6&0x3F]
		dst[di+3] = byte(enc.padChar)
	case 1:
		// 最後に{stdPadding}が2文字追加される
		dst[di+2], dst[di+3] = byte(enc.padChar), byte(enc.padChar)
	}
}

func (enc *Encoding) EncodeToString(src []byte) string {
	buf := make([]byte, enc.EncodedLen(len(src)))
	enc.Encode(buf, src)
	return string(buf)
}

func (enc *Encoding) EncodedLen(n int) int {
	// example:
	// ..., 6　->8
	// 7->12(min), 8->12, 9->12
	// 10->16(min), 11->16, ...
	return (n + 2) / 3 * 4
}

// 分かってないことメモ
// - なぜuintでキャストするのか
