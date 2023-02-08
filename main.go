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
    
    fmt.Println(tks)
  }
}