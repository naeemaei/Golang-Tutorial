package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

var TodoList = []string{}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(50) // 50 - 50 = 0
	for i := 0; i < 50; i++ {
		//wg.Add(1)
		go GetTodo(i+1, &wg)
	}

	wg.Wait()
	fmt.Printf("%v", TodoList)

}

func GetTodo(id int, wg *sync.WaitGroup) {
	//"https://jsonplaceholder.typicode.com/todos/1"
	GetUrl("https://jsonplaceholder.typicode.com/todos/"+strconv.Itoa(id), wg)
}

func GetUrl(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	responseBody, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		panic(err)
	}

	TodoList = append(TodoList, string(responseBody))

}

// main =>
// fork task1 task2 task3 task4 task5

//join

// 5
// -1 => 4
// -1 => 3
// ...
// -1 => 0
