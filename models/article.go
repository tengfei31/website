/*
 * @Author: wtf
 * @Date: 2020-08-21 14:38:33
 * @LastEditors: wtf
 * @LastEditTime: 2020-08-27 17:06:09
 * @Description: plase write Description
 */
package models

import (
	"time"

	"github.com/jinzhu/gorm"
)


type Article struct {
	Model
	TagId int `json:"tag_id" gorm:"index"`
	Tag Tag `json:"tag"`
	Title string `json:"title"`
	Desc string `json:"desc"`
    Content string `json:"content"`
    CreatedBy string `json:"created_by"`
    ModifiedBy string `json:"modified_by"`
    State int `json:"state"`
}

func ExistArticleById(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)
	if article.ID > 0 {
		return true
	}
	return false
}

func GetArticleTotal(maps interface{}) int {
	var count int
	db.Model(&Article{}).Where(maps).Count(&count)
	return count
}

func GetArticles(pageNum int, pageSize int, maps interface{}) []Article {
	var articles []Article
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return articles
}

func GetArticle(id int) Article {
	var article Article
	db.Where("id = ?", id).First(&article)
	db.Model(&article).Related(&article.Tag)
	return article
}

func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Update(data)
	return true
}

func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article {
		TagId: data["tag_id"].(int),
		Title : data["title"].(string),
        Desc : data["desc"].(string),
        Content : data["content"].(string),
        CreatedBy : data["created_by"].(string),
        State : data["state"].(int),
	})
	return true
}

func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(Article{})
	return true
}




func (article *Article) BeforeCreate(scope *gorm.Scope) error {
    scope.SetColumn("CreatedOn", time.Now().Unix())

    return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
    scope.SetColumn("ModifiedOn", time.Now().Unix())

    return nil
}

func CleanAllArticle() bool {
	db.Unscoped().Where("deleted_on != ?", 0).Delete(&Article{})
	return true
}
