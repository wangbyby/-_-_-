package data

type Trie struct {
	inactive         bool
	Key              rune
	Value            interface{}
	Left, Right, Mid *Trie
}

func NewTrie() *Trie {
	return &Trie{}
}

func (t *Trie) InsertString(s string, val interface{}) {
	tmp := t
	for i, k := range s {
		if tmp.Key == 0 || tmp.Key == k {
			tmp.Key = k
			if i+1 == len(s) {
				break
			}
			if tmp.Mid == nil {
				tmp.Mid = &Trie{}
			}

			tmp = tmp.Mid

		} else if tmp.Key > k {
			if tmp.Left == nil {
				tmp.Left = &Trie{}
			}
			tmp = tmp.Left
		} else {
			if tmp.Right == nil {
				tmp.Right = &Trie{}
			}
			tmp = tmp.Right
		}
	}
	tmp.Value = val
}

func (t *Trie) SearchString(s string) (v interface{}) {
	v, node := t.searchString(s)
	if node.inactive == true {
		v = nil
	}
	return
}

func (t *Trie) DelNode(s string) interface{} {
	v, node := t.searchString(s)
	node.inactive = true
	return v
}

func (t *Trie) searchString(s string) (v interface{}, trieNode *Trie) {
	tmp := t
	for i, k := range s {
		if tmp.Key == k {
			if i+1 == len(s) {
				//end of string
				return tmp.Value, tmp
			}
			if tmp.Mid == nil {
				return nil, nil
			}
			tmp = tmp.Mid
		} else if tmp.Key > k {
			if tmp.Left == nil {
				return nil, nil
			}
			tmp = tmp.Left
		} else {
			if tmp.Right == nil {
				return nil, nil
			}
			tmp = tmp.Right
		}
	}
	return tmp.Value, tmp
}
