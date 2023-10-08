package lexer

import (
	"go-interpreter/token"
)

// For simplicity only supports ASCII
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // calling readChar here so our lexer is in a fully working state when initialized
	return l
}

/*
* The purpose of readChar is to give us the next character and advance
* our position in the input string
 */
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // 0 is the ASCII code for "NULL", signifies either "we haven't read anything" or "end of file"
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

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
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			/**
			* The early return here is necessary because when calling "readIdentifier()"
			* we call "readChar()" repeatedly and advance our "readPosition" and "position"
			* fields past the last character of the current identifier. So we don't need
			* the call to "readChar()" after the switch statement again.
			 */
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// Reads in an identifier and advances our lexer's positions until it encounters a non-letter-character
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

/*
* Changing this function has a larger impact on the lanugage our interpreter will be able
* parse than one would expect from such a small function. As you can see, in our case it contains
* the chec "ch == '_' ", which means that we'll treat "_" as a letter and allow it in identifiers
* and keywords. That means we can use variable names like "foo_bar". Other programming languages even allow "!" and "?" in identifiers. If you want to allow that too, this is the place to sneak it in
 */
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

/**
* This helper function is found in a lot of parser. Sometimes it's called "eatWhitespace"
* and sometimes "consumeWhitespace" and sometimes something entirely different.
* Which characters these functions actually skip depends on the language being lexed.
* Some language implementations do create tokens for newline characters for example and throw
* parsing errors if they are not at the correct place in the stream of tokens.
* We skip over newline characters to make the parsing step later on a litter easier
 */
func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

/**
* Note here that we simplified things by a lot. We don't read floats, hex, or octal notation
* We are only reading integers
 */
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
