package main

import (
	"fmt"
	"time"
)

func task(name string) {
	fmt.Println("Starting task:", name)
	time.Sleep(1 * time.Second)
	fmt.Println("Finished task:", name)
	fmt.Println("Concurrency")
}

func main() {
	go task("Alpha")
	go task("Beta")
	go task("Charlie")

	time.Sleep(3 * time.Second)
	fmt.Println("Main function completed.")
}
