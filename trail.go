package main
// import "fmt"

// func Reverse(arr []int){
//     left:=0
//     right:=len(arr)-1
//     for left<right{
//         arr[left],arr[right]=arr[right],arr[left]
//         left++
//         right--
//     }
//     fmt.Println(arr)
// }

// Find Largest Element
// func Max(num []int){
//     max:=num[0]
//     min:=num[0]
//     for i:=1;i<len(num);i++{
//         if num[i]>max{
//             max=num[i]
//         }else{
//             min=num[i]
//         }
//     }
//     fmt.Println("Max number :",max)
//     fmt.Println("Min number :",min)
// }

// Reverse a String
// func RS(){
//     r:=[]rune(array)
//     left:=0
//     right:=len(r)-1
//     for left<right{
//         r[left],r[right]=r[right],r[left]
//         left++
//         right--
//     }
//     fmt.Println(string(r))
// }

// Count Vowels in String
// var str="HY MY NAME IS SHABIN"
// func Vowel(){
//         v:=0
//         c:=0
//         for i:=0;i<len(str);i++{
//             ch:=str[i]
//             if ch=='A'||ch=='E'||ch=='I'||ch=='U'||ch=='U'{
//                 v++
//             }else if ch!=' '{
//                 c++
//             }
//         }
//         fmt.Println(v)
//         fmt.Println(c)
// }

// Find Duplicate Elements in Array
// func main(){
//     array:=[]int{1,2,3,4,5,6,7,7,1,2}
//     m:=make(map[int]int)
//     dupe:=[]int{}
//     for _,n:=range array{
//         m[n]++
//         if m[n]==2{
//             dupe=append(dupe,n)
//         }
//     }
//     fmt.Println(dupe)
// }
