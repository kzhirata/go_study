package main

import "fmt"

// 引数なし関数
func sayHi() {
    fmt.Println("hi!")
}

// 引数あり関数
func sayName(name string) {
    fmt.Println(name)
}

// return関数
func getMessage(name string) string {
    msg := "hi! my name is " + name
    return msg
}

// 名前付きreturn関数
// 関数内で使った変数名を返す
func getHelloMessage(name string) (msg string) {
    msg = "Hello " + name
    return
}

// メイン関数
func main() {
    sayHi()
    sayName("gcfuji")
    fmt.Println(getMessage("gcfuji"))
    fmt.Println(getHelloMessage("Gemcook"))
}
