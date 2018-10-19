package nep9

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type (
	// URI is a Go representation of a NEP-9 parsed URI.
	// Support unified native asset and NEP-5 token transfer
	URI struct {
		Address    string       `json:"address"`
		Amount     float64      `json:"amount,omitempty"`
		Asset      string       `json:"asset,omitempty"`
		Attributes []*Attribute `json:"attributes,omitempty"`
	}
)

const (
	validScheme = "neo"
)

var (
	validAssetAlias = map[string]string{
		"neo": "c56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b",
		"gas": "602c79718b16e442de58778e148d0b1084e3b2dffd5de6b7b16cee7969282de7",
	}
	validAssetIDs = map[string]string{
		"c56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b": "neo",
		"602c79718b16e442de58778e148d0b1084e3b2dffd5de6b7b16cee7969282de7": "gas",
	}

	validURIKeys = map[string]string{
		"asset":  "asset",
		"amount": "amount",
	}
)

// NewURI creates a new URI.
func NewURI(rawURI string) (*URI, error) {
	var uri URI

	parsedURI, err := url.Parse(rawURI)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse raw URI argument: '%s'", err)
	}

	if parsedURI.Scheme != validScheme {
		return nil, fmt.Errorf(
			"Invalid URI scheme, expecting '%s' but got: '%s'", validScheme, parsedURI.Scheme,
		)
	}
	//NEO address should be 34 characters
	if len(parsedURI.Opaque) != 34 {
		return nil, fmt.Errorf("%v is invalid NEOAddress", parsedURI.Opaque)
	}

	//validate using base58 again
	valid := ValidateNEOAddress(parsedURI.Opaque)
	if valid == false {
		return nil, fmt.Errorf("%v is invalid NEOAddress", parsedURI.Opaque)
	}

	uri.Address = parsedURI.Opaque

	//check both assetID and alias(neo|gas)
	assetString := parsedURI.Query().Get("asset")

	//if the asset in query string is not empty
	if assetString != "" {
		//check for both assetID and asset alias
		assetID := validAssetAlias[strings.ToLower(assetString)]
		assetName := validAssetIDs[assetString]
		//if both are empty we then try to validate the nep5 script hash again
		if assetID == "" && assetName == "" {
			validScriptHash := ValidateSmartContractScriptHash(assetString)
			//if it's invalid then return here
			if validScriptHash == false {
				return nil, fmt.Errorf("Invalid asset, got: '%s'", assetString)
			}
		}
		//if it gets here meaning that the asset query string is valid
		uri.Asset = assetString
	}

	stringAmount := parsedURI.Query().Get("amount")
	if stringAmount != "" {
		amount, err := strconv.ParseFloat(stringAmount, 64)
		if err != nil {
			return nil, fmt.Errorf("Invalid Amount format for '%s': '%s'", stringAmount, err)
		}

		uri.Amount = amount
	}

	for key, value := range parsedURI.Query() {
		if validURIKeys[key] != "" {
			continue
		}

		uri.Attributes = append(uri.Attributes, NewAttribute(key, value))
	}

	return &uri, nil
}
