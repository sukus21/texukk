package texukk

type pqEntry[T any] struct {
	item     T
	priority int
}

type pQueue[T any] struct {
	entries []pqEntry[T]
}

func (p *pQueue[T]) Add(priority int, item T) {
	p.entries = append(p.entries, pqEntry[T]{
		item:     item,
		priority: priority,
	})
}

func (p *pQueue[T]) Sort() []T {
	l := len(p.entries)
	sort := make([]pqEntry[T], l)
	copy(sort, p.entries)
	for j := 0; j < l-1; j++ {
		for i := 1; i < l-j; i++ {
			if sort[i].priority > sort[i-1].priority {
				tmp := sort[i]
				sort[i] = sort[i-1]
				sort[i-1] = tmp
			}
		}
	}

	out := make([]T, l)
	for i, v := range sort {
		out[i] = v.item
	}
	return out
}
