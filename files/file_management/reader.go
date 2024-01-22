package file_management

import (
	"fmt"
	"os"
)

func ReadFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("error reading file")
		return "", err
	}
	return string(content), nil
}
