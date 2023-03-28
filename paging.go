package easy

// Paging defines basic params of pagination querying.
type Paging struct {
	Page    int64  `default:"1" query:"page" validate:"min=1"`
	Size    int64  `default:"20" query:"size" validate:"min=1"`
	Keyword string `query:"keyword" validate:"max=255"`
}

func (p Paging) Limit() int64 {
	return p.Size
}

func (p Paging) Offset() int64 {
	return (p.Page - 1) * p.Size
}
