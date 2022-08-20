package main

import "fmt"

func main() {
	//SWITCH
	fmt.Println("1. balans\n2. Sms\n3. mb")
	var tanlov int
	fmt.Println("tanlang! ")
	fmt.Scanf("%d", &tanlov)
	switch tanlov {
	case 1:
		fmt.Println("balansda : 34$ mavjud")
		break

	case 2:
		fmt.Println("balansda : 235sms mavjud")
		break

	case 3:
		fmt.Println("balansda : 34132mb mavjud")
		break
	default:
		fmt.Println("notori")
	}
}
