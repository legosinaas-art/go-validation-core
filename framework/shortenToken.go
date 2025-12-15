package framework

func ShortenToken(token string) string {
	const minLen = 50    // Начинаем сокращать, если токен длиннее 50
	const prefixLen = 20 // Берем первые 20 символов
	const suffixLen = 15 // Берем последние 15 символов

	if len(token) < minLen {
		return token
	}

	// Берем начало и конец, вставляем троеточие
	start := token[:prefixLen]
	end := token[len(token)-suffixLen:]

	return start + "..." + end
}
