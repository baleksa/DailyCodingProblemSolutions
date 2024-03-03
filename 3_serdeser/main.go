package main

import (
	"strings"
)

type Node struct {
	val   string
	left  *Node
	right *Node
}

func serialize(root *Node) string {
	if root == nil {
		return "nil"
	}
	return "Node('" + root.val + "', " + serialize(root.left) + ", " + serialize(root.right) + ")"
}

func deserialize(s string) *Node {
	if len(s) == 0 || s == "nil" {
		return nil
	}
	// fmt.Println("String to deser:", s)
	fh := strings.Index(s, "'")
	sh := strings.Index(s[fh+1:], "'") + fh + 1
	val := s[fh+1 : sh]

	if strings.Count(s, "Node('") == 1 {
		return &Node{val, nil, nil}
	}

	lstart := sh + 3
	if s[lstart:lstart+3] == "nil" {
		rstart := lstart + 5
		rend := len(s) - 2
		return &Node{val, nil, deserialize(s[rstart : rend+1])}
	}

	rend := len(s) - 2
	if s[rend-2:rend+1] == "nil" {
		lend := rend - 5
		return &Node{val, deserialize(s[lstart : lend+1]), nil}
	}

	lend := matchingClosedParen(s, lstart+4)
	l := s[lstart : lend+1]
	rstart := lend + 3
	r := s[rstart : rend+1]

	// fmt.Println("Left child str:", l)
	// fmt.Println("Right child str:", r)

	node := &Node{val, deserialize(l), deserialize(r)}

	return node
}

func matchingClosedParen(s string, io int) int {
	no := 1
	ic := io + 1
	for ic < len(s) && no > 0 {
		switch s[ic] {
		case '(':
			no++
		case ')':
			no--
			if no == 0 {
				return ic
			}
		}
		ic++
	}
	return -1
}

func main() {
	root := &Node{"root", &Node{"left", &Node{"left.left", nil, nil}, nil}, &Node{"right", nil, nil}}

	if deserialize(serialize(root)).left.left.val != "left.left" {
		panic("error")
	}
}
