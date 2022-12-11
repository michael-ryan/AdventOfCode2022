package common

import (
	"os"
	"strings"
)

func ReadFileAsStringArray(path string) []string {
	data, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return strings.Split(string(data), "\n")
}

func Stuff(data string) []string {
	return strings.Split(string(data), "\n")
}
