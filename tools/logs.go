package tools

import "fmt"

func IsError(e error) bool {
	if e != nil {
		fmt.Println(e)
		return true
	}
	return false
}
