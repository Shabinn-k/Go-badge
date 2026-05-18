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

// func main(){
// arr:=[]int{1,2,3,4,5}
// l:=0
// f:=len(arr)-1
// for l<f{
// 	arr[l],arr[f]=arr[f],arr[l]
// 	l++
// 	f--
// }
// fmt.Println(arr)
// }

// type Shape interface{
// 	Circle() int
// 	Rectangle() int
// }

// type Area struct{
// 	Length int
// 	Width int
// }

// func (a Area)Circle()int{
// 	return 2*3*a.Length
// }
// func (a Area)Rectangle()int{
// 	return a.Length*a.Width
// }

// func Trial(s Shape){
// 	fmt.Println(s.Circle())
// 	fmt.Println(s.Rectangle())
// }

// func  main(){
// 	aa:=Area{
// 		Width: 90,
// 		Length: 10,
// 	}
// 	Trial(aa)
// }

// func main(){
// 	str:="Shabin is a good boy"
// 	m:=make(map[rune]int)
// 	for _,v:=range str{
// 		m[v]++
// 	}
// 	for a,b:=range m{
// 		fmt.Printf("%c: %d\n",a,b)
// 	}
// }

// func main(){
// 	str:="Shabin is a good boy"
// 	m:=make(map[int]string)
// 	for i:=0;i<len(str);i++{
// 		m[i]=string(str[i])
// 	}
// 	fmt.Println(m)
// }

// func main(){
// 	str:="Shabin is a good boy "
// 	a:=strings.Fields(str)
// 	m:=make(map[string]int)
// 	for _,v:=range a{
// 		m[v]++
// 	}
// 	fmt.Println(m)
// }

// var mu sync.Mutex
// var bank int

// func Deposits(wg *sync.WaitGroup){
// 	defer wg.Done()
// 	mu.Lock()
// 	bank+=10
// 	mu.Unlock()
// }

// func main(){
// 	var wg sync.WaitGroup
// 	for i:=0;i<10000;i++{
// 		wg.Add(1)
// 		go Deposits(&wg)
// 	}
// 	wg.Wait()
// 	fmt.Println(bank)
// }

// func main(){
// target:=9
// arr:=[]int{2,4,5,6,8}
// m:=make(map[int]int)

// for i,v:=range arr{
// 	get:=target-v
// 	if ind,fou:=m[get];fou{
// 		fmt.Println(ind,i)
// 		return
// 	}
// 	m[v]=i
// }
// fmt.Println("Nothing")
// }

// func Odds(oddch chan int){
// 	for i:=1;i<10;i+=2{
// 		oddch <-i
// 	}
// close(oddch)
// }
// func Evens(evench chan int){
// for i:=0;i<10;i+=2{
// 	evench<-i
// }
// 	close(evench)
// }

// func main(){
// 	oddch:=make(chan int)
// 	evench:=make(chan int)

// 	go Odds(oddch)
// 	go Evens(evench)

// 	for n:=range evench{
// 		fmt.Println(n)
// 	}
// 	for n:=range oddch{
// 		fmt.Println(n)
// 	}
// }
// func main(){
// str:="madam"
// f:=0
// p:=false
// l:=len(str)-1
// for f>l{
// 	if str[f]!=str[l]{
// 		p=true
// 		break
// 	}
// 	f++
// 	l--
// }
// 	fmt.Println(p)
// }
