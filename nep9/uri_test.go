package nep9_test

import (
	"log"
	"testing"

	"github.com/o3labs/NEP9-go/nep9"
	"github.com/stretchr/testify/assert"
)

func TestURI(t *testing.T) {
	t.Run("NewURI()", func(t *testing.T) {
		t.Run("Valid", func(t *testing.T) {
			rawURI := "neo:AeNkbJdiMx49kBStQdDih7BzfDwyTNVRfb?assetID=602c79718b16e442de58778e148d0b1084e3b2dffd5de6b7b16cee7969282de7&amount=0.11&description=This%20is%20awesome&ecdh02=02ed53ad58c838435d4dd7a4b25c1eba01384c814ca53a539405434807afbb04b4"
			uri, err := nep9.NewURI(rawURI)

			assert.NoError(t, err)
			assert.IsType(t, &nep9.URI{}, uri)
		})

		t.Run("InvalidScheme", func(t *testing.T) {
			rawURI := "bitcoin:AeNkbJdiMx49kBStQdDih7BzfDwyTNVRfb?assetID=602c79718b16e442de58778e148d0b1084e3b2dffd5de6b7b16cee7969282de7&amount=0.11"
			uri, err := nep9.NewURI(rawURI)

			assert.Error(t, err)
			assert.Nil(t, uri)
		})

		t.Run("InvalidAssetID", func(t *testing.T) {
			rawURI := "neo:AeNkbJdiMx49kBStQdDih7BzfDwyTNVRfb?assetID=502c79718b16e442de58778e148d0b1084e3b2dffd5de6b7b16cee7969282de7&amount=0.11"
			uri, err := nep9.NewURI(rawURI)

			assert.Error(t, err)
			assert.Nil(t, uri)
		})

		t.Run("InvalidAmount", func(t *testing.T) {
			rawURI := "neo:AeNkbJdiMx49kBStQdDih7BzfDwyTNVRfb?assetID=602c79718b16e442de58778e148d0b1084e3b2dffd5de6b7b16cee7969282de7&amount=foo"
			uri, err := nep9.NewURI(rawURI)

			assert.Error(t, err)
			assert.Nil(t, uri)
		})
	})
}

func TestParseNEP9SmartContract(t *testing.T) {
	//if smart contract conforms the main(operation string, param args []Object)
	//you could pare the URI with this
	rawURI := "neo:ce575ae1bb6153330d20c560acb434dc5755241b?operation=transfer&params=AK2nJJpJr6o664CWJKi1QRXjqeic2zRp8y,AK2nJJpJr6o664CWJKi1QRXjqeic2zRp8y,0.5"
	//rawURI := "neo:ce575ae1bb6153330d20c560acb434dc5755241b?operation=balanceOf"
	uri, err := nep9.NewURI(rawURI)
	if err != nil {
		log.Printf("%v", err)
		t.Fail()
		return
	}
	log.Printf("%+v", uri.SmartContract)
}
