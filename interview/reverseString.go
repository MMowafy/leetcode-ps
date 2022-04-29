package main

func main() {
	reverseString([]byte{'h', 'e', 'l', 'l', 'o'})
}

func reverseString(s []byte) {
	j := len(s) - 1
	for i := 0; i < len(s); i++ {
		if j == i || j < i {
			break
		}
		s[i], s[j] = s[j], s[i]
		j--
	}
}
