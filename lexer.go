package main

import (
  "fmt"
  "strconv"
)

const (
  WHITESPACE = " \r\t"
  NUMBERS    = "0123456789"
)

type Lexer struct {
  text   string
  cursor int
}

func NewLexer(text string) *Lexer {
  return &Lexer { text: text, cursor: 0 }
}

func Lex(l *Lexer) []Token {
  tks := []Token {}
  
  tk := l.nextToken()
  for tk.kind != END && tk.kind != ERROR {
    tks = append(tks, tk)
    
    tk = l.nextToken()
  }
  
  if tk.kind == ERROR {
    return []Token {}
  }
  
  return tks
}

// ---

func (l *Lexer) advance() {
  l.cursor++
}

func (l *Lexer) getChar() uint8 {
  if l.cursor < len(l.text) {
    return l.text[l.cursor]
  }
  
  return 0
}

func (l *Lexer) nextToken() Token {
  for IsWhitespace(l.getChar()) {
    l.advance()
  }
  
  if l.getChar() == 0 {
    return Token { END, "", l.cursor }
  }
  
  if l.getChar() == '+' {
    pos := l.cursor
    l.advance()
    
    return Token { PLUS, "+", pos }
  }
  
  if l.getChar() == '-' {
    pos := l.cursor
    l.advance()
    
    return Token { MINUS, "-", pos }
  }
  
  if l.getChar() == '*' {
    pos := l.cursor
    l.advance()
    
    return Token { TIMES, "*", pos }
  }
  
  if l.getChar() == '/' {
    pos := l.cursor
    l.advance()
    
    return Token { DIVIDE, "/", pos }
  }
  
  if l.getChar() == '(' {
    pos := l.cursor
    l.advance()
    
    return Token { LPAREN, "(", pos }
  }
  
  if l.getChar() == ')' {
    pos := l.cursor
    l.advance()
    
    return Token { RPAREN, ")", pos }
  }
  
  if IsNumber(l.getChar()) || l.getChar() == '.' {
    pos := l.cursor
    
    nStr := l.getNumber()
    
    n, err := strconv.ParseFloat(nStr, 64)
    
    if err != nil {
      PrintError(pos, "Cannot parse token as a number.")
      return Token { ERROR, "", l.cursor }
    }
    
    return Token { NUMBER, n, pos }
  }
  
  PrintError(l.cursor, "Unknown token.")
  return Token { ERROR, "", l.cursor }
}

func (l *Lexer) getNumber() string {
  pos := l.cursor
  
  for IsNumber(l.getChar()) || l.getChar() == '.' {
    l.advance()
  }
  
  return l.text[pos:l.cursor]
}

// ---

func IsWhitespace(c uint8) bool {
  for _, ch := range WHITESPACE {
    if c == uint8(ch) {
      return true
    }
  }
  
  return false
}

func IsNumber(c uint8) bool {
  for _, ch := range NUMBERS {
    if c == uint8(ch) {
      return true
    }
  }
  
  return false
}

func PrintError(pos int, msg string) {
  fmt.Print("  ")
  
  for i := 0; i < pos; i++ {
    fmt.Print(" ")
  }
  
  fmt.Println("^")
  
  fmt.Println(msg)
}