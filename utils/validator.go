package utils

import "net/mail"

func IsEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func IsAgeValid(age int) bool {
	return (1 <= age && age <= 120)
}

func IsGpaValid(gpa float64) bool {
	return (0.0 <= gpa && gpa <= 4.0)
}
