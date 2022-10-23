package msg

var MsgFlags = map[int]string{
	SUCCESS:        "Success",
	ERROR:          "Fail",
	INVALID_PARAMS: "Invalid param",

	// WAGER
	WAGER_NOT_FOUND:            "Wager not found",
	BUYING_WAGER_PRICE_INVALID: "BuyingPrice must be lesser or equal to CurrentSellingPrice",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
