package dsa

// CList is a list with a maximum size.
// It is a doubly linked list.
type CList struct {
	Len, max int
	Q        []interface{}
}

func NewCList(n int) *CList {
	_list := &CList{0, n, make([]interface{}, n)}
	return _list
}

func (p *CList) Enqueue(x ...interface{}) {
	p.Q = append(p.Q, x...)
	p.Len = len(p.Q)
	if _len := p.Len - p.max; _len > 0 {
		p.Q = p.Q[_len:]
	}
}

func (p *CList) Dequeue() interface{} {
	if p.Len == 0 {
		return nil
	}
	_return := p.Q[0]
	p.Q = p.Q[1:]
	p.Len = p.Len - 1
	return _return
}
