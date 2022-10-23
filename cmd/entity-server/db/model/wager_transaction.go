package model

import (
	"errors"
	"wagerservice/internal/pkg/dbtx"
	"wagerservice/internal/pkg/msg"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type WagerTransaction struct {
	gorm.Model
	WagerID     uint
	BuyingPrice float32
}

func (o *WagerTransaction) Insert(db *gorm.DB) error {
	err := dbtx.TransactionDB(db, func(tx *gorm.DB) error {
		wager := &Wager{}
		if err := tx.Clauses(clause.Locking{
			Strength: "UPDATE",
			Options:  "NOWAIT",
		}).Where("id= ?", o.WagerID).First(&wager).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return errors.New(msg.GetMsg(msg.WAGER_NOT_FOUND))
			}
			return err
		}
		if o.BuyingPrice > wager.CurrentSellingPrice {
			return errors.New(msg.GetMsg(msg.BUYING_WAGER_PRICE_INVALID))
		}

		tx.Create(o)
		wager.CurrentSellingPrice = o.BuyingPrice
		var currentAmount uint
		if wager.AmountSold == nil {
			currentAmount = 1
		} else {
			currentAmount = *wager.AmountSold + 1
		}
		// TODO (Phuong): need confirm about how to caculate the PercentageSold
		percentageSold := (wager.CurrentSellingPrice / float32(wager.TotalWagerValue)) * 100

		update := make(map[string]interface{})
		update["current_selling_price"] = o.BuyingPrice
		update["amount_sold"] = currentAmount
		update["percentage_sold"] = percentageSold

		err := tx.Model(wager).Updates(update).Error
		return err
	})
	return err
}
