package main

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/trapedev/go-btpcap/lib"
)

func main() {
	data, err := lib.ReadFileAsByteArray("ubertooth.log")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "AdvA") {
			line = strings.ReplaceAll(line, "AdvA:", "")
			line = strings.TrimSpace(line)
			fmt.Println(line)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error while scanning: %v\n", err)
	}
}