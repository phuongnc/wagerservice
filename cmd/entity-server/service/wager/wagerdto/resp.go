package wagerdto

import (
	"time"
)

type WagerResp struct {
	ID                  uint      `json:"id"`
	TotalWagerValue     uint      `json:"total_wager_value"`
	Odds                uint      `json:"odds"`
	SellingPercentage   uint8     `json:"selling_percentage"`
	SellingPrice        float32   `json:"selling_price"`
	CurrentSellingPrice float32   `json:"current_selling_price"`
	PercentageSold      *float32  `json:"percentage_sold"`
	AmountSold          *uint     `json:"amount_sold"`
	PlaceAt             time.Time `json:"placed_at"`
}

type WagerBuyingResp struct {
	ID          uint      `json:"id"`
	WagerID     uint      `json:"wager_id"`
	BuyingPrice float32   `json:"buying_price"`
	BoughtAt    time.Time `json:"bought_at"`
}
