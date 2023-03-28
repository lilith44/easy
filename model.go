package easy

import "time"

type Model interface {
	PK() int64
	SetPK(pk int64)
	UseSnowflakeIdAsDefault() bool
}

type BaseModel struct {
	Id        int64     `xorm:"pk notnull comment('主键')" bson:"id"`
	CreatedAt time.Time `xorm:"datetime created notnull default('2023-01-01 00:00:00') comment('创建时间')" bson:"created_at"`
}

func (bm *BaseModel) PK() int64 {
	return bm.Id
}

func (bm *BaseModel) SetPK(pk int64) {
	bm.Id = pk
}

func (bm *BaseModel) UseSnowflakeIdAsDefault() bool {
	return true
}

type AutoIncrModel struct {
	Id        int64     `xorm:"pk autoincr notnull comment('主键')" bson:"id"`
	CreatedAt time.Time `xorm:"datetime created notnull default('2023-01-01 00:00:00') comment('创建时间')" bson:"created_at"`
}

func (am *AutoIncrModel) PK() int64 {
	return am.Id
}

func (am *AutoIncrModel) SetPK(pk int64) {
	am.Id = pk
}

func (am *AutoIncrModel) UseSnowflakeIdAsDefault() bool {
	return false
}
