package main

import (
	"files/file_management"
	"fmt"
	"os"
)

func main() {
	rootPath, err := os.Getwd()
	filePath := rootPath + "/data/text.txt"
	if err != nil {
		fmt.Println("Did not work")
	}

	con, err := file_management.ReadFile(filePath)

	if err != nil {
		fmt.Println("did not work")
	}
	fmt.Printf(con)
	newContent := fmt.Sprintf("Original:\n%v\nDouble Original:\n%v%v", con, con, con)
	file_management.WriteToFile(filePath, newContent)

}
