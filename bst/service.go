package bst

type Service interface {
	Insert(v int)
	Search(v int) bool
	Remove(v int)
}

type service struct {
	bst *IntBinarySearchTree
}

func (s *service) Insert(v int) {
	s.bst.Insert(v)
}

func (s *service) Search(v int) bool {
	return s.bst.Search(v)
}

func (s *service) Remove(v int) {
	s.bst.Remove(v)
}

func NewService(input []int) Service {
	return &service{
		bst: NewIntBinarySearchTree(input),
	}
}
