package main

import (
  "fmt"
  "net/http"
)


func main() {
  url := "https://www.shanon.co.jp"
  resp, _ := http.Get(url)
  defer resp.Body.Close()
  fmt.Println(resp.StatusCode) // htmlをstringで取得
}
