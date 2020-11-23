package main

import "fmt"

func main() {

    nums := []int{1,2,3,4,5,6,7,8,9,10}
    sum := 0
    for index, val := range nums {
        sum += val
    
    fmt.Print("[",index,",",val,"]")
     }
fmt.Println("\nSum is ::",sum)
kvs:=map[int]string{1:"divya",2:"chandu"}
for k,v :=range kvs{
fmt.Println(k, "->",v)
}
str :="hello,world!"
for index,c :=range str{
fmt.Print("[",index,",",string(c),"]")
}
}
    

    