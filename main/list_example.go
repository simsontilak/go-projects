package main

import (
	"container/list"
	"fmt"
)

func main() {

	//list do not represent go idioms very well use slice instead
	//this is a doubly linked list
	//ie has reference to its previous element and reference to the next element.
	//you can go forward or backward
	//it preserves order
	var myList list.List

	myList.PushBack(55)
	myList.PushBack(12)
	myList.PushBack(44)
	myList.PushBack(88)
	myList.PushBack(104)
	myList.PushBack(22)

	for element := myList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	var givenElement = myList.Front()

	fmt.Println("Another way to navigate")
	for i := 0; i < myList.Len(); i++ {
		fmt.Println(givenElement.Value)
		givenElement = givenElement.Next()
	}
	fmt.Println("done")

}
