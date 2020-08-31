/*
 * @Author: wtf
 * @Date: 2020-08-28 17:39:00
 * @LastEditors: wtf
 * @LastEditTime: 2020-08-28 23:55:50
 * @Description: plase write Description
 */
package cache_service

import (
	"strconv"
	"strings"

	"github.com/tengfei31/website/pkg/e"
)

type Article struct {
	Id int
	TagId int 
	State int
	PageNum int
	PageSize int
}

func (a *Article) GetArticleKey() string {
	return e.CACHE_ARTICLE + strconv.Itoa(a.Id)
}

func (a *Article) GetArticlesKey() string {
	keys := []string{
		e.CACHE_ARTICLE,
		"LIST",
	}

	if a.Id > 0 {
		keys = append(keys, strconv.Itoa(a.Id))
	}
	if a.TagId > 0 {
		keys  = append(keys, strconv.Itoa(a.TagId))
	}
	if a.State > 0 {
		keys = append(keys, strconv.Itoa(a.State))
	}
	if a.PageNum > 0 {
		keys = append(keys, strconv.Itoa(a.PageNum))
	}
	if a.PageSize > 0 {
		keys = append(keys, strconv.Itoa(a.PageSize))
	}
	return strings.Join(keys, "_")
}


