package internal

import (
	"fmt"
	"io"
	"log"
	"os"
)

func GetInput(day int) string {
	f, err := os.Open(fmt.Sprintf("files/day%d/input.txt", day))
	if err != nil {
		log.Fatal(err)
	}

	content, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func GetTest(day int) string {
	f, err := os.Open(fmt.Sprintf("files/day%d/test.txt", day))
	if err != nil {
		log.Fatal(err)
	}

	content, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}
