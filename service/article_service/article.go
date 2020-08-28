/*
 * @Author: wtf
 * @Date: 2020-08-28 23:16:33
 * @LastEditors: wtf
 * @LastEditTime: 2020-08-28 23:36:27
 * @Description: plase write Description
 */
package article_service

import "github.com/tengfei31/website/models"

type Article struct {
	Id int
	TagId int
	State int

}

func (a *Article) ExistById() (bool, error) {
	return models.ExistArticleById(a.Id), nil
}

func  (a *Article) Get() (models.Article, error) {
	article := models.GetArticle(a.Id)
	return article, nil
}


