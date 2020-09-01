package main

type Trie struct {
	root *node
}

type node struct {
	childes   [MaxCap]*node
	isWordEnd bool
}

const MaxCap = 26

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		root: &node{},
	}
}

/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	current := t.root

	for _, r := range word {
		index := r - 'a'
		if current.childes[index] == nil {
			current.childes[index] = &node{}
		}

		current = current.childes[index]
	}

	current.isWordEnd = true
}

/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	current := t.root

	for _, r := range word {
		index := r - 'a'
		if current.childes[index] == nil {
			return false
		}

		current = current.childes[index]
	}

	return current.isWordEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	current := t.root

	for _, r := range prefix {
		index := r - 'a'
		if current.childes[index] == nil {
			return false
		}

		current = current.childes[index]
	}

	return true
}

func main() {
	obj := Constructor()
	obj.Insert("apple")
	println(obj.Search("apple"))
	println(obj.Search("app"))
	println(obj.StartsWith("app"))
	obj.Insert("apple")
	obj.Search("app")
}
