/*
 * @Author: wtf
 * @Date: 2020-08-28 23:16:33
 * @LastEditors: wtf
 * @LastEditTime: 2020-08-28 23:55:08
 * @Description: plase write Description
 */
package article_service

import (
	"encoding/json"

	"github.com/tengfei31/website/models"
	"github.com/tengfei31/website/pkg/gredis"
	"github.com/tengfei31/website/pkg/logging"
	"github.com/tengfei31/website/service/cache_service"
)

type Article struct {
	Id    int
	TagId int
	State int
}

func (a *Article) ExistById() (bool, error) {
	return models.ExistArticleById(a.Id), nil
}

func (a *Article) Get() (*models.Article, error) {
	var cacheAricle *models.Article

	cache := cache_service.Article{Id: a.Id}
	key := cache.GetArticleKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheAricle)
			return cacheAricle, nil
		}
	}
	article, err := models.GetArticle(a.Id)
	if err != nil {
		return nil, err
	}
	gredis.Set(key, article, 3600)
	return article, nil
}
