package rebase64

import (
	"testing"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		name       string
		src        []byte
		encodedSrc string
	}{
		{
			"1文字のとき",
			[]byte("h"),
			"aA==",
		},
		{
			"2文字のとき",
			[]byte("ho"),
			"aG8=",
		},
		{
			"3文字のとき",
			[]byte("hog"),
			"aG9n",
		},
		{
			"4文字のとき",
			[]byte("hoge"),
			"aG9nZQ==",
		},
		{
			"5文字のとき",
			[]byte("hogeh"),
			"aG9nZWg=",
		},
		{
			"10文字のとき",
			[]byte("hogehogeho"),
			"aG9nZWhvZ2Vobw==",
		},
		{
			"50文字のとき",
			[]byte("HFX8qtTxDgJWyrg4ckeJwdFVQl3dprps3N5YSsmneO4FY7tGlS"),
			"SEZYOHF0VHhEZ0pXeXJnNGNrZUp3ZEZWUWwzZHBycHMzTjVZU3NtbmVPNEZZN3RHbFM=",
		},
		{
			"250文字のとき",
			[]byte("UubBo6mTajg4tINgpYMm6fy5qKi8taaWl0KBEKwTQeNs07cSPFaLGcBw8w1XQWsVLnWwxIhxgRH48Ahmui7n5SOivsC81QdY4KTqRt2c1a2NaYNh1jDu4wJCeRfHghwcWkpJi7MiSF8yUMSo305VHGwCqlc73xsQtPJ8OW0h3eODJyv0sF4huAT8NCv75AkdTRd7O2bcNxLFu8ki0NagWv08PMNlGv1milcRkaei7fnUG8eoiq8apehagr"),
			"VXViQm82bVRhamc0dElOZ3BZTW02Znk1cUtpOHRhYVdsMEtCRUt3VFFlTnMwN2NTUEZhTEdjQnc4dzFYUVdzVkxuV3d4SWh4Z1JINDhBaG11aTduNVNPaXZzQzgxUWRZNEtUcVJ0MmMxYTJOYVlOaDFqRHU0d0pDZVJmSGdod2NXa3BKaTdNaVNGOHlVTVNvMzA1VkhHd0NxbGM3M3hzUXRQSjhPVzBoM2VPREp5djBzRjRodUFUOE5Ddjc1QWtkVFJkN08yYmNOeExGdThraTBOYWdXdjA4UE1ObEd2MW1pbGNSa2FlaTdmblVHOGVvaXE4YXBlaGFncg==",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StdEncoding.EncodeToString(tt.src); got != tt.encodedSrc {
				t.Fatalf("expect %s, but actual %s", tt.encodedSrc, got)
			}
		})
	}
}
