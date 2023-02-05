package main

import (
	"fmt"
)

type element struct {
	data int
	next *element
}

func genNewList(data int) element {
	return element{data, nil}
}

func printList(list *element) {
	for {
		fmt.Println(list.data)
		if list.next != nil {
			list = list.next
		} else {
			break
		}
	}
}

func addToEnd(list *element, data int) {
	for {
		if list.next != nil {
			list = list.next
		} else {
			list.next = &element{data: data, next: nil}
			break
		}
	}
}

func addToStart(list element, data int) element {
	newList := element{data: data, next: &list}
	return newList
}

func addAfter(list *element, index int, data int) {
	counter := 0
	for {
		if counter != index {
			if list.next != nil {
				counter++
				list = list.next
			} else {
				fmt.Println("The specified index goes beyond the bounds of the list!")
				break
			}
		} else {
			newElement := element{data: data, next: list.next}
			list.next = &newElement
			break
		}
	}
}

func delLast(list *element) {
	for {
		if list.next != nil {
			list = list.next
			if list.next.next == nil {
				list.next = nil
				break
			}
		}
	}
}

func delFirst(list *element) element {
	if list.next != nil {
		return *list.next
	} else {
		fmt.Println("There is nothing to delete!")
	}
	return *list
}

func delAfter(list *element, index int) {
	counter := 0
	for {
		if counter != index {
			if list.next != nil {
				counter++
				list = list.next
			} else {
				fmt.Println("The specified index goes beyond the bounds of the list!")
				break
			}
		} else {
			list.next = list.next.next
			break
		}
	}
}

func main() {
	list := genNewList(0)
	addToEnd(&list, 5)
	list = addToStart(list, 10)
	addAfter(&list, 5, 20)
	printList(&list)
	fmt.Println("---")
	delAfter(&list, 0)
	printList(&list)
}
