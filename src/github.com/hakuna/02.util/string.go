package _2_util

func Reverse(value string) string {
	temp := []rune(value)
	for i, j := 0, len(temp)-1; i < len(temp)/2; i, j = i+1, j-1 {
		temp[i], temp[j] = temp[j], temp[i]
	}

	return string(temp)
}
