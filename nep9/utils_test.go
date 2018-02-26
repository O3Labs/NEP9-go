package nep9_test

import (
	"testing"

	"github.com/o3labs/NEP9-go/nep9"
)

func TestValidateNEOAddress(t *testing.T) {
	valid := nep9.ValidateNEOAddress("AeNkbJdiMx49kBStQdDih7BzfDwyTNVRfb")
	if valid == false {
		t.Fail()
	}
}

func TestValidateInvalidNEOAddress(t *testing.T) {
	valid := nep9.ValidateNEOAddress("AeNkbJdiMx49kBStQdDih7BzfDwyTNVRfb1")
	if valid == true {
		t.Fail()
	}
}

func TestValidateScriptHash(t *testing.T) {
	valid := nep9.ValidateSmartContractScriptHash("ce575ae1bb6153330d20c560acb434dc5755241b")
	if valid == false {
		t.Fail()
	}
}
