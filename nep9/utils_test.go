package nep9_test

import (
	"log"
	"testing"

	"github.com/o3labs/NEP9-go/nep9"
)

func TestValidateNEOAddress(t *testing.T) {
	valid := nep9.ValidateNEOAddress("AeNkbJdiMx49kBStQdDih7BzfDwyTNVRfb")
	log.Printf("%v", valid)
}

func TestValidateScriptHash(t *testing.T) {
	valid := nep9.ValidateSmartContractScriptHash("ce575ae1bb6153330d20c560acb434dc5755241b")
	log.Printf("%v", valid)
}
