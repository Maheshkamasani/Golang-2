package main
 
import "fmt"
 
func main() {
 
	numbers :=[]int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
        sum :=0
var i=20
for i=0;i<len(numbers);i++ {
    sum+=numbers[i]
  }
fmt.Println("Sum is ::",sum)
}