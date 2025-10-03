package ast

type SymobolType struct {
	Name string //T
}

func (t *SymobolType) _type() {

}

type ArrayType struct {
	Underlying Type //[]T
}

func (t *ArrayType) _type() {
}
