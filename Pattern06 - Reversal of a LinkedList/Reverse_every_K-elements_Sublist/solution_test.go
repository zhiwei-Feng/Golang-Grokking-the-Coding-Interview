package Reverse_every_K_elements_Sublist

import (
	"reflect"
	"testing"
)

func Test_reverse(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}

	head1 := &ListNode{Value: 1}
	head1.Next = &ListNode{Value: 2}
	node3 := &ListNode{Value: 3}
	head1.Next.Next = node3
	head1.Next.Next.Next = &ListNode{Value: 4}
	head1.Next.Next.Next.Next = &ListNode{Value: 5}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"case1", args{head1, 3}, node3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
