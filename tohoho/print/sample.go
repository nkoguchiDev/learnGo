package main		// mainパッケージであることを宣言

import "fmt"		// fmtモジュールをインポート

func main() {		// 最初に実行されるmain()関数を定義
    num := 123
	str := "ABC"

	fmt.Print("num=", num, " str=", str, "\n")	// 改行無し、空白無し、フォーマット無し
	fmt.Println("num =", num, "str =", str)	// 改行有り、空白有り、フォーマット無し
	fmt.Printf("num=%d str=%s\n", num, str)	// 改行無し、空白無し、フォーマット有り
}

