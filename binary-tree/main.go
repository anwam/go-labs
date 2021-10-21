package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/google/uuid"
)

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

type Node struct {
	ID            string
	Name          string
	Value         int
	Left          *Node
	Right         *Node
	Parent        *Node
	Depth         int
	TotalChildren int
}

func NewNode(value int) *Node {
	n := new(Node)
	n.ID = uuid.NewString()
	n.Value = value
	n.Depth = 1
	return n
}

func (n *Node) GetRootNode() *Node {
	if n.Parent != nil {
		return n.Parent.GetRootNode()
	}
	return n
}

func (n *Node) AddLeft(left *Node) {
	if n.Left == nil {
		n.Left = left
		n.TotalChildren = n.TotalChildren + 1
		r := n.GetRootNode()
		r.TotalChildren = r.TotalChildren + 1
		left.Parent = n
		left.Depth = n.Depth + 1
	} else {
		n.Left.AddChild(left)
	}
}

func (n *Node) AddRight(right *Node) {
	if n.Right == nil {
		n.Right = right
		n.TotalChildren = n.TotalChildren + 1
		r := n.GetRootNode()
		r.TotalChildren = r.TotalChildren + 1
		right.Parent = n
		right.Depth = n.Depth + 1
	} else {
		n.Right.AddChild(right)
	}
}

func (n *Node) AddChild(c *Node) {
	if n.Value > c.Value {
		n.AddLeft(c)
	} else if n.Value < c.Value {
		n.AddRight(c)
	} else {
		return
	}
}

func (n *Node) NewChild(randLen int) {
	randInt := seededRand.Intn(randLen)
	c := NewNode(randInt)
	n.AddChild(c)
}

func (n *Node) DFValue() {
	if n != nil {
		fmt.Printf("%s:%d ", n.Name, n.Value)
		if n.Left != nil {
			n.Left.DFValue()
			if n.Right != nil {
				n.Right.DFValue()
			}
		} else {
			fmt.Println("")
		}
	}
}

func (n Node) HasChild() bool {
	hasChild := false
	if n.Left != nil {
		hasChild = true
	}
	if n.Right != nil {
		hasChild = true
	}
	return hasChild
}

func (n Node) HasLeftChild() bool {
	return n.Left != nil
}

func (n Node) HasRightChild() bool {
	return n.Right != nil
}

func (n Node) HasSiblings() bool {
	return n.Parent.HasRightChild()
}

func (n Node) PrintChildren(connector string, isLeft bool, t int) {
	var prefix string
	var childPrefix string
	if isLeft {
		if n.HasSiblings() {
			childPrefix = connector + "│   "
			prefix = connector + "├── "
		} else {
			prefix = connector + "└── "
			childPrefix = connector + "    "
		}
	} else {
		childPrefix = connector + "    "
		prefix = connector + "└── "
	}
	var value string
	if n.Value == t {
		value = strconv.Itoa(n.Value) + " 🟢"
	} else {
		value = strconv.Itoa(n.Value)
	}
	fmt.Printf("%s%s\n", prefix, value)
	if n.HasChild() {
		if n.HasLeftChild() {
			n.Left.PrintChildren(childPrefix, true, t)
		}
		if n.HasRightChild() {
			n.Right.PrintChildren(childPrefix, false, t)
		}
	}
}

func (n *Node) Find(target int) *Node {
	if n.Value == target {
		return n
	} else {
		if target < n.Value {
			if n.HasLeftChild() {
				return n.Left.Find(target)
			}
		} else {
			if n.HasRightChild() {
				return n.Right.Find(target)
			}
		}
		return nil
	}
}

func main() {
	println("Hello this is binary tree.")
	start := time.Now()

	randLen := 50
	randInt := seededRand.Intn(randLen)
	a := NewNode(randInt)
	n := 0
	for n < randLen {
		a.NewChild(randLen)
		n++
	}

	target := seededRand.Intn(randLen)
	fmt.Printf("a has %d children.\n", a.TotalChildren)
	a.PrintChildren("", false, target)

	targetNode := a.Find(target)
	if targetNode != nil {
		fmt.Printf("found target node %d after node %d \n", target, targetNode.Parent.Value)
	} else {
		fmt.Printf("%d not found.\n", target)
	}
	fmt.Println("execution time took ", time.Since(start))

}
