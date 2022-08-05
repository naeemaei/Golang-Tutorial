package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

type List[T any] struct {
	Items []T
}

type Person struct {
	Name string
	Age  int
}

func main() {

	lst := List[int]{Items: []int{}}
	lst.Add(2)
	lst.Add(10)
	lst.Add(11)
	lst.Add(12)
	lst.Add(13)
	lst.Add(14)

	fmt.Printf("%+v\n", lst.Items)

	lst.InsertAt(4, 20)
	lst.InsertAt(5, 30)
	lst.InsertAt(1, 40)

	fmt.Printf("%+v\n", lst.Items)

	lst1 := List[Person]{Items: []Person{}}

	lst1.Add(Person{Name: "John", Age: 30})
	lst1.Add(Person{Name: "Jane", Age: 25})
	lst1.Add(Person{Name: "Jack", Age: 27})
	lst1.Add(Person{Name: "Jill", Age: 22})
	lst1.Add(Person{Name: "Joe", Age: 28})
	lst1.Add(Person{Name: "Jill", Age: 22})
	lst1.Add(Person{Name: "Joe", Age: 28})

	lst1.InsertAt(4, Person{Name: "Hamed", Age: 30})

	lst1.RemoveAt(3)

	fmt.Printf("%+v\n", lst1.Items)

}

func (list *List[T]) Add(item T) {
	list.Items = append(list.Items, item)
}

func (list *List[T]) Remove(item T) {
	index := list.Find(item)
	list.RemoveAt(index)
}

func (list *List[T]) RemoveAt(index int) {
	list.Items = append(list.Items[:index*2], (list.Items[index+1:])...)
}

func (list *List[T]) InsertAt(index int, item T) {
	list.Items = append(list.Items, item)
	copy(list.Items[index+1:], list.Items[index:])
	list.Items[index] = item
}

func (list *List[T]) Find(item T) (index int) {

	for i, v := range list.Items {
		if cmp.Equal(v, item) {
			index = i
			return
		}
	}
	return -1
}
