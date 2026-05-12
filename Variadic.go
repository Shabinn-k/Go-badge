package main

// import "fmt"

func Variadic(a ...int)int{
    sum:=0
    for _,num:=range a{
        sum+=num
    }
    return sum
}

// func main(){
//     fmt.Println(Variadic(10,1010))
// }