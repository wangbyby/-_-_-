package set

import (
	"errors"
)

//使用 map的key 用作set
type SetHash struct {
	Set   map[interface{}]bool // 使用 key 存储集合
	order int                  //集合的阶数
}

func NewSetHash() SetHash {
	return SetHash{Set: make(map[interface{}]bool, 0), order: 0}
}

func (s *SetHash) Search(v interface{}) error {
	_, ok := s.Set[v]
	if !ok {
		return errors.New("err Element")
	}
	return nil
}
func (s *SetHash) GetOrder() int {
	return s.order
}

func (s *SetHash) AddElement(k interface{}) {
	s.Set[k] = true
	s.order++
}
func (s *SetHash) DelElement(k interface{}) error {
	if _, ok := s.Set[k]; !ok {
		return errors.New("Do not contain this key")
	}
	delete(s.Set, k)
	s.order--
	return nil
}

func (s *SetHash) Equel(a *SetHash) (equ bool) {
	if s.order != a.order {
		return
	}

	// key --> value
	for i, _ := range a.Set {
		if _, ok := s.Set[i]; !ok {
			return
		}
	}
	return true
}
