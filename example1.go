package main
import "fmt"

func main (){
	// BAHOLASH TIZIMI
	var ball int
	fmt.Println("Ballni kiriting ! ")
	fmt.Scanf("%d", &ball)
	if ball>100 {
		fmt.Println("Grandd")
	}else if ball>=80 && 100>=ball {
		fmt.Println("5 baho")
	}else if ball>=60 && 80>ball{
		fmt.Println("4 baho")
	}else if ball>=40 && 60>ball {
		fmt.Println("3 baho")
	}else{
		fmt.Println("Yiqildingiz")
	}
}