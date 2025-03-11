package pq

type QueueElement struct {
	Priority int
	Value    any
}

type Pq struct {
	heap   []QueueElement
	minmax int
}

func Constructor(minmax string) Pq {
	pq := new(Pq)
	if minmax == "max" {
		pq.minmax = -1
	} else {
		pq.minmax = 1
	}
	return *pq
}

func (p *Pq) push(Value QueueElement) {
	p.heap = append(p.heap, Value)
	p.bubbleUp(len(p.heap) - 1)
}

func (p *Pq) pull() any {
	if len(p.heap) == 0 {
		return nil
	}
	root := p.heap[0]
	p.heap[0] = p.heap[len(p.heap)-1]
	p.heap = p.heap[:len(p.heap)-1]
	p.bubbleDown(0)
	return root
}

func (p *Pq) bubbleDown(idx int) {
	for {
		childIdx1 := idx*2 + 1
		childIdx2 := idx*2 + 2
		smallest := idx

		if childIdx1 < len(p.heap) &&
			p.heap[childIdx1].Priority*p.minmax <
				p.heap[smallest].Priority*p.minmax {
			smallest = childIdx1
		}
		if childIdx2 < len(p.heap) &&
			p.heap[childIdx2].Priority*p.minmax <
				p.heap[smallest].Priority*p.minmax {
			smallest = childIdx2
		}
		if smallest == idx {
			break
		}
		p.heap[idx], p.heap[smallest] = p.heap[smallest], p.heap[idx]
		idx = smallest
	}
}

func (p *Pq) bubbleUp(idx int) {
	for idx > 0 {
		parentIdx := (idx - 1) / 2
		if p.heap[parentIdx].Priority*p.minmax > p.heap[idx].Priority*p.minmax {
			p.heap[parentIdx], p.heap[idx] = p.heap[idx], p.heap[parentIdx]
			idx = parentIdx
		} else {
			break
		}
	}
}
