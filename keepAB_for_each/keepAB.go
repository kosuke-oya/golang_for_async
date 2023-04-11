package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//時間経過線を入れるためのコード
	go func() {
		for {
			time.Sleep(300 * time.Millisecond) //わかりやすいように0.3秒ごとに線を表示
			fmt.Println("--------------0.3秒経過-----------------")
		}
	}()
	//時間経過線を入れるためのコード終了
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(2)
		go func() { //funcA
			defer wg.Done()
			fmt.Println("funcA started")
			time.Sleep(1000 * time.Millisecond)
			fmt.Println("funcA finished")
		}()
		go func() { //funcB
			defer wg.Done()
			fmt.Println("funcB started")
			time.Sleep(3000 * time.Millisecond)
			fmt.Println("funcB finished")
		}()
		wg.Wait()
	}

	fmt.Println("All processes finished")

}
