package main

import "fmt"

type Trie struct {
	Root *Node
}

type Node struct {
	IsEnd    bool
	Children map[string]*Node
}

func NewTrie() *Trie {
	return &Trie{Root: NewNode()}
}

func NewNode() *Node {
	return &Node{
		IsEnd:    false, //是否是最后一个元素
		Children: make(map[string]*Node),
	}
}

//插入
func (this *Trie) Insert(str string) {
	current := this.Root
	for _, v := range []rune(str) {
		if _, ok := current.Children[string(v)]; !ok {
			current.Children[string(v)] = NewNode()
		}
		current = current.Children[string(v)]
	}
	current.IsEnd = true
}

//搜索
func (this *Trie) Search(str string) bool {
	current := this.Root
	for _, v := range []rune(str) {
		if _, ok := current.Children[string(v)]; !ok {
			return false
		}
		current = current.Children[string(v)]
	}
	return current.IsEnd
}

func main() {
	strs := []string{"go", "gin", "golang", "goapp", "guest"}
	tree := NewTrie()
	for _, s := range strs {
		tree.Insert(s)
	}
	strs = append(strs, "abc", "gogo")
	for _, s := range strs {
		fmt.Println(tree.Search(s))
	}
}
