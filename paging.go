package easy

// Paging defines basic conditions of pagination querying.
type Paging struct {
	Page    int64  `query:"page" validate:"min=1"`
	Size    int64  `query:"size" validate:"min=1"`
	Keyword string `query:"keyword" validate:"max=255"`
}

// Limit returns limit, offset query conditions for SQL.
func (p Paging) Limit() (int, int) {
	return int(p.Size), int((p.Page - 1) * p.Size)
}
