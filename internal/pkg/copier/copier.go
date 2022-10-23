package copier

import (
	jinzhuCopier "github.com/jinzhu/copier"
)

func CopyWithTransform(toValue interface{}, fromValue interface{}) (err error) {
	return jinzhuCopier.Copy(toValue, fromValue)
}
