package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

func worker(id int, ch <-chan int, chResult chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for value := range ch {
		newValue := value * 2
		time.Sleep(1 * time.Second)
		chResult <- newValue
	}
}

func main() {
	numWorkers := 3
	wg := sync.WaitGroup{}
	ch := make(chan int, 10)
	chResult := make(chan int, 10)

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, ch, chResult, &wg)
	}
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
	wg.Wait()
	close(chResult)

	result := make([]int, 0, 10)
	for value := range chResult {
		result = append(result, value)
	}

	fmt.Println("Отсортированный массив:", result)
	sort.Ints(result)
	fmt.Println("неотсортированный массив:", result)
}
