package naming

func ToSnake(s string) string {
	const k = 'a' - 'A'
	snake := make([]rune, 0, len(s)+1)
	flag := false
	for i, c := range s {
		if c >= 'A' && c <= 'Z' {
			if !flag {
				flag = true
				if i > 0 {
					snake = append(snake, '_')
				}
			}
			snake = append(snake, c+k)
		} else {
			flag = false
			snake = append(snake, c)
		}
	}
	return string(snake)
}
