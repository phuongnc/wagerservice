package wagerdto

type WagerReq struct {
	TotalWagerValue   uint    `json:"total_wager_value" valid:"required"`
	Odds              uint    `json:"odds" valid:"required"`
	SellingPercentage uint8   `json:"selling_percentage" valid:"range(1|100)"`
	SellingPrice      float32 `json:"selling_price" valid:"required, DecimalType, ValidSellingPrice"`
}

type WagerBuyingReq struct {
	WagerID     uint    `json:"wager_id"`
	BuyingPrice float32 `json:"buying_price" valid:"required, DecimalType"`
}
