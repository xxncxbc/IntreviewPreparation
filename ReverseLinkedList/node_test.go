package ReverseLinkedList

import (
	"fmt"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	tests := []struct {
		nums []int
		want string
	}{
		{[]int{1, 2, 3}, "3->2->1->nil"},
		{[]int{1, 0, 1, 0, 1}, "1->0->1->0->1->nil"},
		{[]int{}, "nil"},
		{[]int{1, 2}, "2->1->nil"},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "10->9->8->7->6->5->4->3-2->1->nil"},
	}

	for i, test := range tests {
		name := fmt.Sprintf("test number %d \n %v", i, test.nums)
		t.Run(name, func(t *testing.T) {
			linkedList := GenerateList(test.nums...)
			linkedList = ReverseLinkedList(linkedList)
			got := StringList(linkedList)
			if test.want != got {
				t.Error("wrong results!")
			}
		})
	}
}
