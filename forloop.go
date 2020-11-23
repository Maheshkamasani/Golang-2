package main
 
import "fmt"
 
func main() {
 
	k := 1
	for ; k <= 5; k++ {
		fmt.Println(k)
	}
 
	k = 1
	for k <= 5 {
		fmt.Println(k)
		k++
	}
 
	for k := 1; ; k++ {
		fmt.Println(k)
		if k == 5 {
			break
		}
	}
}