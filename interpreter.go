package main

import (
  "fmt"
)

func Interpret(n Node) Number {
  fn := GetFunc(n)
  
  return fn(n)
}

func GetFunc(n Node) func(n Node) Number {
  switch n.(type) {
    case NumberNode:
      return VisitNumber
    
    case BinNode:
      return VisitBin
    
    case PlusNode:
      return VisitPlus
    
    case MinusNode:
      return VisitMinus
    
    default:
      fmt.Printf("Couldn't get type %T.\n", n)
      return nil
  }
}

// ---

func VisitNumber(n Node) Number {
  return Number { n.(NumberNode).value }
}

func VisitBin(n Node) Number {
  asBin := n.(BinNode)
  
  switch asBin.op {
    case "+":
      return Number { Interpret(asBin.nodeA).value + Interpret(asBin.nodeB).value }
    
    case "-":
      return Number { Interpret(asBin.nodeA).value - Interpret(asBin.nodeB).value }
    
    case "*":
      return Number { Interpret(asBin.nodeA).value * Interpret(asBin.nodeB).value }
    
    case "/":
      valueB := Interpret(asBin.nodeB).value
      
      if valueB == 0 {
        PrintError(0, "Division by zero.")
      }
      
      return Number { Interpret(asBin.nodeA).value / valueB }
    
  }
  
  return Number {}
}

func VisitPlus(n Node) Number {
  return Number { n.(NumberNode).value }
}

func VisitMinus(n Node) Number {
  return Number { -n.(NumberNode).value }
}