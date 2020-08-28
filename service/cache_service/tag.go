/*
 * @Author: wtf
 * @Date: 2020-08-28 17:39:12
 * @LastEditors: wtf
 * @LastEditTime: 2020-08-28 19:09:58
 * @Description: plase write Description
 */
package cache_service

import (
	"strings"

	"github.com/tengfei31/website/pkg/e"
)

type Tag struct {
	Id int
	Name int
	State int
	PageNum int
	PageSize int
}

func (t *Tag) GetTagsKey() string {
	keys := []string{
		e.CACHE_TAG,
		"LIST"
	}
	if t.Name != "" {
		keys = append(keys, t.Name)
	}
	if t.State >= 0 {
		keys = append(keys, strconv.Itoa(t.State))
	}
	if t.PageNum > 0 {
		keys = append(keys, strconv.Itoa(t.PageNum))
	}
	if t.PageSize > 0 {
		keys = append(keys, strconv.Itoa(t.PageSize))
	}
	return strings.Join(keys, "_")
}



