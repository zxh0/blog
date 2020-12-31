package utils

import (
	"encoding/hex"
	"encoding/json"
)

func ToPrettyJSON(v interface{}) string {
	bytes, _ := json.MarshalIndent(v, "", "  ")
	return string(bytes)
}

func CommonPrefixLen(s1, s2 string) int {
	n1, n2 := len(s1), len(s2)
	i := 0
	for ; i < n1 && i < n2; i++ {
		if s1[i] != s2[i] {
			break
		}
	}
	return i
}

func ToHex(s string) string {
	return hex.EncodeToString([]byte(s))
}
func FromHex(s string) string {
	bytes, _ := hex.DecodeString(s)
	return string(bytes)
}

func ToBin(s string) string {
	bs := make([]byte, 0, len(s)*8)
	for _, b := range s {
		for i := 0; i < 8; i++ {
			if (b<<i)&0x80 > 0 {
				bs = append(bs, byte('1'))
			} else {
				bs = append(bs, byte('0'))
			}
		}
	}
	return string(bs)
}

func FromBin(bs string) string {
	s := make([]byte, len(bs)/8)
	for i := 0; i < len(bs); i += 8 {
		var b byte
		for j := 0; j < 8; j++ {
			b <<= 1
			if bs[i+j] == '1' {
				b |= 1
			}
		}
		s[i/8] = b
	}
	return string(s)
}
