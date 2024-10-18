package model

import "github.com/alxrusinov/diploma/internal/mathfn"

type Balance struct {
	Current   float64 `json:"current"`
	Withdrawn float64 `json:"withdrawn"`
}

func (b *Balance) Round() {
	b.Current = mathfn.RoundFloat(b.Current, 5)
	b.Withdrawn = mathfn.RoundFloat(b.Withdrawn, 5)
}
