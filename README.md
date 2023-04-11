# for文とgoroutineとsync.WaitGroupの使い方

func A : funcAスタート　→　1秒待機　→ funcA終了
func B : funcBスタート　→　5秒待機　→ funcB終了

## 大前提
funcAとfuncBをfor文で3回実行する

これらの順序制御するテンプレを作成する
何も使わない基本コードは↓

```go
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
```

当然こうなる　output↓


```
funcA started
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
funcA finished
funcB started
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
funcB finished
funcA started
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
funcA finished
funcB started
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
funcB finished
funcA started
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
funcA finished
funcB started
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
funcB finished
All processes finished
```

流れ
A → B → A → B → A → B → finished

# 順序関係なく、とにかく早く終わらせたい
wg.Wait()をfor文の外につけてfuncAとfuncBの頭にgoをつける
```go
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

```

output↓


```
funcB started
funcB started
funcA started
funcA started
funcB started
funcA started
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
funcA finished
funcA finished
funcA finished
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
funcB finished
funcB finished
funcB finished
All processes finished
```
3回分のABがほぼ同時にスタートして、時間がかかるBが最後に終了する

# for文1サイクル毎にAとBの両方の完了を待ちたい
wg.Wait()をfor文の中につけてfuncAとfuncBの頭にgoをつける

```go

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



```

output↓
```
funcB started
funcA started
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
funcA finished
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
funcB finished
funcB started
funcA started
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
funcA finished
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
funcB finished
funcA started
funcB started
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
funcA finished
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
--------------0.3秒経過-----------------
funcB finished
All processes finished
```
AとBは同時にスタートするが、Bが終了するまで次にはいかない


# まとめ
## 順番通りに実行したい
goroutine使わない

## 順番関係なく早く処理を終わらせたい
for文の外にwg.Wait()を置く
```go
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(2)
		go func() { //funcA
			defer wg.Done()
			//省略
		}()
		go func() { //funcB
			defer wg.Done()
			//省略
		}()
	}
	wg.Wait()
```

## for文1サイクル毎にAとBの両方の完了を待ちたい
for文の中にwg.Wait()を置く
```go
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(2)
		go func() { //funcA
			defer wg.Done()
			//省略
		}()
		go func() { //funcB
			defer wg.Done()
			//省略
		}()
	wg.Wait()
	}
```


参考になれば幸いです