package main

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

type Any interface {}

type Token struct {
  kind  string
  chars Any
  pos   int
}