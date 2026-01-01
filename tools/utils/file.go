package utils

import "os"

func GetCipherText(filePath string) (string, error) {
	contentBytes, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(contentBytes), nil
}