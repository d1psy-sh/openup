package v1

import "github.com/charmbracelet/bubbles/list"

type Node struct {
    value list.Model
    prev *Node
}

type Stack struct {
    length int
    tail *Node
}

func (s *Stack) push(list list.Model)  {
    // TODO: push to Stack
}

func (s *Stack) pop() list.Model {
    // TODO: pop from Stack
}
