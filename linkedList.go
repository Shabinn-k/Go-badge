package main

import "fmt"

// type Node struct {
// 	data int
// 	Next *Node
// }

// func Linked() *Node {
// 	node1 := Node{data: 10}
// 	node2 := Node{data: 20}
// 	node3 := Node{data: 30}
// 	node4 := Node{data: 40}
// 	node5 := Node{data: 50}
// 	node6 := Node{data: 60}
// 	node7 := Node{data: 70}

// 	head := &node1
// 	node1.Next = &node2
// 	node2.Next = &node3
// 	node3.Next = &node4
// 	node4.Next = &node5
// 	node5.Next = &node6
// 	node6.Next = &node7

// 	return head
// }

// func Print(head *Node) {
// 	temp := head
// 	for temp != nil {
// 		fmt.Println(temp.data)
// 		temp = temp.Next
// 	}
// }

// func Middle(head *Node) {
// 	slow := head
// 	fast := head

// 	for fast != nil && fast.Next != nil {
// 		slow = slow.Next
// 		fast = fast.Next.Next
// 	}
// 	fmt.Println("Middle :", slow.data)
// }

// func ReverseLL(head *Node) *Node {
// 	var prev *Node
// 	current := head
// 	for current != nil {
// 		next := current.Next
// 		current.Next = prev
// 		prev = current
// 		current = next
// 	}
// 	return prev
// }

// func RemoveDuplicates(head *Node){
// 	current:=head
// 	for current!=nil&&current.Next!=nil{
// 		if current.data==current.Next.data{
// 			current.Next=current.Next.Next
// 		}else{
// 			current=current.Next
// 		}
// 	}

// }

// type Node struct{
//     data int
//     next *Node
// }

// func LL()*Node{
//     node1:=&Node{data:10}
//     node2:=&Node{data:10}
//     node3:=&Node{data:30}
//     node4:=&Node{data:40}
//     node5:=&Node{data:30}
    
//     node1.next=node2
//     node2.next=node3
//     node3.next=node4
//     node4.next=node5
    
//     head:=node1
//     return head
    
// }

// func Middle(head *Node){
//     slow:=head
//     fast:=head
//     for fast!=nil&&fast.next!=nil{
//         slow=slow.next
//         fast=fast.next.next
//     }
//     fmt.Println("Middle :",slow.data)
// }

// func Print(head *Node){
//     temp:=head
//     for temp!=nil{
//         fmt.Println(temp.data)
//         temp=temp.next
//     }
// }

// func Reverse(head *Node)*Node{
//     var prev *Node
//     current:=head
//     for current!=nil{
//         next:=current.next
//         current.next=prev
//         prev=current
//         current=next
//     }
//     return prev
// }

// func Duplicate(head *Node)*Node{
//     current:=head
//     for current!=nil&&current.next!=nil{
//         if current.data==current.next.data{
//             current.next=current.next.next
//         }else{
//             current=current.next
//         }
//     }
//     return head
// }

// func Duplicate(head *Node)*Node{
//     var prev *Node
//     current:=head
//     seen:=make(map[int]bool)
    
//     for current!=nil{
//         if seen[current.data]{
//             prev.next=current.next
//         }else{
//             seen[current.data]=true
//             prev=current
//         }
//         current=current.next
//     }
//     return head
// }

// func main(){
//     head:=LL()
    
//     fmt.Println("Original :")
//     Print(head)
    
//     Middle(head)
    
//     fmt.Println("Reverse :")
//     tail:=Reverse(head)
//     Print(tail)
    
//     fmt.Println("Duplicate :")
//     dupe:=Duplicate(LL())
//     Print(dupe)
// }

type Node struct{
	data int
	Next *Node
	Prev *Node
}

func DoubleLL(){
	node1:=Node{data: 10}
	node2:=Node{data: 20}
	node3:=Node{data: 30}
	head:=&node3
	node1.Next=&node2
	node2.Prev=&node1
	node2.Next=&node3
	node3.Prev=&node2

	temp:=head
	for temp!=nil{
		fmt.Println(temp.data)
		temp=temp.Prev
	}
}
