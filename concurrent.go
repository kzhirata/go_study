package main

import (
	"fmt"
	"net/http"
)

func request() int {
	url := "https://www.shanon.co.jp"
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	return resp.StatusCode
}

func main() {
	// int型チャネルの作成
	c := make(chan int)

	// 送信専用チャネルを受け取り、1〜10までの数値を送信する
	go func(s chan<- int) {
		for i := 0; i < 10; i++ {
			code := request()
			s <- code
		}
		close(s)
	}(c)

	for {
		// チャネルからの受信を待機
		val, ok := <-c
		if !ok {
			// チャネルがクローズしたので、終了する
			break
		}
		// 受信したデータを表示
		fmt.Println(val)
	}
}
