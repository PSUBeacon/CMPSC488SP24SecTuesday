package messaging

import (
	"log"
	"os"
)

func main() {
	filePath := "chain.json"

	// Open the file in write-only mode with the truncate option
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// No need to write anything, as truncating the file clears its contents
}
