package models

type Payment struct {
	Key    string `json:"key"`
	External_reference      int    `json:"external_reference"`
	Rate float32 `json:"rate"`
	Status string `json:"status"`
}

type Payments []Payment
