package main

import (
  "fmt"
)

type Node interface {
  node() {}
}

type NumberNode struct {
  value float64
}

func (n *NumberNode) node() {}