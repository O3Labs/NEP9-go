package nep9

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type (
	// URI is a Go representation of a NEP-9 parsed URI.
	URI struct {
		Address    string       `json:"address"`
		Amount     float64      `json:"amount,omitempty"`
		AssetID    string       `json:"assetID,omitempty"`
		Attributes []*Attribute `json:"attributes,omitempty"`

		// if the URI contains the script hash this field will be filled
		SmartContract *SmartContract `json:"smartContract"` //optional
	}

	SmartContract struct {
		ScriptHash string        `json:"scriptHash"`
		Operation  string        `json:"operation"`
		Params     []interface{} `json:"params"`
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
	//NEO address should be 34 characters
	//Smart contract's script hash should be 40 characters

	//simple check here
	isNativeAssetTransfer := true
	if len(parsedURI.Opaque) == 34 {
		valid := ValidateNEOAddress(parsedURI.Opaque)
		if valid == false {
			return nil, fmt.Errorf("%v is invalid NEOAddress", parsedURI.Opaque)
		}
		isNativeAssetTransfer = true
	} else {
		valid := ValidateSmartContractScriptHash(parsedURI.Opaque)
		if valid == false {
			return nil, fmt.Errorf("%v is invalid smart contract script hash", parsedURI.Opaque)
		}

		isNativeAssetTransfer = false
	}

	//Smart Contract stuff
	if isNativeAssetTransfer == false {

		smartContract := SmartContract{
			ScriptHash: parsedURI.Opaque,
		}

		operation := parsedURI.Query().Get("operation")
		if operation == "" {
			return nil, fmt.Errorf("Expected SmartContract operation")
		}

		smartContract.Operation = operation

		//eventually we will need to parse the ABI and check the params here
		paramsString := parsedURI.Query().Get("params")
		if paramsString != "" {
			params := strings.Split(paramsString, ",")
			for _, param := range params {
				v := strings.TrimSpace(param)
				if v != "" {
					smartContract.Params = append(smartContract.Params, param)
				}
			}
		}

		uri.SmartContract = &smartContract

		return &uri, nil
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
