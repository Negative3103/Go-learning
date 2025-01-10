package coincap

import "fmt"

type assetsResponse struct {
	Data      []AssetData `json:"data"`
	Timestamp int64       `json:"timestamp"`
}

type assetResponse struct {
	Data      AssetData `json:"data"`
	Timestamp int64     `json:"timestamp"`
}

type AssetData struct {
	ID            string `json:"id"`
	Rank          string `json:"rank"`
	Symbol        string `json:"symbol"`
	Name          string `json:"name"`
	Supply        string `json:"supply"`
	MaxSupply     string `json:"maxSupply"`
	MarketCapUsd  string `json:"marketCapUsd"`
	VolumeUsd24Hr string `json:"volumeUsd24Hr"`
	PriceUsd      string `json:"priceUsd"`
}

func (d AssetData) Info() string {
	return fmt.Sprintf("[ID] %s | [Rank] %s | [Symbol] %s | [Name] %s | [Price] %s", d.ID, d.Rank, d.Symbol, d.Name, d.PriceUsd)
}
