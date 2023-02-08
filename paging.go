package easy

// Paging defines basic conditions of pagination querying.
type Paging struct {
	Page    int64  `default:"1" query:"page" validate:"min=1"`
	Size    int64  `default:"20" query:"size" validate:"min=1"`
	Keyword string `query:"keyword" validate:"max=255"`
	SortBy  string `query:"sortBy" validate:"omitempty"`
}

// OrderBy returns an order by query condition for SQL.
func (p Paging) OrderBy() string {
	return p.SortBy
}

// Limit returns limit, offset query conditions for SQL.
func (p Paging) Limit() (int, int) {
	return int(p.Size), int((p.Page - 1) * p.Size)
}
