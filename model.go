package easy

import "time"

type BaseModel struct {
	Id        int64     `xorm:"pk notnull comment('主键')"`
	CreatedAt time.Time `xorm:"datetime created notnull default('2023-01-01 00:00:00') comment('创建时间')"`
}

type Model interface {
	PK() int64
	SetPK(pk int64)
}

func (bm *BaseModel) PK() int64 {
	return bm.Id
}

func (bm *BaseModel) SetPK(pk int64) {
	bm.Id = pk
}
