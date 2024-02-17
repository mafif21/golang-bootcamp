package main

import (
	"golang-bootcamp/session_three"
	"os"
	"strings"
)

func main() {
	// sessionone.MiniChallengeOne()
	//sessiontwo.MiniChallengeTwo()

	var args = os.Args[1:2]
	studentName := strings.Join(args, "")
	session_three.MiniChallengeThree(studentName)
}
