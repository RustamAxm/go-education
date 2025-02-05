package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	tick := time.NewTicker(time.Second)
	defer tick.Stop()

	wg := new(sync.WaitGroup)

	for i := 0; i < 5; i++ {
		wg.Add(1) // Увеличиваем счетчик горутин в группе
		go work(i, tick.C, wg) // Вызываем функцию work в отдельной горутине
	}

	wg.Wait() // ожидаем завершения всех горутин в группе
	fmt.Println("all gorutine done")
}

func work(id int, limit <- chan time.Time, wg *sync.WaitGroup)  {
	defer wg.Done()
	fmt.Printf("gorutine %d start \n", id)
	<- limit 
	fmt.Printf("gorutine %d stoped \n", id)
}
