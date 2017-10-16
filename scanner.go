package main

import (
	"bufio"
	"fmt"
	"os"
)

func serveScan() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "close" {
			os.Exit(1)
		}
		fmt.Println(scanner.Text())
	}
}
