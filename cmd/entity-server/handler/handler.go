package handler

import (
	"fmt"
	"net/http"
	"strings"
	"wagerservice/cmd/entity-server/service/wager/wagerdto"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type ErrorRes struct {
	Error string `json:"error"`
}

type Gin struct {
	C *gin.Context
}

func init() {
	govalidator.CustomTypeTagMap.Set("ValidSellingPrice", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {
		switch v := context.(type) {
		case wagerdto.WagerReq:
			return v.SellingPrice > (float32(v.TotalWagerValue) * (float32(v.SellingPercentage) / 100))
		}
		return false
	}))

	govalidator.CustomTypeTagMap.Set("DecimalType", func(i interface{}, o interface{}) bool {
		return len(strings.Split(fmt.Sprintf("%v", i.(float32)), ".")[1]) <= 2
	})
}

func (g *Gin) Response(httpCode int, success bool, data interface{}, err error) {
	g.C.JSON(httpCode, gin.H{
		"success": success,
		"data":    data,
		"error":   err,
	})
	return
}

func (g *Gin) BindAndValidate(obj interface{}) bool {
	err := g.C.ShouldBind(obj)
	if err != nil {
		g.Response(http.StatusBadRequest, false, nil, err)
		return false
	}

	isValid, err := govalidator.ValidateStruct(obj)
	if err != nil || !isValid {
		g.Response(http.StatusBadRequest, false, nil, err)
		return false
	} else {
		return true
	}
}
