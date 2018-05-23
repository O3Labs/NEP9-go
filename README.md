# NEP9-go
NEO URI Scheme parser in Go.

#### The [NEP9](https://github.com/neo-project/proposals/pull/25) proposal is not `final` yet.


## Installation

```sh
go get github.com/o3labs/nep9-go
```
### Model
```go
type Attribute struct {
	Key                  string   `json:"key"`
	TransactionAttribute string   `json:"transaction_attribute"`
	Value                []string `json:"value"`
}

type URI struct {
	Address    string       `json:"address"`
	Amount     float64      `json:"amount,omitempty"`
	Asset      string       `json:"asset,omitempty"`
	Attributes []*Attribute `json:"attributes,omitempty"`
}

```
### Usage 

A URI of payment request to `AeNkbJdiMx49kBStQdDih7BzfDwyTNVRfb` with `0.11 GAS` and a description `for a coffee`

```go
rawURI := "neo:AeNkbJdiMx49kBStQdDih7BzfDwyTNVRfb?asset=gas&amount=0.11&description=for%20a%20coffee"
nep9, err := nep9.NewURI(rawURI)
if err != nil {
  return
}
b, _ := json.Marshal(nep9)
fmt.Printf("%v", string(b))

```


A URI of NEP-5 token request to `AeNkbJdiMx49kBStQdDih7BzfDwyTNVRfb` with `2 ONT`  
A script hash of ONT is `ceab719b8baa2310f232ee0d277c061704541cfb`
```go
rawURI := "neo:AeNkbJdiMx49kBStQdDih7BzfDwyTNVRfb?asset=ceab719b8baa2310f232ee0d277c061704541cfb&amount=2"
nep9, err := nep9.NewURI(rawURI)
if err != nil {
  return
}
b, _ := json.Marshal(nep9)
fmt.Printf("%v", string(b))

```