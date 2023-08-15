package easy

// Paging defines a pagination request object.
type Paging struct {
	Page    int64  `default:"1" query:"page" json:"page" validate:"min=1"`
	Size    int64  `default:"20" query:"size" json:"size" validate:"min=1"`
	Keyword string `query:"keyword" json:"keyword" validate:"max=255"`
}

// Limit gets the limit parameter for database query.
func (p Paging) Limit() int64 {
	return p.Size
}

// Offset gets the offset parameter for database query.
func (p Paging) Offset() int64 {
	return (p.Page - 1) * p.Size
}
