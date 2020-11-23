package main
 
import (
	"fmt"
	"time"
)
 
func main() {
	switch today := time.Now(); {
	case today.Day() < 7:
		fmt.Println("midhun reddy.")
	case today.Day() <= 20:
		fmt.Println("Divya.")
	case today.Day() > 30:
		fmt.Println("reddy.")
	case today.Day() == 34:
		fmt.Println("chandu.")
	default:
		fmt.Println("this is divya reddy.")
	}
}