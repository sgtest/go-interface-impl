package ifaceimpl

import (
	"github.com/sgtest/go-interface-def/ifacedef"
)

type JImpl struct{}

func (_ JImpl) F() {}

var jimpl ifacedef.J = JImpl{}

type KImpl struct {
	JImpl
}

func (_ KImpl) G() string { return "qux" }

var kimpl ifacedef.K = KImpl{}

type IImpl struct {
	KImpl
}

func (_ IImpl) H(x string) error { return nil }

var iimpl ifacedef.I = IImpl{}

type LImpl struct{}

func (_ LImpl) F() ifacedef.T { return "qux" }

var limpl ifacedef.L = LImpl{}

type privateJImpl struct{}

func (_ privateJImpl) F() {}

var privatejimpl ifacedef.J = privateJImpl{}
