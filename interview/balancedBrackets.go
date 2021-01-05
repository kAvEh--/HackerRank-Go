package main

import "fmt"

func main() {
	fmt.Println(isBalanced("[{"))
}

// Complete the isBalanced function below.
func isBalanced(s string) string {
	if len(s)%2 == 1 {
		return "NO"
	}
	stack := ""
	for i := 0; i < len(s); i++ {
		tmp := string(s[i])
		if tmp == "{" || tmp == "(" || tmp == "[" {
			stack = tmp + stack
			continue
		}
		if len(stack) < 1 {
			return "NO"
		}
		if tmp == "}" && string(stack[0]) != "{" {
			return "NO"
		}
		if tmp == ")" && string(stack[0]) != "(" {
			return "NO"
		}
		if tmp == "]" && string(stack[0]) != "[" {
			return "NO"
		}
		stack = stack[1:]
	}
	if len(stack) > 0 {
		return "NO"
	}
	return "YES"
}
