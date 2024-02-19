package session_four

import (
	"fmt"
	"strings"
	"sync"
	"testing"
)

func processRandom1(data interface{}, i int, group *sync.WaitGroup) {
	defer group.Done()
	fmt.Printf("[%s] %v\n", data, i+1)
}

func processRandom2(data interface{}, i int, group *sync.WaitGroup) {
	defer group.Done()
	fmt.Printf("[%s] %v\n", data, i+1)
}

func TestRandomNumber(t *testing.T) {
	wg := &sync.WaitGroup{}

	data1 := []string{"coba1", "coba2", "coba3"}
	data2 := []string{"bisa1", "bisa2", "bisa3"}

	wg.Add(8)
	go func() {
		for i := 0; i < 4; i++ {
			go processRandom1(strings.Join(data1, " "), i, wg)
			go processRandom2(strings.Join(data2, " "), i, wg)
		}

	}()

	wg.Wait()
}

func ProcessMutex1(data string, i int, group *sync.WaitGroup, mutex *sync.Mutex) {
	defer group.Done()
	mutex.Lock()
	defer mutex.Unlock()
	fmt.Printf("[%s] %d\n", data, i+1)
}

func ProcessMutex2(data string, i int, group *sync.WaitGroup, mutex *sync.Mutex) {
	defer group.Done()
	mutex.Lock()
	defer mutex.Unlock()
	fmt.Printf("[%s] %d\n", data, i+1)
}

func TestMutexNew(t *testing.T) {
	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}

	data1 := []string{"coba1", "coba2", "coba3"}
	data2 := []string{"bisa1", "bisa2", "bisa3"}

	wg.Add(8)
	go func() {
		for i := 0; i < 4; i++ {
			go ProcessMutex1(strings.Join(data1, " "), i, wg, mutex)
			go ProcessMutex2(strings.Join(data2, " "), i, wg, mutex)
		}

	}()

	wg.Wait()
}
