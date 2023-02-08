package main

import (
  "fmt"
  "bufio"
  "os"
)

func main() {
  scanner := bufio.NewScanner(os.Stdin)
  
  for {
    fmt.Printf("> ")
    scanner.Scan()
    
    text := scanner.Text()
    
    if text == "" {
      continue
    }
    
    l := NewLexer(text)
    tks := Lex(l)
    
    if tks == nil || len(tks) == 0 {
      continue
    }
    
    p := NewParser(tks)
    ast := Parse(p)
    
    if ast == nil {
      continue
    }
    
    value := Interpret(ast)
    
    fmt.Printf("< %f\n", value.value)
  }
}