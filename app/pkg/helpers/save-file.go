package helpers

import "os"

func SaveHash(data string) (bool, error) {
	file, err := os.Create("Verify.txt")

	defer file.Close()

	if err != nil {
		return false, err
	}

	file.WriteString(data)

	return true, nil
}

func ReadHash() (string, error) {
	content, err := os.ReadFile("Verify.txt")

	if err != nil {
		return "", err
	}

	return string(content), nil
}

func DelereHash() error {
	return os.Remove("Verify.txt")
}
