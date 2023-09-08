package dictionary

import (
	"strings"
	"unicode"
)

func RatePassword(password string) int {
	var rating int

	if len(password) < 8 {
		return 0
	}

	// Length-based rating
	lengthRating := len(password) * 7
	if lengthRating > 100 {
		lengthRating = 100
	}
	rating += lengthRating / 4

	// Letter case rating
	upperCount, lowerCount := countLetterCase(password)
	letterCaseRating := 0
	if upperCount > 0 && lowerCount > 0 {
		letterCaseRating = 20
	} else if upperCount > 0 || lowerCount > 0 {
		letterCaseRating = 10
	}
	rating += letterCaseRating

	// Digit rating
	digitCount := countDigits(password)
	digitRating := 0
	if digitCount > 0 && digitCount < len(password) {
		digitRating = 20
	} else if digitCount > 0 {
		digitRating = 10
	}
	rating += digitRating

	// Symbol rating
	symbolCount := countSymbols(password)
	symbolRating := 0
	if symbolCount > 0 {
		symbolRating = 20
	}
	rating += symbolRating

	// Bonus rating for additional criteria
	bonusRating := 0
	if lowerCount > 0 && upperCount > 0 && digitCount > 0 && symbolCount > 0 {
		bonusRating = 20
	}
	rating += bonusRating

	return rating
}

func countLetterCase(password string) (int, int) {
	upperCount := 0
	lowerCount := 0
	for _, ch := range password {
		if unicode.IsUpper(ch) {
			upperCount++
		} else if unicode.IsLower(ch) {
			lowerCount++
		}
	}
	return upperCount, lowerCount
}

func countDigits(password string) int {
	count := 0
	for _, ch := range password {
		if unicode.IsDigit(ch) {
			count++
		}
	}
	return count
}

func countSymbols(password string) int {
	count := 0
	for _, ch := range password {
		if !unicode.IsLetter(ch) && !unicode.IsDigit(ch) {
			count++
		}
	}
	return count
}

const valid = "0123456789qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM!@#$%^&*()_-+={[}]|\\:;<,>.?/"

func Validate(pass string) bool {
	for _, char := range pass {
		if !strings.Contains(valid, string(char)) {
			return false
		}
	}
	return true
}
