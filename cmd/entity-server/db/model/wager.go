package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Wager struct {
	gorm.Model
	TotalWagerValue     uint
	Odds                uint
	SellingPercentage   uint8
	SellingPrice        float32
	CurrentSellingPrice float32
	PercentageSold      *float32
	AmountSold          *uint

	WagerTransactions []*WagerTransaction
}

func (o *Wager) Insert(db *gorm.DB) error {
	if err := db.Create(o).Error; err != nil {
		return err
	}
	return nil
}

func (o *Wager) List(db *gorm.DB, filter map[string]interface{}) ([]*Wager, int16, error) {
	var list []*Wager
	order := "created_at DESC"
	if filter["order"] != "" && filter["order_by"] != "" {
		order = fmt.Sprintf("%s %s", filter["order_by"], filter["order"])
	}
	// Count the total first
	var total int64
	if err := db.Model(&Wager{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	// Query data by filter
	err := db.Order(order).Offset(filter["page_offset"].(int)).Limit(filter["page_limit"].(int)).Find(&list).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	return list, int16(total), nil
}

func (o *Wager) Get(db *gorm.DB) error {
	if err := db.Where("id = ?", o.ID).Find(&o).Error; err != nil {
		return err
	}
	return nil
}
