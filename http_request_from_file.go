package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func request(url string) int {
//	fmt.Println(url)
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	return resp.StatusCode
}

func main() {
	var fp *os.File
	var err error

	if len(os.Args) < 2 {
		fp = os.Stdin
	} else {
		fp, err = os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer fp.Close()
	}

	urls := []string{}
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("ファイル読込完了")

	// int型チャネルの作成
	c := make(chan int)

	// 送信専用チャネルを受け取り、1〜10までの数値を送信する
	go func(s chan<- int) {
		for _, url := range urls {
			code := request(url)
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
        fmt.Println("全http request完了")
}
