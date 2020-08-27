/*
 * @Author: wtf
 * @Date: 2020-08-27 21:28:26
 * @LastEditors: wtf
 * @LastEditTime: 2020-08-27 21:30:36
 * @Description: plase write Description
 */
package util

import (
	"crypto/md5"
	"encoding/hex"
)

func EncodeMd5(val string) string {
	m := md5.New()
	m.Write([]byte(val))
	return hex.EncodeToString(m.Sum(nil))
}


