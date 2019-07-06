package main

import "fmt"

// メイン関数
func main() {
    // mapを宣言する
    // m := make(map[string]int)
    //
    // // mapに値を設定
    // m["fujimoto"] = 200
    // m["arita"] = 300

    // 値を指定しながら宣言する
    m := map[string]int{"fujimoto": 100, "arita": 200} // map[fujimoto:100 arita:200]

    m["itou"] = 3000

    // キーの存在を調べる
    v, has_key := m["fujimoto"]

    fmt.Println(v)  // 100
    fmt.Println(has_key) // true

    // 要素を削除する
    delete(m, "fujimoto")
    fmt.Println(m) // map[arita:200]
}
