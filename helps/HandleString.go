package helps

import (
	"strings"
)

func findString(s, stringFind string, stringCount int) bool {
	if s == "" {
		return false
	}
	for i := 0; i < stringCount; i++ {
		if string(s[i]) == stringFind {
			return true
		}
	}
	return false
}

func HandleString(s string) (string, string) {
	var cmnd, key string
	a := []string{}
	b := []string{}
	s2 := strings.Split(s, " ")
	if s != "" {

		a = append(a, s2[0])
		b = append(b, s2[1])

		cmnd = strings.Join(a, "")
		key = strings.Join(b, "")

		LenOfcmnd := strings.Count(cmnd, "") - 1
		LenOfkey := strings.Count(key, "") - 1
		if findString(cmnd, " ", LenOfcmnd) == true {
			a1 := strings.Split(cmnd, " ")
			cmnd = strings.Join(a1, "")
		}
		if findString(key, " ", LenOfkey) == true {
			b1 := strings.Split(key, " ")
			key = strings.Join(b1, "")
		}
	}

	return cmnd, key
}
