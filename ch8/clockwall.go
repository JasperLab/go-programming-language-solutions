package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	var times []string
	for i, arg := range os.Args[1:] {
		a := strings.Split(arg, "=")
		city := a[0]
		addr := a[1]
		times = append(times, "")
		fmt.Printf("%s\t\t", city)
		go readTime(addr, &times[i])
	}
	fmt.Print("\n")

	for {
		printTime(times)
		time.Sleep(1 * time.Second)
	}

}

func readTime(addr string, val *string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		*val = scanner.Text()
	}
}

func printTime(times []string) {
	carr := len(times)*8 + (len(times)-1)*4
	for i := 0; i < carr; i++ {
		fmt.Print("\r")
	}
	for i, t := range times {
		if i > 0 {
			fmt.Print("    ")
		}
		fmt.Print(t)
	}
}
