package handler

import (
	"errors"
	"net/http"
	"strconv"
	"wagerservice/cmd/entity-server/registry"
	"wagerservice/cmd/entity-server/service/wager"
	"wagerservice/cmd/entity-server/service/wager/wagerdto"
	"wagerservice/internal/pkg/ginutil"
	"wagerservice/internal/pkg/msg"

	"github.com/gin-gonic/gin"
)

type WagerHandler struct {
	wagerService wager.WagerService
	provider     registry.Provider
}

func NewWagerHandler(
	wagerService wager.WagerService,
	provider registry.Provider,
) *WagerHandler {
	return &WagerHandler{
		wagerService: wagerService,
		provider:     provider,
	}
}

func (hdl *WagerHandler) Create(ctx *gin.Context) {
	appG := Gin{C: ctx}
	req := &wagerdto.WagerReq{}

	if isValid := appG.BindAndValidate(req); isValid {
		objRes, err := hdl.wagerService.Create(req)
		if err != nil {
			appG.Response(http.StatusBadRequest, false, nil, nil)
			return
		}
		appG.Response(http.StatusCreated, true, objRes, nil)
	}
}

func (hdl *WagerHandler) List(ctx *gin.Context) {
	appG := Gin{C: ctx}
	pageOffset, pageLimit := ginutil.GetPage(ctx, hdl.provider.Config.DefaultPageNum, hdl.provider.Config.DefaultPageLimit)
	filter := map[string]interface{}{
		"page_limit":  pageLimit,
		"page_offset": pageOffset,
		"order_by":    ctx.Query("order_by"),
		"order":       ctx.Query("order"),
	}

	objRes, total, err := hdl.wagerService.List(filter)
	if err != nil {
		appG.Response(http.StatusBadRequest, false, nil, nil)
		return
	}

	data := make(map[string]interface{})
	data["data"] = objRes
	data["total"] = total

	appG.Response(http.StatusOK, true, data, nil)
}

func (hdl *WagerHandler) Buy(ctx *gin.Context) {
	appG := Gin{C: ctx}

	wagerID, err := strconv.Atoi(ctx.Param("wager_id"))
	if err != nil {
		appG.Response(http.StatusBadRequest, false, nil, errors.New(msg.GetMsg(msg.INVALID_PARAMS)))
		return
	}

	req := &wagerdto.WagerBuyingReq{
		WagerID: uint(wagerID),
	}

	if isValid := appG.BindAndValidate(req); isValid {
		objRes, err := hdl.wagerService.Buy(req)
		if err != nil {
			appG.Response(http.StatusBadRequest, false, nil, err)
			return
		}
		appG.Response(http.StatusCreated, true, objRes, nil)
	}
}
