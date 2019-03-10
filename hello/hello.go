//Goでは１つのファイルに記述できるのは単一のパッケージのみ
//２つ以上記述するとコンパイルエラーになる
package main

//ライブラリのインポート
//省略できる
//fmtは入出力関係
//参照しないパケージが記述されている場合はコンパイルエラー
import (
	"fmt"
)

//Go(mainパッケージの中のmain)のエントリポイント
func main() {
	fmt.Println("Hello, World!")

	//エラーのインターフェースerrorに定義されているメソッドはError()なのでZ
	//構造体MyErrorのメソッドとしてEorro()を作成
	//RaiseErrorでMyErrorをnewしたものをもらってインターフェースで定義されていたErrorを呼び出してみる
	err := RaiseError()
	fmt.Println(err.Error()) //エラーが発生
}

///////////////////////////////////////////////////////////////////////////
//以下はerrorのインターフェースを適用した構造体作成テスト
//////////////////////////////////////////////////////////////////////////

//独自のえらーを表す型
type MyError struct {
	Message string
	ErrCode int
}

//errorインターフェースのメソッド
func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{Message: "エラーが発生したよ！", ErrCode: 1234}
}
