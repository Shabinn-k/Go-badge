package main

import "fmt"

// func main() {
// 1
// result1 := TwoSum(nums, target)
// fmt.Println(result1)

// 2
// result2 := Reverse(rev)
// fmt.Println(result2)

// 3
// result3,result4:=Max(maxNums)
// fmt.Println(result3,result4)

// 4
// result5:=Sum(Array1)
// fmt.Println(result5)

// 5
// Mix()

// 6
// head:=Linked()
// Print(head)
// Middle(head)
// reverseHead:=ReverseLL(head)
// Print(reverseHead)

// 7
// array:=[]int{1,2,3,4,5,2,3}
// value:=2
// for i,v:=range array{
// 	if v==value{
// 		array = append(array[:i],array[i+1:]...)
// 	break
// 	}
// }
// array = append(array[:ind],array[ind+1:]...)
// 	dupe:=make(map[int]bool)
// 	res:=[]int{}
// 	for _,v:=range array{
// 		if !dupe[v]{
// 			dupe[v]=true
// 			res = append(res, v)
// 		}
// 	}
// 	fmt.Println(res)

// 8
// Stack()
// Queue()

// 9
// DoubleLL()

// 10
// 	Race()
// }

// var str = "madamk"

// func main() {
// 	l := 0
// 	r := len(str) - 1
// 	g := []rune(str)
// 	for l < r {
// 		if g[r] != g[l] {
// 			fmt.Println(false)
// 			return
// 		}
// 		l++
// 		r--
// 	}
// 	fmt.Println(true)
// }

// func main(){
// 	arr:=[]int{1,2,3,4,0,5}
// 	max:=arr[0]
// 	min:=arr[0]

// 	for i:=1;i<len(arr);i++{
// 		if arr[i]>max{
// 			max=arr[i]
// 		}else{
// 			min=arr[i]
// 		}
// 	}
// 	fmt.Println("max",max)
// 	fmt.Println("Min",min)
// }

// func main(){
// 	arr := []int{10, 5, 8, 20, 15}
// 	large:=arr[0]
// 	second:=arr[0]

// 	for i:=1;i<len(arr);i++{
// 		if arr[i]>large{
// 			second=large
// 			large=arr[i]
// 		}else if arr[i]>second&&arr[i]!=large{
// 			second=arr[i]
// 		}
// 	}
// 	fmt.Println(second)
// }


// var arr=[]int{1,2,3,4,5}
// var targets=9

// func main(){

// seen:=make(map[int]int)
// 	for i,n:=range arr{
// 		get:=targets-n
// 		if ind,found:=seen[get];found{
// 			fmt.Println(ind,i)
// 			return
// 		}
// 		seen[n]=i
// 	}
// 	fmt.Println("nothing")
// }

func main(){
arr:=[]int{1,2,3,4,5}
l:=0
f:=len(arr)-1
for l<f{
	arr[l],arr[f]=arr[f],arr[l]
	l++
	f--	
}
fmt.Println(arr)
}
