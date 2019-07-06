package main

import "fmt"

// メイン関数
func main() {
    // make関数でいきなりスライスを作成する
    s1 := make([]int, 3) // [0 0 0]

    // いきなり値を割り当てたスライスを作成する
    s2 := []int{1, 3, 5} // 配列の宣言と似ている
    fmt.Println(s1)
    fmt.Println(s2)

    // appendでスライスの末尾に要素を追加
    s3 := append(s2, 8, 2, 10)
    fmt.Println(s3)

    // copyでスライスをコピー
    // 返り値は要素数
    t := make([]int, len(s3))
    s4 := copy(t, s3)
    fmt.Println(s4)
}
