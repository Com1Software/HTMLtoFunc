package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Printf("HTMLtoFunc Convert HTML Page to Go Function \n")
	fmt.Printf("Operating System : %s\n", runtime.GOOS)

	switch {
	//----------------------------
	case len(os.Args) == 2:
		cmd := os.Args[1]
		switch {
		case cmd == "overview":
			fmt.Printf("Overview\n")

		default:

		}
	case len(os.Args) == 1:
	}

}
