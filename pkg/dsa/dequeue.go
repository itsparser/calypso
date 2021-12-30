package dsa

type Queue struct {
	len        int
	head, tail int
	q          []float64
}

func New(n int) *Queue {
	return &Queue{n, 0, 0, make([]float64, n)}
}

func (p *Queue) Enqueue(x float64) bool {
	p.q[p.tail] = x
	p.tail = (p.tail + 1) % p.len
	return p.head != p.tail
}

func (p *Queue) Dequeue() (float64, bool) {
	if p.head == p.tail {
		return 0, false
	}
	x := p.q[p.head]
	p.head = (p.head + 1) % p.len
	return x, true
}
