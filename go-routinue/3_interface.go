package main

import (
	"fmt"
	"sort"
)

type SortableStrings []string

//type SortableStrings [3]string

type Sortable interface {
	sort.Interface
	Sort()
}

func (self SortableStrings) Len() int {
	return len(self)
}

func (self SortableStrings) Less(i, j int) bool {
	return self[i] < self[j]
}

func (self SortableStrings) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

func (self *SortableStrings) Sort() {
	sort.Sort(self)
}

func main() {
	_, ok := interface{}(SortableStrings{}).(sort.Interface)
	_, ok1 := interface{}(SortableStrings{}).(Sortable)
	_, ok2 := interface{}(&SortableStrings{}).(Sortable)
	fmt.Println("interface ok:", ok)
	fmt.Println("interface ok1:", ok1)
	fmt.Println("interface ok2:", ok2)

	ss := SortableStrings{"2", "3", "1"}
	ss.Sort()
	fmt.Printf("Sortable strings: %v\n", ss)
}
