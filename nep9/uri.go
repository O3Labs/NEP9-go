package nep9

import (
	"fmt"
	"net/url"
	"strconv"
)

type (
	// URI is a Go representation of a NEP-9 parsed URI.
	URI struct {
		Address    string       `json:"address"`
		Amount     float64      `json:"amount,omitempty"`
		AssetID    string       `json:"assetID,omitempty"`
		Attributes []*Attribute `json:"attributes,omitempty"`
	}
)

const (
	validScheme = "neo"
)

var (
	validAssets = map[string]string{
		"c56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b": "neo",
		"602c79718b16e442de58778e148d0b1084e3b2dffd5de6b7b16cee7969282de7": "gas",
	}

	validURIKeys = map[string]string{
		"assetID": "assetID",
		"amount":  "amount",
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

	uri.Address = parsedURI.Opaque

	assetID := parsedURI.Query().Get("assetID")
	if assetID != "" {
		if validAssets[assetID] == "" {
			return nil, fmt.Errorf("Invalid AssetID, got: '%s'", assetID)
		}

		uri.AssetID = assetID
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
