package main

/*
  解题思路：注意数据类型，for 遍历出来的字符和索引访问的不一致。
  时间复杂度：O(n)
 */

type Trie struct {
	isWord bool
	node map[byte]*Trie
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{isWord: false, node: make(map[byte]*Trie)}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	tree := this
	for i := 0; i < len(word); i++ {
		nextTree, ok := tree.node[word[i]]
		if !ok {
			newNode := Constructor()
			tree.node[word[i]] = &newNode
			tree = &newNode
		} else {
			tree = nextTree
		}
	}
	tree.isWord = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	tree := this
	for i := 0; i < len(word); i++ {
		nextTree, ok := tree.node[word[i]]
		if !ok {
			return false
		} else {
			tree = nextTree
		}
	}
	if tree.isWord {
		return true
	}
	return false
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	tree := this
	for i := 0; i < len(prefix); i++ {
		nextTree, ok := tree.node[prefix[i]]
		if !ok {
			return false
		} else {
			tree = nextTree
		}
	}
	return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */