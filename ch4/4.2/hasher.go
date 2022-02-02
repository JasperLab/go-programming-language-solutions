package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		t := input.Text()
		formats := os.Args[1:]
		if len(formats) == 0 {
			fmt.Printf("%v\n", sum256(t))
		} else {
			switch formats[0] {
			case "384":
				fmt.Printf("%v\n", sum384(t))
			case "512":
				fmt.Printf("%v\n", sum512(t))
			default:
				fmt.Printf("%v\n", sum256(t))
			}
		}

	}
}

func sum256(s string) [32]byte {
	return sha256.Sum256([]byte(s))
}

func sum384(s string) [48]byte {
	return sha512.Sum384([]byte(s))
}

func sum512(s string) [64]byte {
	return sha512.Sum512([]byte(s))
}
