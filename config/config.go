package config

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ConfingApi() string {
	file, err := os.Open(".env")
	if err != nil {
		log.Fatal("Error opening .env file:", err)
	}
	defer file.Close()

	var port string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "PORT=") {
			port = strings.TrimPrefix(line, "PORT=")
			break
		}
	}

	if port == "" {
		log.Fatal("PORT is not set in the .env file")
	}

	fmt.Println(port)
	return port
}
