package main

import "fmt"

const sizeOfCache = 5

func main() {
	fmt.Println("Start cache")
	cache := NewCache()
	initialSample := []string{"Real Madrid", "Bayern Munich", "Manchester City", "Barcelona", "PSG"}
	testSample := []string{"Liverpool", "Manchester City", "Bayern Munich", "Real Madrid", "Chelsea", "Barcelona", "PSG"}

	fmt.Println("Populating cache")
	for _, word := range initialSample {
		cache.Check(word)
		cache.Display()
	}
	fmt.Println()
	fmt.Println("Testing cache")
	for _, word := range testSample {
		cache.Check(word)
		cache.Display()
	}

	fmt.Println()
	cache.Flush()
	cache.Display()

}

type Node struct {
	Left  *Node
	Value string
	Right *Node
}

type Queue struct {
	Head   *Node
	Length int
	Tail   *Node
}

type Cache struct {
	Queue Queue
	Hash  Hash
}

type Hash map[string]*Node

func NewCache() Cache {
	return Cache{Queue: NewQueue(), Hash: Hash{}}
}

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}
	head.Right = tail
	tail.Left = head

	return Queue{Head: head, Tail: tail}
}

func (c *Cache) Check(str string) {
	node := &Node{}

	if val, ok := c.Hash[str]; ok {
		node = c.Remove(val)
	} else {
		node = &Node{Value: str}
	}
	c.Add(node)
	c.Hash[str] = node
}

func (c *Cache) Add(node *Node) {
	fmt.Printf("add: %s\n", node.Value)
	tmp := c.Queue.Head.Right
	c.Queue.Head.Right = node
	node.Left = c.Queue.Head
	node.Right = tmp
	tmp.Left = node
	c.Queue.Length += 1

	if c.Queue.Length > sizeOfCache {
		c.Remove(c.Queue.Tail.Left)
	}
}

func (c *Cache) Remove(node *Node) *Node {
	fmt.Printf("remove: %s\n", node.Value)
	left := node.Left
	right := node.Right
	left.Right = right
	right.Left = left

	c.Queue.Length -= 1
	delete(c.Hash, node.Value)
	return node
}

func (c *Cache) Display() {
	c.Queue.Display()
}

func (q *Queue) Display() {
	node := q.Head.Right
	fmt.Printf("%d - [", q.Length)
	for i := 0; i < q.Length; i++ {
		fmt.Printf("{%s}", node.Value)
		if i < q.Length-1 {
			fmt.Printf("<-->")
		}
		node = node.Right
	}
	fmt.Println("]")
}

func (c *Cache) Flush() {
	fmt.Println("Flushing cache")
	*c = NewCache()
}
