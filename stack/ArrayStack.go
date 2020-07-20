package stack

type ArrayStack struct {
	S []interface{}
}

func (a *ArrayStack) IsEmpty() bool {
	return len(a.S) == 0
}

func (a *ArrayStack) AddElement(ele interface{}) error {
	a.S = append(a.S, ele)
	return nil
}

func (a *ArrayStack) RemoveLastElementAdded() interface{}  {
	ele := a.S[len(a.S)-1]
	a.S = a.S[:len(a.S)-1]
	return ele
}