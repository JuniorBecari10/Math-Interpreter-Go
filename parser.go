package main

type Parser struct {
  tokens []Token
  cursor int
}

func NewParser(tokens []Token) *Parser {
  return &Parser { tokens: tokens }
}

func Parse(p *Parser) Node {
  return p.exp()
}

// ---

func (p *Parser) advance() {
  p.cursor++
}

func (p *Parser) getToken() Token {
  if p.cursor < len(p.tokens) {
    return p.tokens[p.cursor]
  }
  
  return Token { ERROR, "", p.cursor }
}

func (p *Parser) exp() Node {
  if p.getToken().kind == ERROR {
    return nil
  }
  
  res := p.term()
  
  for p.getToken().kind != ERROR && (p.getToken().kind == PLUS || p.getToken().kind == MINUS) {
    if p.getToken().kind == PLUS {
      p.advance()
      
      res = BinNode { res, p.term(), "+" }
    } else if p.getToken().kind == MINUS {
      p.advance()
      
      res = BinNode { res, p.term(), "-" }
    }
  }
  
  return res
}

func (p *Parser) term() Node {
  if p.getToken().kind == ERROR {
    return nil
  }
  
  res := p.factor()
  
  for p.getToken().kind != ERROR && (p.getToken().kind == TIMES || p.getToken().kind == DIVIDE) {
    if p.getToken().kind == TIMES {
      p.advance()
     
      res = BinNode { res, p.term(), "*" }
    } else if p.getToken().kind == DIVIDE {
      p.advance()
      
      res = BinNode { res, p.term(), "/" }
    }
  }
  
  return res
}

func (p *Parser) factor() Node {
  tk := p.getToken()
  
  if tk.kind == LPAREN {
    p.advance()
    res := p.exp()
    
    if tk.kind != RPAREN {
      PrintError(p.cursor, "Unclosed parenthesis.")
      //return nil
    }
    
    p.advance()
    return res
  }
  
  if tk.kind == RPAREN {
    PrintError(p.cursor, "Unopened parenthesis.")
  }
  
  if tk.kind == PLUS {
    p.advance()
    
    return PlusNode { p.factor() }
  }
  
  if tk.kind == MINUS {
    p.advance()
    
    return MinusNode { p.factor() }
  }
  
  if tk.kind == NUMBER {
    p.advance()
    
    vl, _ := tk.value.(float64)
    
    return NumberNode { vl }
  }
  
  return nil
}