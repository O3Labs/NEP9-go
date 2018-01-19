package nep9

import (
	"errors"
	"net/url"
	"strconv"
)

const (
	scheme = "neo"
)

var validURIKeys = map[string]string{
	"assetID": "assetID",
	"amount":  "amount",
}

var validAssets = map[string]string{
	"c56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b": "neo",
	"602c79718b16e442de58778e148d0b1084e3b2dffd5de6b7b16cee7969282de7": "gas",
}

type NEP9Attribute struct {
	Key                     string   `json:"key"`
	NEOTransactionAttribute string   `json:"attribute"`
	Value                   []string `json:"value"`
}
type NEP9 struct {
	To         string          `json:"to"`
	AssetID    string          `json:"assetID"`
	Amount     float64         `json:"amount"`
	Attributes []NEP9Attribute `json:"attributes,omitempty"`
}

var transactionAttributes = map[string]string{
	"contractHash":   "0x00",
	"ecdh02":         "0x02",
	"ecdh03":         "0x03",
	"script":         "0x20",
	"vote":           "0x30",
	"certUrl":        "0x80",
	"descriptionUrl": "0x81",
	"description":    "0x90",
	"hash1":          "0xa1",
	"hash2":          "0xa2",
	"hash3":          "0xa3",
	"hash4":          "0xa4",
	"hash5":          "0xa5",
	"hash6":          "0xa6",
	"hash7":          "0xa7",
	"hash8":          "0xa8",
	"hash9":          "0xa9",
	"hash10":         "0xaa",
	"hash11":         "0xab",
	"hash12":         "0xac",
	"hash13":         "0xad",
	"hash14":         "0xae",
	"hash15":         "0xaf",
	"remark1":        "0xf1",
	"remark2":        "0xf2",
	"remark3":        "0xf3",
	"remark4":        "0xf4",
	"remark5":        "0xf5",
	"remark6":        "0xf6",
	"remark7":        "0xf7",
	"remark8":        "0xf8",
	"remark9":        "0xf9",
	"remark10":       "0xfa",
	"remark11":       "0xfb",
	"remark12":       "0xfc",
	"remark13":       "0xfd",
	"remark14":       "0xfe",
	"remark15":       "0xff",
}

func ParseNEP9URI(uri string) (*NEP9, error) {
	out := &NEP9{}
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}
	if u.Scheme != scheme {
		return nil, errors.New("invalid scheme")
	}
	assetID := u.Query().Get("assetID")

	if validAssets[assetID] == "" {
		return nil, errors.New("Invalid Asset")
	}
	out.AssetID = assetID
	out.To = u.Opaque
	if f, err := strconv.ParseFloat(u.Query().Get("amount"), 64); err == nil {
		out.Amount = f
	}

	//parse transactions attributes
	for key, value := range u.Query() {

		if validURIKeys[key] != "" {
			continue
		}

		byteValue := transactionAttributes[key]
		if byteValue == "" {
			continue
		}
		out.Attributes = append(out.Attributes, NEP9Attribute{
			Key: key,
			NEOTransactionAttribute: byteValue,
			Value: value,
		})
	}

	return out, nil
}
