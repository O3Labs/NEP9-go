package nep9

type (
	// Attribute is an optional field that can be included within a NEP-9 URI.
	Attribute struct {
		Key                  string   `json:"key"`
		TransactionAttribute string   `json:"transaction_attribute"`
		Value                []string `json:"value"`
	}
)

var (
	transactionAttributes = map[string]string{
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
)

// NewAttribute creates a new Attribute.
func NewAttribute(key string, value []string) *Attribute {
	byteValue := transactionAttributes[key]
	if byteValue == "" {
		return nil
	}

	return &Attribute{
		Key:                  key,
		TransactionAttribute: byteValue,
		Value:                value,
	}
}
