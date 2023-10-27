package addTwoNumbers

import (
	"testing"
)

type test struct {
	l1 *ListNode
	l2 *ListNode
	result *ListNode
}



var testCases = []test {
	{
		&ListNode{2, &ListNode{4, &ListNode{3, nil}}}, 
		&ListNode{5, &ListNode{6, &ListNode{4, nil}}}, 
		&ListNode{7, &ListNode{0, &ListNode{8, nil}}}, 
	},
	{
		&ListNode{0, nil},
		&ListNode{0, nil},
		&ListNode{0, nil},
	},
	{
		&ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}},
		&ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}},
		&ListNode{8, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{1, nil}}}}}}}},
	},
}

func TestAddTwoNumbers(t *testing.T) {
	for _, test := range testCases {
		t.Run("AddTwoNumbers", func(t *testing.T) {
			result := AddTwoNumbers(test.l1,test.l2)
			expected := test.result

			for (result != nil) && (expected != nil) {
				if expected == nil {
					t.Errorf("Length Problem Expected is lower")
				}

				if result == nil {
					t.Errorf("Length Problem Result is lower")
				}
				
				if expected.Val != result.Val {
					t.Errorf("got %v, want %v", result.Val, expected.Val)
				}

				result = result.Next
				expected = expected.Next
			}

			if (result != nil) || (expected != nil) {
				t.Errorf("Length Problem")
			}
			
		})
	}
}