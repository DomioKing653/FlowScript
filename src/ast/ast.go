package ast

import "github.com/DomioKing653/FlowScript/src/runtime"

type Statement interface {
	stmt() runtime.RuntimeValue
}

type Expression interface {
	expr() runtime.RuntimeValue
}

type Type interface {
	_type()
}
