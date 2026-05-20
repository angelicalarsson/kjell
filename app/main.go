package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("$ ")
	for scanner.Scan() {

		err := scanner.Err()

		if err != nil {
			log.Print(err)
		}

		fmt.Printf("%s: command not found\n", scanner.Text())
		fmt.Print("$ ")

	}
}
