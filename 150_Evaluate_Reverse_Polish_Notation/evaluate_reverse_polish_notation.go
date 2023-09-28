package evaluatereversepolishnotation

import "strconv"

type stack struct {
	stack []int
}

func newStack() *stack {
	return &stack{}
}

func (st *stack) getTwoValues() (int, int) {
	first := st.stack[len(st.stack) - 2]
	second := st.stack[len(st.stack) - 1]
	st.stack = st.stack[:len(st.stack) - 2]
	return first, second
}

func evalRPN(tokens []string) int {
    stack := newStack()

	for _, token := range tokens {
		switch token {
		case "+":
			first, second := stack.getTwoValues()
			stack.stack = append(stack.stack, first+second)
		case "-":
			first, second := stack.getTwoValues()
			stack.stack = append(stack.stack, first-second)
		case "*":
			first, second := stack.getTwoValues()
			stack.stack = append(stack.stack, first*second)
		case "/":
			first, second := stack.getTwoValues()
			stack.stack = append(stack.stack, first/second)
		default:
			integer, _ := strconv.Atoi(token)
			stack.stack = append(stack.stack, integer)
		}
	}
	return stack.stack[0]
}