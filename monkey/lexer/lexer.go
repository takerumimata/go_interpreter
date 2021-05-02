package lexer

type Lexer struct {
	input        string
	position     int  // 入力における現在の位置
	readPosition int  //これから読み込むいち（現在の文字の次）
	ch           byte // 現在検査中の文字
}
