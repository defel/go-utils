package utils

import (
	"fmt"
	"os"
	"log"
	"bufio"
)

// naive implementation to iterate over a file with one password each line and check if password from file matches the given password
func CheckPasswordIsUnique(password string, passwordFile string) bool {
	passwordIsUnique := true

	// TODO: duration to scan 2150838 password: 150ms - make this better
	file, err := os.Open(passwordFile)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner1 := bufio.NewScanner(file)
	for scanner1.Scan() {
		text := scanner1.Text()

		if text == password {
			passwordIsUnique = false
		}
	}

	if err := scanner1.Err(); err != nil {
		log.Fatal(err)
	}

	return passwordIsUnique
}