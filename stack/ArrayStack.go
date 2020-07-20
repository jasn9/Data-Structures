package stack

type ArrayStack struct {
	S []interface{}
}

func (a *ArrayStack) IsEmpty() bool {
	return len(a.S) == 0
}

func (a *ArrayStack) AddAtLast(ele interface{}) error {
	a.S = append(a.S, ele)
	return nil
}

func (a *ArrayStack) RemoveFromLast() interface{}  {
	ele := a.S[len(a.S)-1]
	a.S = a.S[:len(a.S)-1]
	return ele
}