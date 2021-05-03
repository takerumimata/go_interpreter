# Goのプロジェクトをどう始めれば良い？
以前少しだけGoを書いた記憶がある。その際はGoはローカル環境でも特定のディレクトリにプロジェクトを作らなければならなかったはずだ。
だが今は`go mod`の登場でそんなけったいなことをしなくていいようだ。

# プロジェクトの初期化
以下のコマンドを叩く
```zsh
$ go mod init github.com/takerumimata/<repository-name>
```
そうするとルートディレクトリにgo.modができる。

# localで作ったpackageをimportする
これですら情報がとっ散らかっててよくわからなかった。
相対パスでかけるもんかと思ったらそうでもないようだったり、絶対パスでかけと言われたりでややこしかった。

go.modで指定したmodule名をprefixにして、プロジェクトルートからの絶対パスを記述することでimportすることができるようだ

```go
module github.com/takerumimata/go_interpreter

go 1.16

```
上記のような`go.mod`ファイルが与えられているのであれば、以下のような形式で指定すれば良いようだ

```go
// monkey/lexer/lexer_test.go の例
package lexer

import (
	"testing"

	"github.com/takerumimata/go_interpreter/monkey/token"
)
```