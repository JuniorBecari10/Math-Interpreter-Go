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
    
    l := NewLexer(text)
    tks := Lex(l)
    
    p := NewParser(tks)
    ast := Parse(p)
    
    fmt.Println(ast)
  }
}