package session_four

import (
	"fmt"
	"testing"
	"time"
)

func sayHello1(index int) {
	fmt.Println("hello ke-", index)
}

func sayHello2(index int) {
	fmt.Println("hello ke-", index)
}

func TestMutex(t *testing.T) {
	for i := 1; i <= 4; i++ {
		sayHello1(i)
		sayHello2(i)

		time.Sleep(3 * time.Second)
	}
}
