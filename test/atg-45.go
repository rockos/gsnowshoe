package main

import (
	"fmt"
	"log"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	log.Println("Responsing to /hello request")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}
