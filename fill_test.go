package easy

import (
	"strings"
	"testing"
)

type M struct {
	Id   int64
	Name string
}

type V struct {
	Id        int64
	Name      string
	FirstName string
}

func MtoV(m any) any {
	ms := m.(*M)
	return &V{
		Id:        ms.Id,
		Name:      ms.Name,
		FirstName: strings.Split(ms.Name, " ")[0],
	}
}

func TestFill(t *testing.T) {
	ms := []*M{
		{
			Id:   1,
			Name: "Alice White",
		},
		{
			Id:   2,
			Name: "Ace Black",
		},
		{
			Id:   3,
			Name: "Bob Black",
		},
		{
			Id:   4,
			Name: "Helen Black",
		},
		{
			Id:   5,
			Name: "Monkey Black",
		},
		{
			Id:   6,
			Name: "Stand Black",
		},
		{
			Id:   7,
			Name: "Jack Black",
		},
		{
			Id:   8,
			Name: "Panda Black",
		},
	}

	vs := make([]*V, len(ms))
	Fill(ms, vs, MtoV, 5)

	ms = nil
	vs = make([]*V, len(ms))
	Fill(ms, vs, MtoV, 5)
}
