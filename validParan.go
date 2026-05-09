package main

func isValid(s string)bool{
	stack:=[]rune{}
pairs:=map[rune]rune{
	')':'(',
	'}':'{',
	']':'[',
}
for _,char:=range s{
	if opn,exist:=pairs[char];exist{
		if len(stack)==0||stack[len(stack)-1]!=opn{
			stack=stack[:len(stack)-1]
		}else{
			stack = append(stack,char)
		}
	}
}
return len(stack)==0
}