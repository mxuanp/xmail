//utils包包含一些作者封装的用于本项目的工具
//sliceutil包含一些和slice相关的工具
package utils

import (
	"xmail/model"
)

func Remove(slice []model.User, item model.User) []model.User {
	if len(slice) == 0 {
		return nil
	}
	for i, s := range slice {
		if s == item {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
