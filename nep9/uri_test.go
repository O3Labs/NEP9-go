package nep9_test

import (
	"testing"

	"github.com/o3labs/NEP9-go/nep9"
	"github.com/stretchr/testify/assert"
)

func TestURI(t *testing.T) {
	t.Run("NewURI()", func(t *testing.T) {
		t.Run("Valid", func(t *testing.T) {
			rawURI := "neo:AeNkbJdiMx49kBStQdDih7BzfDwyTNVRfb?asset=gas&amount=0.11&description=This%20is%20awesome&ecdh02=02ed53ad58c838435d4dd7a4b25c1eba01384c814ca53a539405434807afbb04b4"
			uri, err := nep9.NewURI(rawURI)

			assert.NoError(t, err)
			assert.IsType(t, &nep9.URI{}, uri)
		})

		t.Run("ValidWithAssetID", func(t *testing.T) {
			rawURI := "neo:AeNkbJdiMx49kBStQdDih7BzfDwyTNVRfb?asset=602c79718b16e442de58778e148d0b1084e3b2dffd5de6b7b16cee7969282de7&amount=0.11&description=This%20is%20awesome&ecdh02=02ed53ad58c838435d4dd7a4b25c1eba01384c814ca53a539405434807afbb04b4"
			uri, err := nep9.NewURI(rawURI)

			assert.NoError(t, err)
			assert.IsType(t, &nep9.URI{}, uri)
		})

		t.Run("InvalidScheme", func(t *testing.T) {
			rawURI := "bitcoin:AeNkbJdiMx49kBStQdDih7BzfDwyTNVRfb?asset=gas&amount=0.11"
			uri, err := nep9.NewURI(rawURI)

			assert.Error(t, err)
			assert.Nil(t, uri)
		})

		t.Run("InvalidAsset", func(t *testing.T) {
			rawURI := "neo:AeNkbJdiMx49kBStQdDih7BzfDwyTNVRfb?asset=unknown&amount=0.11"
			uri, err := nep9.NewURI(rawURI)

			assert.Error(t, err)
			assert.Nil(t, uri)
		})

		t.Run("InvalidAmount", func(t *testing.T) {
			rawURI := "neo:AeNkbJdiMx49kBStQdDih7BzfDwyTNVRfb?asset=gas&amount=foo"
			uri, err := nep9.NewURI(rawURI)

			assert.Error(t, err)
			assert.Nil(t, uri)
		})

		t.Run("ValidNEP5Transfer", func(t *testing.T) {
			rawURI := "neo:AeNkbJdiMx49kBStQdDih7BzfDwyTNVRfb?asset=ceab719b8baa2310f232ee0d277c061704541cfb&amount=10.0"
			uri, err := nep9.NewURI(rawURI)

			assert.NoError(t, err)
			assert.IsType(t, &nep9.URI{}, uri)
		})

		t.Run("InvalidNEP5Transfer", func(t *testing.T) {
			rawURI := "neo:AeNkbJdiMx49kBStQdDih7BzfDwyTNVRfb?asset=ce719b8baa2310f232ee0d277c061704541cfb&amount=10.0"
			uri, err := nep9.NewURI(rawURI)

			assert.Error(t, err)
			assert.Nil(t, uri)
		})

		t.Run("ValidNEP5TransferWithoutAmount", func(t *testing.T) {
			rawURI := "neo:AeNkbJdiMx49kBStQdDih7BzfDwyTNVRfb?asset=ceab719b8baa2310f232ee0d277c061704541cfb"
			uri, err := nep9.NewURI(rawURI)

			assert.NoError(t, err)
			assert.IsType(t, &nep9.URI{}, uri)
		})
	})
}
