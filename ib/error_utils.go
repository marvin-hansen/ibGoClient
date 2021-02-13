package ib

import "fmt"

func CheckNoError(err error, msg string) bool {
	if err != nil {
		fmt.Println(msg, err)
		return false
	} else {
		return true
	}
}
