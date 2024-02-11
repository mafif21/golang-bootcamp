package sessiontwo

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func MiniChallengeTwo() {

	fmt.Print("Enter the sentences:")

	in := bufio.NewReader(os.Stdin)
	sentences, err := in.ReadString('\n')
	if err != nil {
		panic(err)
	}

	sentences = strings.TrimSpace(sentences)
	separateWord := strings.Split(sentences, "") // array
	penampungData := make(map[string]int)

	for _, word := range separateWord {
		if _, exists := penampungData[word]; exists {
			penampungData[word] += 1
		} else {
			penampungData[word] = 1
		}
	}

	fmt.Println(penampungData)

}
