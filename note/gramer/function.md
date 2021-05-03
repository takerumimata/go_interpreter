# Go言語の関数について学ぶ
Goにはメソッドと関数がある。なんのこっちゃわからなかったので調べながら分かったことをまとめていく。例に取るのは[/monkey/lexer/lexer.go](https://github.com/takerumimata/go_interpreter/blob/main/monkey/lexer/lexer.go)

```go
package lexer

type Lexer struct {
	input        string
	position     int  // 入力における現在の位置
	readPosition int  //これから読み込むいち（現在の文字の次）
	ch           byte // 現在検査中の文字
}

// 関数だよ
func New(input string) *Lexer { // 関数Newの引数にinputという変数名をつけていて、それはstring型だよ、と言っている。そして返り値は*lexer。つまりLexerのポインタを返している。
	l := &Lexer{input: input}   // &はアドレス演算子かな？ ポインタ型としてlを生成しているっぽい
	l.readChar()
	return l
}

// メソッドだよ
func (l *Lexer) readChar() {  // こっちはメソッド。
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

```

# Goはオブジェクト指向ではないが...
構造体にメソッドを実装することができるようだ。.演算子で構造体からメソッドを呼び出すことができるようだ。
具体的にはこんな感じ↓で本プロジェクトでは構造体を初期化し、readchar()始め、さまざまなメソッドを呼び出している
```go
// repl.go
l := lexer.New(line)

// lexer.go
~~~
l.readChar()
~~~
```
もうちょっとましな例をあとで出そう。今だからわかるけど後で見たらわからなそう。（雰囲気はわかるが）

# メソッドの定義方法
```go
func (<レシーバー>) <関数名>([引数]) [戻り値の型] {
    [関数の本体]
}
```

### 雑なまとめ
関数とメソッドに関して、関数は普通に引数を取るが、メソッドはレシーバからメソッドを呼ぶ。あと関数とメソッドは名前空間が違うらしいのでおんなじ関数名で宣言してもエラーにならないらしい。

# 参考
- [golang メソッドについてまとめてみた。](https://qiita.com/pei0804/items/2a3010df39623fadc3c6)
- [Goで学ぶポインタとアドレス](https://qiita.com/Sekky0905/items/447efa04a95e3fec217f)