package mergeksortedlists

import (
	"reflect"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}

	link1 := &ListNode{5, nil}
	link1.Next = &ListNode{8, nil}
	link1.Next.Next = &ListNode{9, nil}

	link2 := &ListNode{1, nil}
	link2.Next = &ListNode{7, nil}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", args{[]*ListNode{link1, link2}}, []int{1, 5, 7, 8, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeKLists(tt.args.lists)
			gotArr := make([]int, 0)
			for got != nil {
				gotArr = append(gotArr, got.Val)
				got = got.Next
			}

			if !reflect.DeepEqual(gotArr, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", gotArr, tt.want)
			}
		})
	}
}
