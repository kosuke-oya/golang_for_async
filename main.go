package main

import (
	"fmt"
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
	for i := 0; i < 3; i++ {
		func() { //funcA
			fmt.Println("funcA started")
			time.Sleep(1000 * time.Millisecond)
			fmt.Println("funcA finished")
		}()
		func() { //funcB
			fmt.Println("funcB started")
			time.Sleep(3000 * time.Millisecond)
			fmt.Println("funcB finished")
		}()
	}

	fmt.Println("All processes finished")

}
