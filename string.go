package main
import (
"fmt"
)

func main(){
s :="hello,world!"
r :=[]rune(s)
r(0) ='h'
s2 :=string(s)
fmt.Println(s2)
}