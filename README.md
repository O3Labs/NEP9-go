# NEP9-go
NEO URI Scheme parser in Go.

#### The [NEP9](https://github.com/neo-project/proposals/pull/25) proposal is not `final` yet.


## Installation

```sh
go get github.com/o3labs/nep9-go
```
### Model
```go
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

```
### Usage 

A URI of payment request to `AeNkbJdiMx49kBStQdDih7BzfDwyTNVRfb` with `0.11 GAS` and a description `for a coffee`

```go
uri := "neo:AeNkbJdiMx49kBStQdDih7BzfDwyTNVRfb?assetID=602c79718b16e442de58778e148d0b1084e3b2dffd5de6b7b16cee7969282de7&amount=0.11&description=for%20a%20coffee"
nep9, err := ParseNEP9URI(uri)
if err != nil {
  return
}
b, _ := json.Marshal(nep9)
fmt.Printf("%v", string(b))

```
