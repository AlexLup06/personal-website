package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LoadEnv() map[string]string {
	file, err := os.OpenFile(".env", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Printf("Error opneing .env: %v\n", err)
	}
	defer file.Close()

	envVars := map[string]string{}

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()
		envPair := strings.Split(line, "=")
		envVars[envPair[0]] = envPair[1]
	}
	return envVars
}
