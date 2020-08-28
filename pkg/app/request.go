/*
 * @Author: wtf
 * @Date: 2020-08-28 22:56:08
 * @LastEditors: wtf
 * @LastEditTime: 2020-08-28 22:57:34
 * @Description: plase write Description
 */
package app

import (
	"github.com/astaxie/beego/validation"
	"github.com/tengfei31/website/pkg/logging"
)

func MakeErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Info(err.Key, err.Message)
	}
	return
}


