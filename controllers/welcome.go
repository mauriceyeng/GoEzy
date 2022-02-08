package main

import (
	"fmt"
	"GoEzy/structs"
)
func Welcome() string {
	msg:=structs.WelcomeMsg{
		Msg: "Welcome to GoEzy"
	}
	return Welcome()
}
