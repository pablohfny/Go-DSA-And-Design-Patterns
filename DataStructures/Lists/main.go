package main

import "fmt"

type List struct {
	Val  any
	Next *List
}

func (list *List) PushBack(val any) {
	next := list.Next

	if next != nil {
		next.PushBack(val)
	} else {
		list.Next = &List{
			Val: val,
		}
	}
}

func main() {
	var intList List
	intList.PushBack(11)
	intList.PushBack(23)
	intList.PushBack(34)

	for element := intList.Next; element != nil; element = element.Next {
		fmt.Println(element.Val)
	}
}
