package main

import (
	"fmt"
	// "sync"
)

// Print numbers from 1–100
// func main(){
// 	for i:=1;i<=100;i++{
// 		fmt.Println(i)

// 	}
// }

// Print even numbers from 1–100
// func main(){
// 	for i:=1;i<=100;i+=2{
// 		fmt.Println(i)
// 	}
// }

// Reverse a string
// func main(){
// str:="Hello World"
// r:=[]rune(str)
// 	f:=0
// 	l:=len(r)-1
// 	for f<l{
// 		r[f],r[l]=r[l],r[f]
// 		f++
// 		l--
// 	}
// 	fmt.Println(string(r))
// }

// Check palindrome string
// func main(){
// 	str:="madam"
// 	p:=true
// 	f:=0
// 	l:=len(str)-1
// 	for f<l{
// 		if str[f]!=str[l]{
// 			p=false
// 		}
// 		f++
// 		l--
// 	}
// 	fmt.Println(p)
// }

// Find factorial using recursion
// func Recursion(n int)int{
// 	if n<=1{
// 		return 1
// 	}
// 	return n*Recursion(n-1)
// }
// func main(){
// 	answer:=Recursion(10)
// 	fmt.Println(answer)
// }

// Find factorial using loop
// func main(){
// 	n:=10
// 	fact:=1
// 	if n<=1{
// 		fmt.Println(1)
// 		return
// 	}
// 	for i:=1;i<=n;i++{
// 		fact=fact*i
// 	}
// 	fmt.Println(fact)
// }

// Swap two numbers without third variable
// func main(){
// 	a:=432
// 	b:=6423
// 	fmt.Println("Before Swap")
// 	fmt.Println(a)
// 	fmt.Println(b)

// 	a=a+b
// 	b=a-b
// 	a=a-b
// 	fmt.Println("After Swap")
// 	fmt.Println(a)
// 	fmt.Println(b)
// }

// Find largest element in array
// func main(){
// 	array:=[]int{1,10,2,3,4,5}
// 	large:=array[0]
// 	for i:=1;i<len(array);i++{
// 		if array[i]>large{
// 			large=array[i]
// 		}
// 	}
// 	fmt.Println(large)
// }

// Find second largest element
// func main(){
// 	array:=[]int{5,2,10,3,9,4,1}
// 	large:=array[0]
// 	second:=array[0]

// 	for i:=1;i<len(array);i++{
// 		if array[i]>large{
// 			second=large
// 			large=array[i]
// 		}else if array[i]>second&&array[i]!=large{
// 			second=array[i]
// 		}
// 	}
// 	fmt.Println(second)
// }

// Remove duplicates from slice
// func main(){
// 	array:=[]int{1,2,3,4,1,2,5,6}
// 	dupe:=[]int{}
// 	m:=make(map[int]bool)
// 	for _,v:=range array{
// 		if !m[v]{
// 			m[v]=true
// 			dupe=append(dupe, v)
// 		}
// 	}
// 	fmt.Println(dupe)
// }

// Reverse an array
// func main(){
// array:=[]int{1,2,3,4,5}
// f:=0
// l:=len(array)-1
// for f<l{
// 	array[f],array[l]=array[l],array[f]
// 	f++
// 	l--
// }
// fmt.Println(array)
// }

// Count vowels in string
// func main(){
// 	str:="hyy my name is shabin"
// 	v:=0
// 	for i:=0;i<len(str);i++{
// 		ch:=str[i]
// 		if ch=='a'||ch=='e'||ch=='i'||ch=='o'||ch=='u'{
// 			v++
// 		}
// 	}
// 	fmt.Println(v)
// }

// Count words in sentence
// func main(){
// 	words:="Hyy my name is shabin"
// 	count:=0
// 	inWord:=false

// 	for i:=0;i<len(words);i++{
// 		if words[i]!=' '&&!inWord{
// 			inWord=true
// 			count++
// 		}
// 		if words[i]==' '{
// 			inWord=false
// 		}
// 	}
// 	fmt.Println(count)
// }

//Check anagram strings
// func main(){
// 	a:="listen"
// 	b:="silent"
// 	if len(a)!=len(b){
// 		fmt.Println(false)
// 		return
// 	}
// 	m:=make(map[rune]int)
// 	for _,v:=range a{
// 		m[v]++
// 	}
// 	for _,v:=range b{
// 		m[v]--
// 	}
// 	for _,v:=range m{
// 		if v!=0{
// 			fmt.Println(false)
// 			return
// 		}
// 	}
// 	fmt.Println(true)
// }

// Fibonacci series
// func main(){
// 	a:=0
// 	b:=1
// 	for i:=0;i<10;i++{
// 		fmt.Println(a)
// 		next:=a+b
// 		a=b
// 		b=next
// 	}
// }

// func main(){
// 	num:=70
// isPrime:=true
// 	if num<=1{
// 		isPrime=false
// 	}else{
// 	for i:=2;i<num;i++{
// 		if num%i==0{
// 			isPrime=false
// 			break
// 		}
// 	}
// }
// if isPrime{
// 	fmt.Println("Is prime")
// }else{
// 	fmt.Println("Not prime")
// }
// }

// Sum of digits of a number
// func main(){
// 	result:=0
// 	n:=10
// 	for i:=1;i<=n;i++{
// 		result+=i
// 	}
// 	fmt.Println(result)
// }

// Sort slice without built-in sort
// func main(){
// 	array:=[]int{2,3,4,1,5,7,8,6}
// 	for i:=0;i<len(array);i++{
// 		for j:=i+1;j<len(array);j++{
// 			if array[i]>array[j]{
// 				array[i],array[j]=array[j],array[i]
// 			}
// 		}
// 	}
// 	fmt.Println(array)
// }

// Frequency count of characters
// func main(){
// 	str:="hyy my name is shabin"
// 	m:=make(map[rune]int)
// 	for _,v:=range str{
// 		m[v]++
// 	}
// 	for a,b:=range m{
// 		fmt.Printf("%c:%d\n",a,b)
// 	}
// }

// Frequency count of integers in slice
// func main(){
// 	array:=[]int{2,1,2,5,6,3,2,3,10,9,8,7,4,6,7,3,4,2,5,7,6,3,4,4,2,5,5,6}
// 	m:=make(map[int]int)
// 	for _,v:=range array{
// 		m[v]++
// 	}
// 	fmt.Println(m)
// }

// Find first non-repeating character
// func main(){
// 	array:=[]int{1, 2, 3, 4, 3, 5, 6, 7}
// 	m:=make(map[int]bool)
// 	n:=make(map[int]int)
// 	for _,v:=range array{
// 		if m[v]{
// 			fmt.Println("First repeating : ",v)
// 		}
// 		m[v]=true
// 	}
// 	for _,v:=range array{
// 		n[v]++
// 	}
// 	for _,v:=range array{
// 		if n[v]==1{
// 			fmt.Println("First non-repeating :",v)
// 			break
// 		}
// 	}
// }

// Find missing number in array
// func main(){
// 	array:=[]int{2,3,4,5}
// 	n:=len(array)+1
// 	sum:=0
// 	for _,v:=range array{
// 		sum+=v
// 	}
// 	exp:=n*(n+1)/2
// 	miss:=exp-sum
// 	fmt.Println(miss)
// }

// Two Sum problem
// func main(){
// 	array:=[]int{1,2,3,5,0}
// 	target:=8
// 	m:=make(map[int]int)
// 	for i,n:=range array{
// 		get:=target-n
// 		if ind,fou:=m[get];fou{
// 			fmt.Println(ind,i)
// 			return
// 		}
// 		m[n]=i
// 	}
// 	fmt.Println("Nothing")
// }

// Merge two arrays
// func main(){
// 	a:=[]int{1,2,3}
// 	b:=[]int{4,5,6}
// 	merge:=append(a,b...)
// 	fmt.Println(merge)
// }

//odd and even concurrently
// func Odd(oddch chan bool,evench chan bool,wg *sync.WaitGroup){
// 	defer wg.Done()
// 	for i:=1;i<=10;i+=2{
// 		<-oddch
// 		fmt.Println("odd :",i)
// 		evench<-true
// 	}
// }
// func Even(oddch chan bool,evench chan bool,wg *sync.WaitGroup){
// 	defer wg.Done()
// 	for i:=2;i<=10;i+=2{
// 		<-evench
// 		fmt.Println("even :",i)
// 		if i!=10{
// 			oddch<-true
// 		}
// 	}
// }
// func main(){
// 		oddch:=make(chan bool)
// 		evench:=make(chan bool)

// 		var wg sync.WaitGroup
// 		wg.Add(2)
// 		go Odd(oddch,evench,&wg)
// 		go Even(oddch,evench,&wg)
// 	oddch<-true
// 		wg.Wait()
// }

// func main(){
// 	a:=[]int{1,2,3,6}
// 	b:=[]int{2,3,4,5}

// 	m:=make(map[int]int)
// 	for _,v:=range a{
// 		m[v]++
// 	}
// 	for _,v:=range b{
// 		m[v]--
// 	}
// 	found:=false
// 	small:=0
// 	for i,n:=range m{
// 		if n==0{
// 			if !found||i<small{
// 				small=i
// 				found=true
// 			}
// 		}
// 	}
// 	if found{
// 		fmt.Println(small)
// 		return
// 	}
// 	for s,h:=range m{
// 		if h>0{
// 			fmt.Println(s)
// 			return
// 		}
// 	}
// }


// Intersection of arrays
// func main(){
// 	arr1:=[]int{1,2,3,4,5}
// 	arr2:=[]int{6,7,2,2,8,9,1}
// 	m:=make(map[int]bool)
// 	res:=[]int{}
// 	for _,v:=range arr1{
// 		m[v]=true
// 	}
// 	for _,v:=range arr2{
// 		if m[v]{
// 			res=append(res,v)
// 		}
// 	}
// 	fmt.Println(res)
// }

func main(){
	fmt.Printf("asdfghjk")
}