package two_pointers

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	l, r := 0, len(s)-1

	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
