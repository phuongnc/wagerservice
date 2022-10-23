package wagerdto

import (
	"wagerservice/cmd/entity-server/db/model"
	"wagerservice/internal/pkg/copier"
)

func CopyWagerReqToModel(src *WagerReq) (dst *model.Wager, err error) {
	dst = &model.Wager{}
	if err = copier.CopyWithTransform(&dst, &src); err != nil {
		return nil, err
	}
	dst.CurrentSellingPrice = src.SellingPrice
	return
}

func CopyWagerModelToResp(src *model.Wager) (dst *WagerResp, err error) {
	dst = &WagerResp{}
	if err = copier.CopyWithTransform(&dst, &src); err != nil {
		return nil, err
	}
	dst.PlaceAt = src.CreatedAt
	return
}

func CopyListWagerModelToResp(src []*model.Wager) (dst []WagerResp, err error) {
	for _, srcObj := range src {
		dstObj := &WagerResp{}
		dstObj, err = CopyWagerModelToResp(srcObj)
		if err != nil {
			return nil, err
		}
		dst = append(dst, *dstObj)
	}
	return dst, nil
}

func CopyReqToWagerTransactionModel(src *WagerBuyingReq) (dst *model.WagerTransaction, err error) {
	dst = &model.WagerTransaction{}
	if err = copier.CopyWithTransform(dst, &src); err != nil {
		return nil, err
	}
	return
}

func CopyWagerTransactionModelToResp(src *model.WagerTransaction) (dst *WagerBuyingResp, err error) {
	dst = &WagerBuyingResp{}
	if err = copier.CopyWithTransform(&dst, &src); err != nil {
		return nil, err
	}
	dst.BoughtAt = src.CreatedAt
	return
}
