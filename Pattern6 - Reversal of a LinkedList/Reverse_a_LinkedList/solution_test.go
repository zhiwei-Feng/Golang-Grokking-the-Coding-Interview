package Reverse_a_LinkedList

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_reverse(t *testing.T) {
	type args struct {
		head *ListNode
	}

	head := &ListNode{Value: 2}
	head.Next = &ListNode{Value: 4}
	head.Next.Next = &ListNode{Value: 6}
	head.Next.Next.Next = &ListNode{Value: 8}
	tail := &ListNode{Value: 10}
	head.Next.Next.Next.Next = tail

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"case1", args{head: head}, tail},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			} else {
				for p := got; p != nil; p = p.Next {
					fmt.Printf("%d ", p.Value)
				}
			}
		})
	}
}
