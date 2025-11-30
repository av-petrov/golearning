package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
	s5 := &List[int]{val: 1969, next: nil}
	s4 := &List[string]{val: "Appolo", next: nil}
	s3 := &List[int]{val: 1961, next: s5}
	s2 := &List[string]{val: "Vostok", next: s4}
	s1 := &List[int]{val: 1957, next: s3}
	s := &List[string]{val: "Sputnik", next: s2}
	currentS := s
	currentY := s1
	for {
		fmt.Printf("Event: %v, Year: %v\n", currentS.val, currentY.val)
		if currentS.next == nil || currentY.next == nil {
			break
		}
		currentS = currentS.next
		currentY = currentY.next
	}
}
