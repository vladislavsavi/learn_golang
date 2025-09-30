package utils

import (
	"crypto/rand"
	"errors"
	"math/big"
	"strings"
)

const (
	MinPasswordLength = 4
	MinPasswordsCount = 1
	MaxPasswordsCount = 50
)

var (
	ErrPasswordLengthTooLow = errors.New("password length too low")
	ErrTooLowPasswordsCount = errors.New("too low passwords count")
	ErrTooBigPasswordsCount = errors.New("too big passwords count")
)

var (
	upperChars   = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	lowerChars   = []rune("abcdefghijklmnopqrstuvwxyz")
	digitChars   = []rune("0123456789")
	specialChars = []rune("!@#$%^&*")
)

func validateArgs(length int, count int) error {
	if length < MinPasswordLength {
		return ErrPasswordLengthTooLow
	}

	if count > MaxPasswordsCount {
		return ErrTooBigPasswordsCount
	}

	if count < MinPasswordsCount {
		return ErrTooLowPasswordsCount
	}

	return nil
}

func getRandomNum(maxVal int64) int {
	if randomNum, err := rand.Int(rand.Reader, big.NewInt(maxVal)); err != nil {
		return 0
	} else {
		return int(randomNum.Int64())
	}
}

func createOnePass(length int) string {
	pass := make([]string, length)

	chars := [][]rune{upperChars, lowerChars, digitChars, specialChars}
	checker := make(map[int]struct{}, 4)

	passIndex := 0
	for passIndex < length {
		selectedCharScice := getRandomNum(4)
		selectedRuneIndex := getRandomNum(int64(len(chars[selectedCharScice])))

		if _, ok := checker[selectedCharScice]; ok && passIndex+1 <= 4 {
			continue
		} else {
			checker[selectedCharScice] = struct{}{}
			pass[passIndex] = string(chars[selectedCharScice][selectedRuneIndex])
			passIndex++
		}
	}

	return strings.Join(pass, "")
}

func GeneratePassword(length int, count int) ([]string, error) {
	if err := validateArgs(length, count); err != nil {
		return nil, err
	}

	passwords := make([]string, count)

	for index := range passwords {
		passwords[index] = createOnePass(length)
	}

	return passwords, nil
}
