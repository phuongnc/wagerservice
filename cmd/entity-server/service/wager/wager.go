package wager

import (
	"wagerservice/cmd/entity-server/db/model"
	"wagerservice/cmd/entity-server/service/wager/wagerdto"

	"gorm.io/gorm"
)

type WagerService interface {
	Create(req *wagerdto.WagerReq) (*wagerdto.WagerResp, error)
	List(filter map[string]interface{}) ([]wagerdto.WagerResp, int16, error)
	Buy(req *wagerdto.WagerBuyingReq) (*wagerdto.WagerBuyingResp, error)
}

func NewWagerService(
	db *gorm.DB,
) WagerService {
	return &wagerService{
		db: db,
	}
}

type wagerService struct {
	db *gorm.DB
}

func (sv *wagerService) Create(
	req *wagerdto.WagerReq,
) (*wagerdto.WagerResp, error) {
	objModel, err := wagerdto.CopyWagerReqToModel(req)
	if err != nil {
		return nil, err
	}
	if err = objModel.Insert(sv.db); err != nil {
		return nil, err
	}
	return wagerdto.CopyWagerModelToResp(objModel)
}

func (sv *wagerService) List(
	filter map[string]interface{},
) ([]wagerdto.WagerResp, int16, error) {
	list, total, err := (&model.Wager{}).List(sv.db, filter)
	if err != nil {
		return nil, 0, err
	}
	listResp, err := wagerdto.CopyListWagerModelToResp(list)
	if err != nil {
		return nil, 0, err
	}
	return listResp, total, err
}

func (sv *wagerService) Buy(
	req *wagerdto.WagerBuyingReq,
) (*wagerdto.WagerBuyingResp, error) {
	objModel, err := wagerdto.CopyReqToWagerTransactionModel(req)
	if err != nil {
		return nil, err
	}
	if err := objModel.Insert(sv.db); err != nil {
		return nil, err
	}
	return wagerdto.CopyWagerTransactionModelToResp(objModel)
}
