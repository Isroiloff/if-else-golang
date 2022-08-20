package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("a sonini kiriting ")
	fmt.Scanln(&a)
	fmt.Println("b sonini kiriting ")
	fmt.Scanln(&b)

	if a > b {
		fmt.Println("a soni katta")
	} else if a == b {
		fmt.Println("a=b ga")
	} else {
		fmt.Println("b soni katta")
	}

	if son := 34; son > 0 {
		fmt.Println("pozitiv")
	} else {
		fmt.Println("nefativ")
	}

}
