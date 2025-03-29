package pq

type QueueElement struct {
	Priority int
	Value    interface{}
}

type Pq struct {
	Heap   []QueueElement
	Minmax int // whether this is a min-pq or max-pq
}

func New(Minmax string) *Pq {
	pq := new(Pq)
	if Minmax == "max" {
		pq.Minmax = -1
	} else {
		pq.Minmax = 1
	}
	return pq
}

func (p *Pq) Length() int {
	return len(p.Heap)
}

func (p *Pq) Push(priority int, value interface{}) {
	qe := QueueElement{Priority: priority, Value: value}
	p.Heap = append(p.Heap, qe)
	p.bubbleUp(len(p.Heap) - 1)
}

func (p *Pq) Pull() interface{} {
	if len(p.Heap) == 0 {
		return nil
	}
	root := p.Heap[0]
	p.Heap[0] = p.Heap[len(p.Heap)-1]
	p.Heap = p.Heap[:len(p.Heap)-1]
	p.bubbleDown(0)
	// return root
	return root.Value
}

func (p *Pq) bubbleDown(idx int) {
	for {
		childIdx1 := idx*2 + 1
		childIdx2 := idx*2 + 2
		smallest := idx

		if childIdx1 < len(p.Heap) &&
			p.Heap[childIdx1].Priority*p.Minmax <
				p.Heap[smallest].Priority*p.Minmax {
			smallest = childIdx1
		}
		if childIdx2 < len(p.Heap) &&
			p.Heap[childIdx2].Priority*p.Minmax <
				p.Heap[smallest].Priority*p.Minmax {
			smallest = childIdx2
		}
		if smallest == idx {
			break
		}
		p.Heap[idx], p.Heap[smallest] = p.Heap[smallest], p.Heap[idx]
		idx = smallest
	}
}

func (p *Pq) bubbleUp(idx int) {
	for idx > 0 {
		parentIdx := (idx - 1) / 2
		if p.Heap[parentIdx].Priority*p.Minmax > p.Heap[idx].Priority*p.Minmax {
			p.Heap[parentIdx], p.Heap[idx] = p.Heap[idx], p.Heap[parentIdx]
			idx = parentIdx
		} else {
			break
		}
	}
}
