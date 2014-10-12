package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

const end_symbol rune = 4

/* The rule types inferred from the grammar are below. */
type pegRule uint8

const (
	ruleUnknown pegRule = iota
	ruleStart
	ruleInterfaceDef
	ruleInterfaceBody
	ruleConstantDef
	ruleAttributeDef
	ruleFunctionDef
	ruleFunctionArgs
	ruleFunctionArg
	ruleTypeDef
	ruleToken
	ruleNumber
	ruleEndOfLine
	ruleSpace
	ruleComment
	ruleSpaceComment
	ruleSpacing
	ruleMustSpacing
	ruleAction0
	ruleAction1
	ruleAction2
	ruleAction3
	ruleAction4
	ruleAction5
	ruleAction6
	ruleAction7
	ruleAction8
	ruleAction9
	ruleAction10
	ruleAction11
	ruleAction12
	ruleAction13
	ruleAction14
	ruleAction15
	ruleAction16
	ruleAction17
	ruleAction18
	ruleAction19
	ruleAction20
	ruleAction21
	ruleAction22
	ruleAction23
	ruleAction24
	ruleAction25
	rulePegText
	ruleAction26
	ruleAction27
	ruleAction28

	rulePre_
	rule_In_
	rule_Suf
)

var rul3s = [...]string{
	"Unknown",
	"Start",
	"InterfaceDef",
	"InterfaceBody",
	"ConstantDef",
	"AttributeDef",
	"FunctionDef",
	"FunctionArgs",
	"FunctionArg",
	"TypeDef",
	"Token",
	"Number",
	"EndOfLine",
	"Space",
	"Comment",
	"SpaceComment",
	"Spacing",
	"MustSpacing",
	"Action0",
	"Action1",
	"Action2",
	"Action3",
	"Action4",
	"Action5",
	"Action6",
	"Action7",
	"Action8",
	"Action9",
	"Action10",
	"Action11",
	"Action12",
	"Action13",
	"Action14",
	"Action15",
	"Action16",
	"Action17",
	"Action18",
	"Action19",
	"Action20",
	"Action21",
	"Action22",
	"Action23",
	"Action24",
	"Action25",
	"PegText",
	"Action26",
	"Action27",
	"Action28",

	"Pre_",
	"_In_",
	"_Suf",
}

type tokenTree interface {
	Print()
	PrintSyntax()
	PrintSyntaxTree(buffer string)
	Add(rule pegRule, begin, end, next, depth int)
	Expand(index int) tokenTree
	Tokens() <-chan token32
	AST() *node32
	Error() []token32
	trim(length int)
}

type node32 struct {
	token32
	up, next *node32
}

func (node *node32) print(depth int, buffer string) {
	for node != nil {
		for c := 0; c < depth; c++ {
			fmt.Printf(" ")
		}
		fmt.Printf("\x1B[34m%v\x1B[m %v\n", rul3s[node.pegRule], strconv.Quote(buffer[node.begin:node.end]))
		if node.up != nil {
			node.up.print(depth+1, buffer)
		}
		node = node.next
	}
}

func (ast *node32) Print(buffer string) {
	ast.print(0, buffer)
}

type element struct {
	node *node32
	down *element
}

/* ${@} bit structure for abstract syntax tree */
type token16 struct {
	pegRule
	begin, end, next int16
}

func (t *token16) isZero() bool {
	return t.pegRule == ruleUnknown && t.begin == 0 && t.end == 0 && t.next == 0
}

func (t *token16) isParentOf(u token16) bool {
	return t.begin <= u.begin && t.end >= u.end && t.next > u.next
}

func (t *token16) getToken32() token32 {
	return token32{pegRule: t.pegRule, begin: int32(t.begin), end: int32(t.end), next: int32(t.next)}
}

func (t *token16) String() string {
	return fmt.Sprintf("\x1B[34m%v\x1B[m %v %v %v", rul3s[t.pegRule], t.begin, t.end, t.next)
}

type tokens16 struct {
	tree    []token16
	ordered [][]token16
}

func (t *tokens16) trim(length int) {
	t.tree = t.tree[0:length]
}

func (t *tokens16) Print() {
	for _, token := range t.tree {
		fmt.Println(token.String())
	}
}

func (t *tokens16) Order() [][]token16 {
	if t.ordered != nil {
		return t.ordered
	}

	depths := make([]int16, 1, math.MaxInt16)
	for i, token := range t.tree {
		if token.pegRule == ruleUnknown {
			t.tree = t.tree[:i]
			break
		}
		depth := int(token.next)
		if length := len(depths); depth >= length {
			depths = depths[:depth+1]
		}
		depths[depth]++
	}
	depths = append(depths, 0)

	ordered, pool := make([][]token16, len(depths)), make([]token16, len(t.tree)+len(depths))
	for i, depth := range depths {
		depth++
		ordered[i], pool, depths[i] = pool[:depth], pool[depth:], 0
	}

	for i, token := range t.tree {
		depth := token.next
		token.next = int16(i)
		ordered[depth][depths[depth]] = token
		depths[depth]++
	}
	t.ordered = ordered
	return ordered
}

type state16 struct {
	token16
	depths []int16
	leaf   bool
}

func (t *tokens16) AST() *node32 {
	tokens := t.Tokens()
	stack := &element{node: &node32{token32: <-tokens}}
	for token := range tokens {
		if token.begin == token.end {
			continue
		}
		node := &node32{token32: token}
		for stack != nil && stack.node.begin >= token.begin && stack.node.end <= token.end {
			stack.node.next = node.up
			node.up = stack.node
			stack = stack.down
		}
		stack = &element{node: node, down: stack}
	}
	return stack.node
}

func (t *tokens16) PreOrder() (<-chan state16, [][]token16) {
	s, ordered := make(chan state16, 6), t.Order()
	go func() {
		var states [8]state16
		for i, _ := range states {
			states[i].depths = make([]int16, len(ordered))
		}
		depths, state, depth := make([]int16, len(ordered)), 0, 1
		write := func(t token16, leaf bool) {
			S := states[state]
			state, S.pegRule, S.begin, S.end, S.next, S.leaf = (state+1)%8, t.pegRule, t.begin, t.end, int16(depth), leaf
			copy(S.depths, depths)
			s <- S
		}

		states[state].token16 = ordered[0][0]
		depths[0]++
		state++
		a, b := ordered[depth-1][depths[depth-1]-1], ordered[depth][depths[depth]]
	depthFirstSearch:
		for {
			for {
				if i := depths[depth]; i > 0 {
					if c, j := ordered[depth][i-1], depths[depth-1]; a.isParentOf(c) &&
						(j < 2 || !ordered[depth-1][j-2].isParentOf(c)) {
						if c.end != b.begin {
							write(token16{pegRule: rule_In_, begin: c.end, end: b.begin}, true)
						}
						break
					}
				}

				if a.begin < b.begin {
					write(token16{pegRule: rulePre_, begin: a.begin, end: b.begin}, true)
				}
				break
			}

			next := depth + 1
			if c := ordered[next][depths[next]]; c.pegRule != ruleUnknown && b.isParentOf(c) {
				write(b, false)
				depths[depth]++
				depth, a, b = next, b, c
				continue
			}

			write(b, true)
			depths[depth]++
			c, parent := ordered[depth][depths[depth]], true
			for {
				if c.pegRule != ruleUnknown && a.isParentOf(c) {
					b = c
					continue depthFirstSearch
				} else if parent && b.end != a.end {
					write(token16{pegRule: rule_Suf, begin: b.end, end: a.end}, true)
				}

				depth--
				if depth > 0 {
					a, b, c = ordered[depth-1][depths[depth-1]-1], a, ordered[depth][depths[depth]]
					parent = a.isParentOf(b)
					continue
				}

				break depthFirstSearch
			}
		}

		close(s)
	}()
	return s, ordered
}

func (t *tokens16) PrintSyntax() {
	tokens, ordered := t.PreOrder()
	max := -1
	for token := range tokens {
		if !token.leaf {
			fmt.Printf("%v", token.begin)
			for i, leaf, depths := 0, int(token.next), token.depths; i < leaf; i++ {
				fmt.Printf(" \x1B[36m%v\x1B[m", rul3s[ordered[i][depths[i]-1].pegRule])
			}
			fmt.Printf(" \x1B[36m%v\x1B[m\n", rul3s[token.pegRule])
		} else if token.begin == token.end {
			fmt.Printf("%v", token.begin)
			for i, leaf, depths := 0, int(token.next), token.depths; i < leaf; i++ {
				fmt.Printf(" \x1B[31m%v\x1B[m", rul3s[ordered[i][depths[i]-1].pegRule])
			}
			fmt.Printf(" \x1B[31m%v\x1B[m\n", rul3s[token.pegRule])
		} else {
			for c, end := token.begin, token.end; c < end; c++ {
				if i := int(c); max+1 < i {
					for j := max; j < i; j++ {
						fmt.Printf("skip %v %v\n", j, token.String())
					}
					max = i
				} else if i := int(c); i <= max {
					for j := i; j <= max; j++ {
						fmt.Printf("dupe %v %v\n", j, token.String())
					}
				} else {
					max = int(c)
				}
				fmt.Printf("%v", c)
				for i, leaf, depths := 0, int(token.next), token.depths; i < leaf; i++ {
					fmt.Printf(" \x1B[34m%v\x1B[m", rul3s[ordered[i][depths[i]-1].pegRule])
				}
				fmt.Printf(" \x1B[34m%v\x1B[m\n", rul3s[token.pegRule])
			}
			fmt.Printf("\n")
		}
	}
}

func (t *tokens16) PrintSyntaxTree(buffer string) {
	tokens, _ := t.PreOrder()
	for token := range tokens {
		for c := 0; c < int(token.next); c++ {
			fmt.Printf(" ")
		}
		fmt.Printf("\x1B[34m%v\x1B[m %v\n", rul3s[token.pegRule], strconv.Quote(buffer[token.begin:token.end]))
	}
}

func (t *tokens16) Add(rule pegRule, begin, end, depth, index int) {
	t.tree[index] = token16{pegRule: rule, begin: int16(begin), end: int16(end), next: int16(depth)}
}

func (t *tokens16) Tokens() <-chan token32 {
	s := make(chan token32, 16)
	go func() {
		for _, v := range t.tree {
			s <- v.getToken32()
		}
		close(s)
	}()
	return s
}

func (t *tokens16) Error() []token32 {
	ordered := t.Order()
	length := len(ordered)
	tokens, length := make([]token32, length), length-1
	for i, _ := range tokens {
		o := ordered[length-i]
		if len(o) > 1 {
			tokens[i] = o[len(o)-2].getToken32()
		}
	}
	return tokens
}

/* ${@} bit structure for abstract syntax tree */
type token32 struct {
	pegRule
	begin, end, next int32
}

func (t *token32) isZero() bool {
	return t.pegRule == ruleUnknown && t.begin == 0 && t.end == 0 && t.next == 0
}

func (t *token32) isParentOf(u token32) bool {
	return t.begin <= u.begin && t.end >= u.end && t.next > u.next
}

func (t *token32) getToken32() token32 {
	return token32{pegRule: t.pegRule, begin: int32(t.begin), end: int32(t.end), next: int32(t.next)}
}

func (t *token32) String() string {
	return fmt.Sprintf("\x1B[34m%v\x1B[m %v %v %v", rul3s[t.pegRule], t.begin, t.end, t.next)
}

type tokens32 struct {
	tree    []token32
	ordered [][]token32
}

func (t *tokens32) trim(length int) {
	t.tree = t.tree[0:length]
}

func (t *tokens32) Print() {
	for _, token := range t.tree {
		fmt.Println(token.String())
	}
}

func (t *tokens32) Order() [][]token32 {
	if t.ordered != nil {
		return t.ordered
	}

	depths := make([]int32, 1, math.MaxInt16)
	for i, token := range t.tree {
		if token.pegRule == ruleUnknown {
			t.tree = t.tree[:i]
			break
		}
		depth := int(token.next)
		if length := len(depths); depth >= length {
			depths = depths[:depth+1]
		}
		depths[depth]++
	}
	depths = append(depths, 0)

	ordered, pool := make([][]token32, len(depths)), make([]token32, len(t.tree)+len(depths))
	for i, depth := range depths {
		depth++
		ordered[i], pool, depths[i] = pool[:depth], pool[depth:], 0
	}

	for i, token := range t.tree {
		depth := token.next
		token.next = int32(i)
		ordered[depth][depths[depth]] = token
		depths[depth]++
	}
	t.ordered = ordered
	return ordered
}

type state32 struct {
	token32
	depths []int32
	leaf   bool
}

func (t *tokens32) AST() *node32 {
	tokens := t.Tokens()
	stack := &element{node: &node32{token32: <-tokens}}
	for token := range tokens {
		if token.begin == token.end {
			continue
		}
		node := &node32{token32: token}
		for stack != nil && stack.node.begin >= token.begin && stack.node.end <= token.end {
			stack.node.next = node.up
			node.up = stack.node
			stack = stack.down
		}
		stack = &element{node: node, down: stack}
	}
	return stack.node
}

func (t *tokens32) PreOrder() (<-chan state32, [][]token32) {
	s, ordered := make(chan state32, 6), t.Order()
	go func() {
		var states [8]state32
		for i, _ := range states {
			states[i].depths = make([]int32, len(ordered))
		}
		depths, state, depth := make([]int32, len(ordered)), 0, 1
		write := func(t token32, leaf bool) {
			S := states[state]
			state, S.pegRule, S.begin, S.end, S.next, S.leaf = (state+1)%8, t.pegRule, t.begin, t.end, int32(depth), leaf
			copy(S.depths, depths)
			s <- S
		}

		states[state].token32 = ordered[0][0]
		depths[0]++
		state++
		a, b := ordered[depth-1][depths[depth-1]-1], ordered[depth][depths[depth]]
	depthFirstSearch:
		for {
			for {
				if i := depths[depth]; i > 0 {
					if c, j := ordered[depth][i-1], depths[depth-1]; a.isParentOf(c) &&
						(j < 2 || !ordered[depth-1][j-2].isParentOf(c)) {
						if c.end != b.begin {
							write(token32{pegRule: rule_In_, begin: c.end, end: b.begin}, true)
						}
						break
					}
				}

				if a.begin < b.begin {
					write(token32{pegRule: rulePre_, begin: a.begin, end: b.begin}, true)
				}
				break
			}

			next := depth + 1
			if c := ordered[next][depths[next]]; c.pegRule != ruleUnknown && b.isParentOf(c) {
				write(b, false)
				depths[depth]++
				depth, a, b = next, b, c
				continue
			}

			write(b, true)
			depths[depth]++
			c, parent := ordered[depth][depths[depth]], true
			for {
				if c.pegRule != ruleUnknown && a.isParentOf(c) {
					b = c
					continue depthFirstSearch
				} else if parent && b.end != a.end {
					write(token32{pegRule: rule_Suf, begin: b.end, end: a.end}, true)
				}

				depth--
				if depth > 0 {
					a, b, c = ordered[depth-1][depths[depth-1]-1], a, ordered[depth][depths[depth]]
					parent = a.isParentOf(b)
					continue
				}

				break depthFirstSearch
			}
		}

		close(s)
	}()
	return s, ordered
}

func (t *tokens32) PrintSyntax() {
	tokens, ordered := t.PreOrder()
	max := -1
	for token := range tokens {
		if !token.leaf {
			fmt.Printf("%v", token.begin)
			for i, leaf, depths := 0, int(token.next), token.depths; i < leaf; i++ {
				fmt.Printf(" \x1B[36m%v\x1B[m", rul3s[ordered[i][depths[i]-1].pegRule])
			}
			fmt.Printf(" \x1B[36m%v\x1B[m\n", rul3s[token.pegRule])
		} else if token.begin == token.end {
			fmt.Printf("%v", token.begin)
			for i, leaf, depths := 0, int(token.next), token.depths; i < leaf; i++ {
				fmt.Printf(" \x1B[31m%v\x1B[m", rul3s[ordered[i][depths[i]-1].pegRule])
			}
			fmt.Printf(" \x1B[31m%v\x1B[m\n", rul3s[token.pegRule])
		} else {
			for c, end := token.begin, token.end; c < end; c++ {
				if i := int(c); max+1 < i {
					for j := max; j < i; j++ {
						fmt.Printf("skip %v %v\n", j, token.String())
					}
					max = i
				} else if i := int(c); i <= max {
					for j := i; j <= max; j++ {
						fmt.Printf("dupe %v %v\n", j, token.String())
					}
				} else {
					max = int(c)
				}
				fmt.Printf("%v", c)
				for i, leaf, depths := 0, int(token.next), token.depths; i < leaf; i++ {
					fmt.Printf(" \x1B[34m%v\x1B[m", rul3s[ordered[i][depths[i]-1].pegRule])
				}
				fmt.Printf(" \x1B[34m%v\x1B[m\n", rul3s[token.pegRule])
			}
			fmt.Printf("\n")
		}
	}
}

func (t *tokens32) PrintSyntaxTree(buffer string) {
	tokens, _ := t.PreOrder()
	for token := range tokens {
		for c := 0; c < int(token.next); c++ {
			fmt.Printf(" ")
		}
		fmt.Printf("\x1B[34m%v\x1B[m %v\n", rul3s[token.pegRule], strconv.Quote(buffer[token.begin:token.end]))
	}
}

func (t *tokens32) Add(rule pegRule, begin, end, depth, index int) {
	t.tree[index] = token32{pegRule: rule, begin: int32(begin), end: int32(end), next: int32(depth)}
}

func (t *tokens32) Tokens() <-chan token32 {
	s := make(chan token32, 16)
	go func() {
		for _, v := range t.tree {
			s <- v.getToken32()
		}
		close(s)
	}()
	return s
}

func (t *tokens32) Error() []token32 {
	ordered := t.Order()
	length := len(ordered)
	tokens, length := make([]token32, length), length-1
	for i, _ := range tokens {
		o := ordered[length-i]
		if len(o) > 1 {
			tokens[i] = o[len(o)-2].getToken32()
		}
	}
	return tokens
}

func (t *tokens16) Expand(index int) tokenTree {
	tree := t.tree
	if index >= len(tree) {
		expanded := make([]token32, 2*len(tree))
		for i, v := range tree {
			expanded[i] = v.getToken32()
		}
		return &tokens32{tree: expanded}
	}
	return nil
}

func (t *tokens32) Expand(index int) tokenTree {
	tree := t.tree
	if index >= len(tree) {
		expanded := make([]token32, 2*len(tree))
		copy(expanded, tree)
		t.tree = expanded
	}
	return nil
}

type Parser struct {
	idl            *IDL
	currentWord    string
	currentNumber  string
	currentIntDef  *InterfaceDef
	currentAttDef  *AttributeDef
	currentConDef  *ConstantDef
	currentFunDef  *FunctionDef
	currentArgDef  *ArgumentDef
	currentTypDef  *TypeDef
	currentArgList []ArgumentDef

	Buffer string
	buffer []rune
	rules  [48]func() bool
	Parse  func(rule ...int) error
	Reset  func()
	tokenTree
}

type textPosition struct {
	line, symbol int
}

type textPositionMap map[int]textPosition

func translatePositions(buffer string, positions []int) textPositionMap {
	length, translations, j, line, symbol := len(positions), make(textPositionMap, len(positions)), 0, 1, 0
	sort.Ints(positions)

search:
	for i, c := range buffer[0:] {
		if c == '\n' {
			line, symbol = line+1, 0
		} else {
			symbol++
		}
		if i == positions[j] {
			translations[positions[j]] = textPosition{line, symbol}
			for j++; j < length; j++ {
				if i != positions[j] {
					continue search
				}
			}
			break search
		}
	}

	return translations
}

type parseError struct {
	p *Parser
}

func (e *parseError) Error() string {
	tokens, error := e.p.tokenTree.Error(), "\n"
	positions, p := make([]int, 2*len(tokens)), 0
	for _, token := range tokens {
		positions[p], p = int(token.begin), p+1
		positions[p], p = int(token.end), p+1
	}
	translations := translatePositions(e.p.Buffer, positions)
	for _, token := range tokens {
		begin, end := int(token.begin), int(token.end)
		error += fmt.Sprintf("parse error near \x1B[34m%v\x1B[m (line %v symbol %v - line %v symbol %v):\n%v\n",
			rul3s[token.pegRule],
			translations[begin].line, translations[begin].symbol,
			translations[end].line, translations[end].symbol,
			/*strconv.Quote(*/ e.p.Buffer[begin:end] /*)*/)
	}

	return error
}

func (p *Parser) PrintSyntaxTree() {
	p.tokenTree.PrintSyntaxTree(p.Buffer)
}

func (p *Parser) Highlighter() {
	p.tokenTree.PrintSyntax()
}

func (p *Parser) Execute() {
	buffer, begin, end := p.Buffer, 0, 0
	for token := range p.tokenTree.Tokens() {
		switch token.pegRule {
		case rulePegText:
			begin, end = int(token.begin), int(token.end)
		case ruleAction0:
			p.currentIntDef = new(InterfaceDef)
		case ruleAction1:
			p.currentIntDef.Name = p.currentWord
		case ruleAction2:
			p.currentIntDef.Inherits = p.currentWord
		case ruleAction3:
			p.idl.Interfaces = append(p.idl.Interfaces, *p.currentIntDef)
			p.currentIntDef = nil
		case ruleAction4:
			p.currentIntDef.Constants = append(p.currentIntDef.Constants, *p.currentConDef)
			p.currentConDef = nil
		case ruleAction5:
			p.currentIntDef.Attributes = append(p.currentIntDef.Attributes, *p.currentAttDef)
			p.currentAttDef = nil
		case ruleAction6:
			p.currentIntDef.Functions = append(p.currentIntDef.Functions, *p.currentFunDef)
			p.currentFunDef = nil
		case ruleAction7:
			p.currentConDef = new(ConstantDef)
		case ruleAction8:
			p.currentConDef.Type = *p.currentTypDef
			p.currentTypDef = nil
		case ruleAction9:
			p.currentConDef.Name = p.currentWord
		case ruleAction10:
			p.currentConDef.Value = p.currentNumber
		case ruleAction11:
			p.currentAttDef = new(AttributeDef)
		case ruleAction12:
			p.currentAttDef.Type = *p.currentTypDef
			p.currentTypDef = nil
		case ruleAction13:
			p.currentAttDef.Name = p.currentWord
		case ruleAction14:
			p.currentFunDef = new(FunctionDef)
		case ruleAction15:
			p.currentFunDef.Type = *p.currentTypDef
			p.currentTypDef = nil
		case ruleAction16:
			p.currentFunDef.Name = p.currentWord
		case ruleAction17:
			p.currentFunDef.Arguments = p.currentArgList
			p.currentArgList = nil
		case ruleAction18:
			p.currentArgList = append(p.currentArgList, *p.currentArgDef)
			p.currentArgDef = nil
		case ruleAction19:
			p.currentArgList = append(p.currentArgList, *p.currentArgDef)
			p.currentArgDef = nil
		case ruleAction20:
			p.currentArgDef = new(ArgumentDef)
		case ruleAction21:
			p.currentArgDef.Type = *p.currentTypDef
			p.currentTypDef = nil
		case ruleAction22:
			p.currentArgDef.Name = p.currentWord
		case ruleAction23:
			p.currentTypDef = new(TypeDef)
		case ruleAction24:
			p.currentTypDef.Unsigned = true
		case ruleAction25:
			p.currentTypDef.Type = p.currentWord
		case ruleAction26:
			p.currentWord = string(buffer[begin:end])
		case ruleAction27:
			p.currentNumber = "0x" + string(buffer[begin:end])
		case ruleAction28:
			p.currentNumber = string(buffer[begin:end])

		}
	}
}

func (p *Parser) Init() {
	p.buffer = []rune(p.Buffer)
	if len(p.buffer) == 0 || p.buffer[len(p.buffer)-1] != end_symbol {
		p.buffer = append(p.buffer, end_symbol)
	}

	var tree tokenTree = &tokens16{tree: make([]token16, math.MaxInt16)}
	position, depth, tokenIndex, buffer, _rules := 0, 0, 0, p.buffer, p.rules

	p.Parse = func(rule ...int) error {
		r := 1
		if len(rule) > 0 {
			r = rule[0]
		}
		matches := p.rules[r]()
		p.tokenTree = tree
		if matches {
			p.tokenTree.trim(tokenIndex)
			return nil
		}
		return &parseError{p}
	}

	p.Reset = func() {
		position, tokenIndex, depth = 0, 0, 0
	}

	add := func(rule pegRule, begin int) {
		if t := tree.Expand(tokenIndex); t != nil {
			tree = t
		}
		tree.Add(rule, begin, position, depth, tokenIndex)
		tokenIndex++
	}

	matchDot := func() bool {
		if buffer[position] != end_symbol {
			position++
			return true
		}
		return false
	}

	/*matchChar := func(c byte) bool {
		if buffer[position] == c {
			position++
			return true
		}
		return false
	}*/

	/*matchRange := func(lower byte, upper byte) bool {
		if c := buffer[position]; c >= lower && c <= upper {
			position++
			return true
		}
		return false
	}*/

	_rules = [...]func() bool{
		nil,
		/* 0 Start <- <(Spacing (InterfaceDef Spacing)*)> */
		func() bool {
			position0, tokenIndex0, depth0 := position, tokenIndex, depth
			{
				position1 := position
				depth++
				if !_rules[ruleSpacing]() {
					goto l0
				}
			l2:
				{
					position3, tokenIndex3, depth3 := position, tokenIndex, depth
					if !_rules[ruleInterfaceDef]() {
						goto l3
					}
					if !_rules[ruleSpacing]() {
						goto l3
					}
					goto l2
				l3:
					position, tokenIndex, depth = position3, tokenIndex3, depth3
				}
				depth--
				add(ruleStart, position1)
			}
			return true
		l0:
			position, tokenIndex, depth = position0, tokenIndex0, depth0
			return false
		},
		/* 1 InterfaceDef <- <('i' 'n' 't' 'e' 'r' 'f' 'a' 'c' 'e' MustSpacing Action0 Token Spacing Action1 (':' Spacing Token Spacing Action2)? InterfaceBody Spacing ';' Action3)> */
		func() bool {
			position4, tokenIndex4, depth4 := position, tokenIndex, depth
			{
				position5 := position
				depth++
				if buffer[position] != rune('i') {
					goto l4
				}
				position++
				if buffer[position] != rune('n') {
					goto l4
				}
				position++
				if buffer[position] != rune('t') {
					goto l4
				}
				position++
				if buffer[position] != rune('e') {
					goto l4
				}
				position++
				if buffer[position] != rune('r') {
					goto l4
				}
				position++
				if buffer[position] != rune('f') {
					goto l4
				}
				position++
				if buffer[position] != rune('a') {
					goto l4
				}
				position++
				if buffer[position] != rune('c') {
					goto l4
				}
				position++
				if buffer[position] != rune('e') {
					goto l4
				}
				position++
				if !_rules[ruleMustSpacing]() {
					goto l4
				}
				if !_rules[ruleAction0]() {
					goto l4
				}
				if !_rules[ruleToken]() {
					goto l4
				}
				if !_rules[ruleSpacing]() {
					goto l4
				}
				if !_rules[ruleAction1]() {
					goto l4
				}
				{
					position6, tokenIndex6, depth6 := position, tokenIndex, depth
					if buffer[position] != rune(':') {
						goto l6
					}
					position++
					if !_rules[ruleSpacing]() {
						goto l6
					}
					if !_rules[ruleToken]() {
						goto l6
					}
					if !_rules[ruleSpacing]() {
						goto l6
					}
					if !_rules[ruleAction2]() {
						goto l6
					}
					goto l7
				l6:
					position, tokenIndex, depth = position6, tokenIndex6, depth6
				}
			l7:
				if !_rules[ruleInterfaceBody]() {
					goto l4
				}
				if !_rules[ruleSpacing]() {
					goto l4
				}
				if buffer[position] != rune(';') {
					goto l4
				}
				position++
				if !_rules[ruleAction3]() {
					goto l4
				}
				depth--
				add(ruleInterfaceDef, position5)
			}
			return true
		l4:
			position, tokenIndex, depth = position4, tokenIndex4, depth4
			return false
		},
		/* 2 InterfaceBody <- <('{' Spacing (((ConstantDef Action4) / (AttributeDef Action5) / (FunctionDef Action6)) Spacing)* '}')> */
		func() bool {
			position8, tokenIndex8, depth8 := position, tokenIndex, depth
			{
				position9 := position
				depth++
				if buffer[position] != rune('{') {
					goto l8
				}
				position++
				if !_rules[ruleSpacing]() {
					goto l8
				}
			l10:
				{
					position11, tokenIndex11, depth11 := position, tokenIndex, depth
					{
						position12, tokenIndex12, depth12 := position, tokenIndex, depth
						if !_rules[ruleConstantDef]() {
							goto l13
						}
						if !_rules[ruleAction4]() {
							goto l13
						}
						goto l12
					l13:
						position, tokenIndex, depth = position12, tokenIndex12, depth12
						if !_rules[ruleAttributeDef]() {
							goto l14
						}
						if !_rules[ruleAction5]() {
							goto l14
						}
						goto l12
					l14:
						position, tokenIndex, depth = position12, tokenIndex12, depth12
						if !_rules[ruleFunctionDef]() {
							goto l11
						}
						if !_rules[ruleAction6]() {
							goto l11
						}
					}
				l12:
					if !_rules[ruleSpacing]() {
						goto l11
					}
					goto l10
				l11:
					position, tokenIndex, depth = position11, tokenIndex11, depth11
				}
				if buffer[position] != rune('}') {
					goto l8
				}
				position++
				depth--
				add(ruleInterfaceBody, position9)
			}
			return true
		l8:
			position, tokenIndex, depth = position8, tokenIndex8, depth8
			return false
		},
		/* 3 ConstantDef <- <('c' 'o' 'n' 's' 't' MustSpacing Action7 TypeDef MustSpacing Action8 Token Spacing Action9 '=' Spacing Number Action10 ';')> */
		func() bool {
			position15, tokenIndex15, depth15 := position, tokenIndex, depth
			{
				position16 := position
				depth++
				if buffer[position] != rune('c') {
					goto l15
				}
				position++
				if buffer[position] != rune('o') {
					goto l15
				}
				position++
				if buffer[position] != rune('n') {
					goto l15
				}
				position++
				if buffer[position] != rune('s') {
					goto l15
				}
				position++
				if buffer[position] != rune('t') {
					goto l15
				}
				position++
				if !_rules[ruleMustSpacing]() {
					goto l15
				}
				if !_rules[ruleAction7]() {
					goto l15
				}
				if !_rules[ruleTypeDef]() {
					goto l15
				}
				if !_rules[ruleMustSpacing]() {
					goto l15
				}
				if !_rules[ruleAction8]() {
					goto l15
				}
				if !_rules[ruleToken]() {
					goto l15
				}
				if !_rules[ruleSpacing]() {
					goto l15
				}
				if !_rules[ruleAction9]() {
					goto l15
				}
				if buffer[position] != rune('=') {
					goto l15
				}
				position++
				if !_rules[ruleSpacing]() {
					goto l15
				}
				if !_rules[ruleNumber]() {
					goto l15
				}
				if !_rules[ruleAction10]() {
					goto l15
				}
				if buffer[position] != rune(';') {
					goto l15
				}
				position++
				depth--
				add(ruleConstantDef, position16)
			}
			return true
		l15:
			position, tokenIndex, depth = position15, tokenIndex15, depth15
			return false
		},
		/* 4 AttributeDef <- <(('r' 'e' 'a' 'd' 'o' 'n' 'l' 'y' MustSpacing)? ('a' 't' 't' 'r' 'i' 'b' 'u' 't' 'e') MustSpacing Action11 TypeDef MustSpacing Action12 Token Spacing Action13 ';')> */
		func() bool {
			position17, tokenIndex17, depth17 := position, tokenIndex, depth
			{
				position18 := position
				depth++
				{
					position19, tokenIndex19, depth19 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l19
					}
					position++
					if buffer[position] != rune('e') {
						goto l19
					}
					position++
					if buffer[position] != rune('a') {
						goto l19
					}
					position++
					if buffer[position] != rune('d') {
						goto l19
					}
					position++
					if buffer[position] != rune('o') {
						goto l19
					}
					position++
					if buffer[position] != rune('n') {
						goto l19
					}
					position++
					if buffer[position] != rune('l') {
						goto l19
					}
					position++
					if buffer[position] != rune('y') {
						goto l19
					}
					position++
					if !_rules[ruleMustSpacing]() {
						goto l19
					}
					goto l20
				l19:
					position, tokenIndex, depth = position19, tokenIndex19, depth19
				}
			l20:
				if buffer[position] != rune('a') {
					goto l17
				}
				position++
				if buffer[position] != rune('t') {
					goto l17
				}
				position++
				if buffer[position] != rune('t') {
					goto l17
				}
				position++
				if buffer[position] != rune('r') {
					goto l17
				}
				position++
				if buffer[position] != rune('i') {
					goto l17
				}
				position++
				if buffer[position] != rune('b') {
					goto l17
				}
				position++
				if buffer[position] != rune('u') {
					goto l17
				}
				position++
				if buffer[position] != rune('t') {
					goto l17
				}
				position++
				if buffer[position] != rune('e') {
					goto l17
				}
				position++
				if !_rules[ruleMustSpacing]() {
					goto l17
				}
				if !_rules[ruleAction11]() {
					goto l17
				}
				if !_rules[ruleTypeDef]() {
					goto l17
				}
				if !_rules[ruleMustSpacing]() {
					goto l17
				}
				if !_rules[ruleAction12]() {
					goto l17
				}
				if !_rules[ruleToken]() {
					goto l17
				}
				if !_rules[ruleSpacing]() {
					goto l17
				}
				if !_rules[ruleAction13]() {
					goto l17
				}
				if buffer[position] != rune(';') {
					goto l17
				}
				position++
				depth--
				add(ruleAttributeDef, position18)
			}
			return true
		l17:
			position, tokenIndex, depth = position17, tokenIndex17, depth17
			return false
		},
		/* 5 FunctionDef <- <(Action14 TypeDef MustSpacing Action15 Token Spacing Action16 FunctionArgs Spacing Action17 ('r' 'a' 'i' 's' 'e' 's' Spacing '(' Spacing Token Spacing ')' Spacing)? ';')> */
		func() bool {
			position21, tokenIndex21, depth21 := position, tokenIndex, depth
			{
				position22 := position
				depth++
				if !_rules[ruleAction14]() {
					goto l21
				}
				if !_rules[ruleTypeDef]() {
					goto l21
				}
				if !_rules[ruleMustSpacing]() {
					goto l21
				}
				if !_rules[ruleAction15]() {
					goto l21
				}
				if !_rules[ruleToken]() {
					goto l21
				}
				if !_rules[ruleSpacing]() {
					goto l21
				}
				if !_rules[ruleAction16]() {
					goto l21
				}
				if !_rules[ruleFunctionArgs]() {
					goto l21
				}
				if !_rules[ruleSpacing]() {
					goto l21
				}
				if !_rules[ruleAction17]() {
					goto l21
				}
				{
					position23, tokenIndex23, depth23 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l23
					}
					position++
					if buffer[position] != rune('a') {
						goto l23
					}
					position++
					if buffer[position] != rune('i') {
						goto l23
					}
					position++
					if buffer[position] != rune('s') {
						goto l23
					}
					position++
					if buffer[position] != rune('e') {
						goto l23
					}
					position++
					if buffer[position] != rune('s') {
						goto l23
					}
					position++
					if !_rules[ruleSpacing]() {
						goto l23
					}
					if buffer[position] != rune('(') {
						goto l23
					}
					position++
					if !_rules[ruleSpacing]() {
						goto l23
					}
					if !_rules[ruleToken]() {
						goto l23
					}
					if !_rules[ruleSpacing]() {
						goto l23
					}
					if buffer[position] != rune(')') {
						goto l23
					}
					position++
					if !_rules[ruleSpacing]() {
						goto l23
					}
					goto l24
				l23:
					position, tokenIndex, depth = position23, tokenIndex23, depth23
				}
			l24:
				if buffer[position] != rune(';') {
					goto l21
				}
				position++
				depth--
				add(ruleFunctionDef, position22)
			}
			return true
		l21:
			position, tokenIndex, depth = position21, tokenIndex21, depth21
			return false
		},
		/* 6 FunctionArgs <- <('(' Spacing (FunctionArg Spacing Action18 (',' Spacing FunctionArg Spacing Action19)*)? ')')> */
		func() bool {
			position25, tokenIndex25, depth25 := position, tokenIndex, depth
			{
				position26 := position
				depth++
				if buffer[position] != rune('(') {
					goto l25
				}
				position++
				if !_rules[ruleSpacing]() {
					goto l25
				}
				{
					position27, tokenIndex27, depth27 := position, tokenIndex, depth
					if !_rules[ruleFunctionArg]() {
						goto l27
					}
					if !_rules[ruleSpacing]() {
						goto l27
					}
					if !_rules[ruleAction18]() {
						goto l27
					}
				l29:
					{
						position30, tokenIndex30, depth30 := position, tokenIndex, depth
						if buffer[position] != rune(',') {
							goto l30
						}
						position++
						if !_rules[ruleSpacing]() {
							goto l30
						}
						if !_rules[ruleFunctionArg]() {
							goto l30
						}
						if !_rules[ruleSpacing]() {
							goto l30
						}
						if !_rules[ruleAction19]() {
							goto l30
						}
						goto l29
					l30:
						position, tokenIndex, depth = position30, tokenIndex30, depth30
					}
					goto l28
				l27:
					position, tokenIndex, depth = position27, tokenIndex27, depth27
				}
			l28:
				if buffer[position] != rune(')') {
					goto l25
				}
				position++
				depth--
				add(ruleFunctionArgs, position26)
			}
			return true
		l25:
			position, tokenIndex, depth = position25, tokenIndex25, depth25
			return false
		},
		/* 7 FunctionArg <- <(Action20 ('i' 'n' MustSpacing)? TypeDef MustSpacing Action21 Token Action22)> */
		func() bool {
			position31, tokenIndex31, depth31 := position, tokenIndex, depth
			{
				position32 := position
				depth++
				if !_rules[ruleAction20]() {
					goto l31
				}
				{
					position33, tokenIndex33, depth33 := position, tokenIndex, depth
					if buffer[position] != rune('i') {
						goto l33
					}
					position++
					if buffer[position] != rune('n') {
						goto l33
					}
					position++
					if !_rules[ruleMustSpacing]() {
						goto l33
					}
					goto l34
				l33:
					position, tokenIndex, depth = position33, tokenIndex33, depth33
				}
			l34:
				if !_rules[ruleTypeDef]() {
					goto l31
				}
				if !_rules[ruleMustSpacing]() {
					goto l31
				}
				if !_rules[ruleAction21]() {
					goto l31
				}
				if !_rules[ruleToken]() {
					goto l31
				}
				if !_rules[ruleAction22]() {
					goto l31
				}
				depth--
				add(ruleFunctionArg, position32)
			}
			return true
		l31:
			position, tokenIndex, depth = position31, tokenIndex31, depth31
			return false
		},
		/* 8 TypeDef <- <(Action23 ('u' 'n' 's' 'i' 'g' 'n' 'e' 'd' MustSpacing Action24)? Token Action25)> */
		func() bool {
			position35, tokenIndex35, depth35 := position, tokenIndex, depth
			{
				position36 := position
				depth++
				if !_rules[ruleAction23]() {
					goto l35
				}
				{
					position37, tokenIndex37, depth37 := position, tokenIndex, depth
					if buffer[position] != rune('u') {
						goto l37
					}
					position++
					if buffer[position] != rune('n') {
						goto l37
					}
					position++
					if buffer[position] != rune('s') {
						goto l37
					}
					position++
					if buffer[position] != rune('i') {
						goto l37
					}
					position++
					if buffer[position] != rune('g') {
						goto l37
					}
					position++
					if buffer[position] != rune('n') {
						goto l37
					}
					position++
					if buffer[position] != rune('e') {
						goto l37
					}
					position++
					if buffer[position] != rune('d') {
						goto l37
					}
					position++
					if !_rules[ruleMustSpacing]() {
						goto l37
					}
					if !_rules[ruleAction24]() {
						goto l37
					}
					goto l38
				l37:
					position, tokenIndex, depth = position37, tokenIndex37, depth37
				}
			l38:
				if !_rules[ruleToken]() {
					goto l35
				}
				if !_rules[ruleAction25]() {
					goto l35
				}
				depth--
				add(ruleTypeDef, position36)
			}
			return true
		l35:
			position, tokenIndex, depth = position35, tokenIndex35, depth35
			return false
		},
		/* 9 Token <- <(<([a-z] / [A-Z] / [0-9] / '_')+> Action26)> */
		func() bool {
			position39, tokenIndex39, depth39 := position, tokenIndex, depth
			{
				position40 := position
				depth++
				{
					position41 := position
					depth++
					{
						position44, tokenIndex44, depth44 := position, tokenIndex, depth
						if c := buffer[position]; c < rune('a') || c > rune('z') {
							goto l45
						}
						position++
						goto l44
					l45:
						position, tokenIndex, depth = position44, tokenIndex44, depth44
						if c := buffer[position]; c < rune('A') || c > rune('Z') {
							goto l46
						}
						position++
						goto l44
					l46:
						position, tokenIndex, depth = position44, tokenIndex44, depth44
						if c := buffer[position]; c < rune('0') || c > rune('9') {
							goto l47
						}
						position++
						goto l44
					l47:
						position, tokenIndex, depth = position44, tokenIndex44, depth44
						if buffer[position] != rune('_') {
							goto l39
						}
						position++
					}
				l44:
				l42:
					{
						position43, tokenIndex43, depth43 := position, tokenIndex, depth
						{
							position48, tokenIndex48, depth48 := position, tokenIndex, depth
							if c := buffer[position]; c < rune('a') || c > rune('z') {
								goto l49
							}
							position++
							goto l48
						l49:
							position, tokenIndex, depth = position48, tokenIndex48, depth48
							if c := buffer[position]; c < rune('A') || c > rune('Z') {
								goto l50
							}
							position++
							goto l48
						l50:
							position, tokenIndex, depth = position48, tokenIndex48, depth48
							if c := buffer[position]; c < rune('0') || c > rune('9') {
								goto l51
							}
							position++
							goto l48
						l51:
							position, tokenIndex, depth = position48, tokenIndex48, depth48
							if buffer[position] != rune('_') {
								goto l43
							}
							position++
						}
					l48:
						goto l42
					l43:
						position, tokenIndex, depth = position43, tokenIndex43, depth43
					}
					depth--
					add(rulePegText, position41)
				}
				if !_rules[ruleAction26]() {
					goto l39
				}
				depth--
				add(ruleToken, position40)
			}
			return true
		l39:
			position, tokenIndex, depth = position39, tokenIndex39, depth39
			return false
		},
		/* 10 Number <- <(('0' 'x' <[0-9]+> Action27) / (<[0-9]+> Action28))> */
		func() bool {
			position52, tokenIndex52, depth52 := position, tokenIndex, depth
			{
				position53 := position
				depth++
				{
					position54, tokenIndex54, depth54 := position, tokenIndex, depth
					if buffer[position] != rune('0') {
						goto l55
					}
					position++
					if buffer[position] != rune('x') {
						goto l55
					}
					position++
					{
						position56 := position
						depth++
						if c := buffer[position]; c < rune('0') || c > rune('9') {
							goto l55
						}
						position++
					l57:
						{
							position58, tokenIndex58, depth58 := position, tokenIndex, depth
							if c := buffer[position]; c < rune('0') || c > rune('9') {
								goto l58
							}
							position++
							goto l57
						l58:
							position, tokenIndex, depth = position58, tokenIndex58, depth58
						}
						depth--
						add(rulePegText, position56)
					}
					if !_rules[ruleAction27]() {
						goto l55
					}
					goto l54
				l55:
					position, tokenIndex, depth = position54, tokenIndex54, depth54
					{
						position59 := position
						depth++
						if c := buffer[position]; c < rune('0') || c > rune('9') {
							goto l52
						}
						position++
					l60:
						{
							position61, tokenIndex61, depth61 := position, tokenIndex, depth
							if c := buffer[position]; c < rune('0') || c > rune('9') {
								goto l61
							}
							position++
							goto l60
						l61:
							position, tokenIndex, depth = position61, tokenIndex61, depth61
						}
						depth--
						add(rulePegText, position59)
					}
					if !_rules[ruleAction28]() {
						goto l52
					}
				}
			l54:
				depth--
				add(ruleNumber, position53)
			}
			return true
		l52:
			position, tokenIndex, depth = position52, tokenIndex52, depth52
			return false
		},
		/* 11 EndOfLine <- <(('\r' '\n') / '\n')> */
		func() bool {
			position62, tokenIndex62, depth62 := position, tokenIndex, depth
			{
				position63 := position
				depth++
				{
					position64, tokenIndex64, depth64 := position, tokenIndex, depth
					if buffer[position] != rune('\r') {
						goto l65
					}
					position++
					if buffer[position] != rune('\n') {
						goto l65
					}
					position++
					goto l64
				l65:
					position, tokenIndex, depth = position64, tokenIndex64, depth64
					if buffer[position] != rune('\n') {
						goto l62
					}
					position++
				}
			l64:
				depth--
				add(ruleEndOfLine, position63)
			}
			return true
		l62:
			position, tokenIndex, depth = position62, tokenIndex62, depth62
			return false
		},
		/* 12 Space <- <(' ' / '\t' / EndOfLine)> */
		func() bool {
			position66, tokenIndex66, depth66 := position, tokenIndex, depth
			{
				position67 := position
				depth++
				{
					position68, tokenIndex68, depth68 := position, tokenIndex, depth
					if buffer[position] != rune(' ') {
						goto l69
					}
					position++
					goto l68
				l69:
					position, tokenIndex, depth = position68, tokenIndex68, depth68
					if buffer[position] != rune('\t') {
						goto l70
					}
					position++
					goto l68
				l70:
					position, tokenIndex, depth = position68, tokenIndex68, depth68
					if !_rules[ruleEndOfLine]() {
						goto l66
					}
				}
			l68:
				depth--
				add(ruleSpace, position67)
			}
			return true
		l66:
			position, tokenIndex, depth = position66, tokenIndex66, depth66
			return false
		},
		/* 13 Comment <- <('/' '/' (!EndOfLine .)* EndOfLine)> */
		func() bool {
			position71, tokenIndex71, depth71 := position, tokenIndex, depth
			{
				position72 := position
				depth++
				if buffer[position] != rune('/') {
					goto l71
				}
				position++
				if buffer[position] != rune('/') {
					goto l71
				}
				position++
			l73:
				{
					position74, tokenIndex74, depth74 := position, tokenIndex, depth
					{
						position75, tokenIndex75, depth75 := position, tokenIndex, depth
						if !_rules[ruleEndOfLine]() {
							goto l75
						}
						goto l74
					l75:
						position, tokenIndex, depth = position75, tokenIndex75, depth75
					}
					if !matchDot() {
						goto l74
					}
					goto l73
				l74:
					position, tokenIndex, depth = position74, tokenIndex74, depth74
				}
				if !_rules[ruleEndOfLine]() {
					goto l71
				}
				depth--
				add(ruleComment, position72)
			}
			return true
		l71:
			position, tokenIndex, depth = position71, tokenIndex71, depth71
			return false
		},
		/* 14 SpaceComment <- <(Space / Comment)> */
		func() bool {
			position76, tokenIndex76, depth76 := position, tokenIndex, depth
			{
				position77 := position
				depth++
				{
					position78, tokenIndex78, depth78 := position, tokenIndex, depth
					if !_rules[ruleSpace]() {
						goto l79
					}
					goto l78
				l79:
					position, tokenIndex, depth = position78, tokenIndex78, depth78
					if !_rules[ruleComment]() {
						goto l76
					}
				}
			l78:
				depth--
				add(ruleSpaceComment, position77)
			}
			return true
		l76:
			position, tokenIndex, depth = position76, tokenIndex76, depth76
			return false
		},
		/* 15 Spacing <- <SpaceComment*> */
		func() bool {
			{
				position81 := position
				depth++
			l82:
				{
					position83, tokenIndex83, depth83 := position, tokenIndex, depth
					if !_rules[ruleSpaceComment]() {
						goto l83
					}
					goto l82
				l83:
					position, tokenIndex, depth = position83, tokenIndex83, depth83
				}
				depth--
				add(ruleSpacing, position81)
			}
			return true
		},
		/* 16 MustSpacing <- <SpaceComment+> */
		func() bool {
			position84, tokenIndex84, depth84 := position, tokenIndex, depth
			{
				position85 := position
				depth++
				if !_rules[ruleSpaceComment]() {
					goto l84
				}
			l86:
				{
					position87, tokenIndex87, depth87 := position, tokenIndex, depth
					if !_rules[ruleSpaceComment]() {
						goto l87
					}
					goto l86
				l87:
					position, tokenIndex, depth = position87, tokenIndex87, depth87
				}
				depth--
				add(ruleMustSpacing, position85)
			}
			return true
		l84:
			position, tokenIndex, depth = position84, tokenIndex84, depth84
			return false
		},
		/* 18 Action0 <- <{ p.currentIntDef = new(InterfaceDef) }> */
		func() bool {
			{
				add(ruleAction0, position)
			}
			return true
		},
		/* 19 Action1 <- <{ p.currentIntDef.Name = p.currentWord }> */
		func() bool {
			{
				add(ruleAction1, position)
			}
			return true
		},
		/* 20 Action2 <- <{ p.currentIntDef.Inherits = p.currentWord }> */
		func() bool {
			{
				add(ruleAction2, position)
			}
			return true
		},
		/* 21 Action3 <- <{ p.idl.Interfaces = append(p.idl.Interfaces, *p.currentIntDef); p.currentIntDef = nil }> */
		func() bool {
			{
				add(ruleAction3, position)
			}
			return true
		},
		/* 22 Action4 <- <{ p.currentIntDef.Constants = append(p.currentIntDef.Constants, *p.currentConDef); p.currentConDef = nil }> */
		func() bool {
			{
				add(ruleAction4, position)
			}
			return true
		},
		/* 23 Action5 <- <{ p.currentIntDef.Attributes = append(p.currentIntDef.Attributes, *p.currentAttDef); p.currentAttDef = nil }> */
		func() bool {
			{
				add(ruleAction5, position)
			}
			return true
		},
		/* 24 Action6 <- <{ p.currentIntDef.Functions = append(p.currentIntDef.Functions, *p.currentFunDef); p.currentFunDef = nil }> */
		func() bool {
			{
				add(ruleAction6, position)
			}
			return true
		},
		/* 25 Action7 <- <{ p.currentConDef = new(ConstantDef) }> */
		func() bool {
			{
				add(ruleAction7, position)
			}
			return true
		},
		/* 26 Action8 <- <{ p.currentConDef.Type = *p.currentTypDef; p.currentTypDef = nil }> */
		func() bool {
			{
				add(ruleAction8, position)
			}
			return true
		},
		/* 27 Action9 <- <{ p.currentConDef.Name = p.currentWord }> */
		func() bool {
			{
				add(ruleAction9, position)
			}
			return true
		},
		/* 28 Action10 <- <{ p.currentConDef.Value = p.currentNumber }> */
		func() bool {
			{
				add(ruleAction10, position)
			}
			return true
		},
		/* 29 Action11 <- <{ p.currentAttDef = new(AttributeDef) }> */
		func() bool {
			{
				add(ruleAction11, position)
			}
			return true
		},
		/* 30 Action12 <- <{ p.currentAttDef.Type = *p.currentTypDef; p.currentTypDef = nil }> */
		func() bool {
			{
				add(ruleAction12, position)
			}
			return true
		},
		/* 31 Action13 <- <{ p.currentAttDef.Name = p.currentWord }> */
		func() bool {
			{
				add(ruleAction13, position)
			}
			return true
		},
		/* 32 Action14 <- <{ p.currentFunDef = new(FunctionDef); }> */
		func() bool {
			{
				add(ruleAction14, position)
			}
			return true
		},
		/* 33 Action15 <- <{ p.currentFunDef.Type = *p.currentTypDef; p.currentTypDef = nil }> */
		func() bool {
			{
				add(ruleAction15, position)
			}
			return true
		},
		/* 34 Action16 <- <{ p.currentFunDef.Name = p.currentWord }> */
		func() bool {
			{
				add(ruleAction16, position)
			}
			return true
		},
		/* 35 Action17 <- <{ p.currentFunDef.Arguments = p.currentArgList; p.currentArgList = nil }> */
		func() bool {
			{
				add(ruleAction17, position)
			}
			return true
		},
		/* 36 Action18 <- <{ p.currentArgList = append(p.currentArgList, *p.currentArgDef); p.currentArgDef = nil }> */
		func() bool {
			{
				add(ruleAction18, position)
			}
			return true
		},
		/* 37 Action19 <- <{ p.currentArgList = append(p.currentArgList, *p.currentArgDef); p.currentArgDef = nil }> */
		func() bool {
			{
				add(ruleAction19, position)
			}
			return true
		},
		/* 38 Action20 <- <{ p.currentArgDef = new(ArgumentDef) }> */
		func() bool {
			{
				add(ruleAction20, position)
			}
			return true
		},
		/* 39 Action21 <- <{ p.currentArgDef.Type = *p.currentTypDef; p.currentTypDef = nil }> */
		func() bool {
			{
				add(ruleAction21, position)
			}
			return true
		},
		/* 40 Action22 <- <{ p.currentArgDef.Name = p.currentWord }> */
		func() bool {
			{
				add(ruleAction22, position)
			}
			return true
		},
		/* 41 Action23 <- <{ p.currentTypDef = new(TypeDef) }> */
		func() bool {
			{
				add(ruleAction23, position)
			}
			return true
		},
		/* 42 Action24 <- <{ p.currentTypDef.Unsigned = true }> */
		func() bool {
			{
				add(ruleAction24, position)
			}
			return true
		},
		/* 43 Action25 <- <{ p.currentTypDef.Type = p.currentWord }> */
		func() bool {
			{
				add(ruleAction25, position)
			}
			return true
		},
		nil,
		/* 45 Action26 <- <{ p.currentWord = string(buffer[begin:end]) }> */
		func() bool {
			{
				add(ruleAction26, position)
			}
			return true
		},
		/* 46 Action27 <- <{ p.currentNumber = "0x" + string(buffer[begin:end]) }> */
		func() bool {
			{
				add(ruleAction27, position)
			}
			return true
		},
		/* 47 Action28 <- <{ p.currentNumber = string(buffer[begin:end]) }> */
		func() bool {
			{
				add(ruleAction28, position)
			}
			return true
		},
	}
	p.rules = _rules
}
