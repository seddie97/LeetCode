package main

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	sK := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			sK = append(sK, s[i])

		} else {
			if len(sK) == 0 {
				return false
			}
			if s[i] == ')' && sK[len(sK)-1] == '(' {
				sK = sK[:len(sK)-1]
				continue
			}
			if s[i] == ']' && sK[len(sK)-1] == '[' {
				sK = sK[:len(sK)-1]
				continue
			}
			if s[i] == '}' && sK[len(sK)-1] == '{' {
				sK = sK[:len(sK)-1]
				continue
			}
			return false
		}
	}

	return len(sK) == 0
}
