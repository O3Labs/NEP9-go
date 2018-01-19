package nep9

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestParseNEP9(t *testing.T) {
	uri := "neo:AeNkbJdiMx49kBStQdDih7BzfDwyTNVRfb?assetID=602c79718b16e442de58778e148d0b1084e3b2dffd5de6b7b16cee7969282de7&amount=0.11&description=This%20is%20awesome&ecdh02=02ed53ad58c838435d4dd7a4b25c1eba01384c814ca53a539405434807afbb04b4"
	nep9, err := ParseNEP9URI(uri)
	if err != nil {
		t.Fail()
		return
	}
	b, _ := json.Marshal(nep9)
	fmt.Printf("%v", string(b))
}
