package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFromFile(filename string) ([]string, error) {
	var lines []string
	file, err := os.Open(filename)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	scanner.Err()

	return lines, nil
}

func main() {
	fmt.Println(ReadFromFile("filename.txt"))
}
