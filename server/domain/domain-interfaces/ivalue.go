package domain

type Value interface {
	Equal(target Value) bool
}