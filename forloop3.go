package main
 
import "fmt"
 
func main() {
 
	numbers :=[]int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
        sum :=0
       i:=20 
       var n=10
for i<n{

    sum+=numbers[i]
    i++
  }
fmt.Println("Sum is ::",sum)
}