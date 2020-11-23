package main
 
import (
	"fmt"
	"time"
)
 
func main() {
	today := time.Now()
 
	switch today.Day() {
	case 1:
		fmt.Println("Today is 1th. monday.")
	case 2:
		fmt.Println("Today is 2th. thusday.")
	case 3:
		fmt.Println("Today is 3th. friday.")
	case 4:
		fmt.Println("Today is 4th. sunday.")
	case 5:
		fmt.Println("birthaday tonight.")
	default:
		fmt.Println("this is divya.")
	}
}