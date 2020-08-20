/*
 * @Author: wtf
 * @Date: 2020-08-19 21:01:58
 * @LastEditors: wtf
 * @LastEditTime: 2020-08-20 16:20:38
 * @Description: plase write Description
 */
package models

type Tag struct {
    Model

    Name string `json:"name"`
    CreatedBy string `json:"created_by"`
    ModifiedBy string `json:"modified_by"`
    State int `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface {}) (tags []Tag) {
    db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

    return
}

func GetTagTotal(maps interface {}) (count int){
    db.Model(&Tag{}).Where(maps).Count(&count)

    return
}

