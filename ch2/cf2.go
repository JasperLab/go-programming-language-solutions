package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"gopl/ch2/tempconv"
)

func main() {
	args := os.Args[1:]

	if len(args) > 0 {
		for _, a := range(args) {
			v, err := strconv.ParseFloat(a, 64)
			if err != nil {
				log.Println("Invalid argument " + a)
				continue
			}
			print(v)
		}
	} else {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			 a := input.Text()
			v, err := strconv.ParseFloat(a, 64)
			if err != nil { 
				log.Println("Invalid input " + a)
				continue
			}
			print(v)
		}
	}
}

func print(val float64) {
	fmt.Printf("%g°C = %g°F\t%g°F = %g°C\n", val, tempconv.CToF(tempconv.Celsius(val)), val, tempconv.FToC(tempconv.Fahrenheit(val)))
}
