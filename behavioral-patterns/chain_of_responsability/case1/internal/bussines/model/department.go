package model

type Department interface {
	Execute(*Patient)
	SetNext(Department)
}
