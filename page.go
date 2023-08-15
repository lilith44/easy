package easy

// Page defines a pagination object.
type Page struct {
	Page       int64 `json:"page"`
	Size       int64 `json:"size"`
	HasNext    bool  `json:"hasNext"`
	HasPrev    bool  `json:"hasPrev"`
	TotalPage  int64 `json:"totalPage"`
	TotalItems int64 `json:"totalItems"`
	Items      any   `json:"items"`
}

// NewPage creates a new pagination object.
func NewPage(items any, total int64, page int64, size int64) *Page {
	totalPage := total / size
	if (total % size) != 0 {
		totalPage++
	}

	return &Page{
		Page:       page,
		Size:       size,
		HasNext:    page < totalPage,
		HasPrev:    page > 1,
		TotalPage:  totalPage,
		TotalItems: total,
		Items:      items,
	}
}
