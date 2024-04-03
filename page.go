package easy

// Page defines a pagination object.
type Page struct {
	Page       int64 `json:"page"`
	Size       int64 `json:"size"`
	HasNext    bool  `json:"has_next"`
	HasPrev    bool  `json:"has_prev"`
	TotalPage  int64 `json:"total_page"`
	TotalItems int64 `json:"total_items"`
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
