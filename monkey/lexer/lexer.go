package lexer

import (
	"github.com/takerumimata/go_interpreter/monkey/token"
)

type Lexer struct {
	input        string
	position     int  // 入力における現在の位置
	readPosition int  //これから読み込むいち（現在の文字の次）
	ch           byte // 現在検査中の文字
}

// 字句解析を行うにあたり、初期化を行う。任意の入力文字列を受け取って、トークンを作成する
// positon、readpositon、chの初期値って何？ 多分0かnullだろうけど
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// 次の文字を読んでいく関数
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// 現在検査中の文字l.chをみてその文字がなんであるかに応じてトークンを返す。
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
