package main

import (
  "bytes"
)

const (
  NUMBER = "NUMBER"
  
  PLUS   = "PLUS"
  MINUS  = "MINUS"
  TIMES  = "TIMES"
  DIVIDE = "DIVIDE"
  
  LPAREN = "LPAREN"
  RPAREN = "RPAREN"
  
  END    = "END"
  ERROR  = "ERROR"
)

type Token struct {
  kind  string
  chars string
  pos   int
}

func (t *Token) String() string {
  var out bytes.Buffer
  
  out.WriteString(t.kind)
  
  if t.kind == NUMBER {
    out.WriteString(":" + t.chars)
  }
  
  return out.String()
}