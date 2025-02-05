package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)

	for i := 0; i < 5; i++ {
		wg.Add(1) // Увеличиваем счетчик горутин в группе
		go work(i, wg) // Вызываем функцию work в отдельной горутине
	}

	wg.Wait() // ожидаем завершения всех горутин в группе
	fmt.Println("all gorutine done")
}

func work(id int, wg *sync.WaitGroup)  {
	defer wg.Done()
	fmt.Printf("gorutine %d start \n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("gorutine %d stoped \n", id)
}
