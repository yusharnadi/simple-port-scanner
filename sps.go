package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage : go run sps.go target.com 1 80")
		os.Exit(1)
	}
	hostname := os.Args[1]
	start, _ := strconv.Atoi(os.Args[2])
	end, _ := strconv.Atoi(os.Args[3])

	var ports uint16
	var wg sync.WaitGroup
	for i := start; i < end; i++ {
		wg.Add(1)
		go func(i int) {
			target := fmt.Sprintf(hostname+":%d", i)
			_, err := net.Dial("tcp", target)
			if err != nil {
				// continue
			} else {
				fmt.Println(i, "is Open.")
				ports += 1
			}
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Println("==============================")
	fmt.Println(ports, "ports open.")
	fmt.Println("==============================")
}
