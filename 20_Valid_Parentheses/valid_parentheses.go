package validparentheses


func isValid(s string) bool {
    stack := make([]rune, 0)
    for _, symbol := range s {
        switch symbol {
        case '(','{','[':
            stack = append(stack, symbol)
        case ')':
            if len(stack) == 0 {
                return false
            }
            if stack[len(stack) - 1] != '(' {
                return false
            }
            stack = stack[:len(stack) - 1]
        case '}':
            if len(stack) == 0 {
                return false
            }
            if stack[len(stack) - 1] != '{' {
                return false
            }
            stack = stack[:len(stack) - 1]
        case ']':
            if len(stack) == 0 {
                return false
            }
            if stack[len(stack) - 1] != '[' {
                return false
            }
            stack = stack[:len(stack) - 1]
        }
    }
    return len(stack) == 0 
}