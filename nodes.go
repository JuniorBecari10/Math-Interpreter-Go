package main

type Node interface {
  node()
}

type NumberNode struct {
  value float64
}

type BinNode struct {
  nodeA Node
  nodeB Node
  op string
}

type PlusNode struct {
  value Node
}

type MinusNode struct {
  value Node
}

func (n NumberNode) node() {}
func (n BinNode)    node() {}
func (n PlusNode)   node() {}
func (n MinusNode)  node() {}