package ch1

// 检查输入字符串的括号是否配对完整，对于[()]{}{[()()]()} 程序应该打印true，对于[(]) 则打印false
func parenthesesValid(src string) bool {
	stack := []rune{}
	for _, v := range src {

		if v == '(' || v == '{' || v == '[' {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if v == '}' && top != '{' {
				return false
			} else if v == ']' && top != '[' {
				return false
			} else if v == ')' && top != '(' {
				return false
			}
			stack = stack[0 : len(stack)-1]
		}
	}

	return len(stack) == 0
}

func completeLeftParentheses(src string) string {

}
