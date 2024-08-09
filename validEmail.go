package helper

import "regexp"

func isValidEmail(email string) bool {
	// regarx untuk menvalidasi email
	const emailRegexPattern = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegexPattern)
	return re.MatchString(email)
}