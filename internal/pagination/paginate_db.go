package pagination

import "math"

type Pagination struct {
	items   Paginator
	curPage int
	limit   int
}

func NewPagination(items Paginator, curPage int, limit int) *Pagination {
	if curPage < 1 {
		curPage = 1
	}

	if limit <= 0 {
		limit = 100
	}

	return &Pagination{
		items:   items,
		curPage: curPage,
		limit:   limit,
	}
}

func (p *Pagination) CurrentPage() (map[string]interface{}, error) {
	total, err := p.items.Count()
	if err != nil {
		return nil, err
	}

	limit := p.limit
	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	nextPage := p.curPage + 1
	if nextPage > totalPages {
		nextPage = totalPages
	}

	prevPage := p.curPage - 1
	if prevPage < 1 {
		prevPage = 1
	}

	start := (p.curPage - 1) * limit
	data, err := p.items.GetAll(limit, start)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"pageNum":    p.curPage,
		"totalPages": totalPages,
		"nextPage":   nextPage,
		"prevPage":   prevPage,
		"items":      data,
		"total":      total,
	}, nil
}
