package camel

import (
	"bytes"
	"strings"
)

// Marshal Marshal
func Marshal(name string, firstUpper bool) string {
	if name == "" {
		return ""
	}
	temp := strings.Split(name, "_")
	var s string
	for _, v := range temp {

		vv := []rune(v)
		if len(vv) > 0 {
			if bool(vv[0] >= 'a' && vv[0] <= 'z') { //首字母大写
				vv[0] -= 32
			}
			s += string(vv)
		}
	}
	if !firstUpper {
		sv := []rune(s)
		sv[0] += 32
		s = string(sv)
	}
	return s
}

// UnMarshal UnMarshal
func UnMarshal(name string) string {
	if name == "" {
		return ""
	}

	buf := bytes.NewBufferString("")
	sv := []rune(name)
	sl := len(sv)
	for i := 0; i < sl-1; i++ {
		buf.WriteRune(sv[i])
		if sv[i+1] >= 'A' && sv[i+1] <= 'Z' {
			buf.WriteRune('_')
		}
	}
	buf.WriteRune(sv[sl-1])
	s := strings.ToLower(buf.String())
	return s
}
