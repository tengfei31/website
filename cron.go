/*
 * @Author: wtf
 * @Date: 2020-08-27 17:07:00
 * @LastEditors: wtf
 * @LastEditTime: 2020-08-31 17:22:25
 * @Description: plase write Description
 */
package main

import (
	"time"

	"github.com/robfig/cron"
	"github.com/tengfei31/website/models"
	"github.com/tengfei31/website/pkg/logging"
	"github.com/tengfei31/website/pkg/setting"
)

func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()

	logging.Info("cron starting...")
	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		logging.Info("run models.CleanAllArticle...")
		models.CleanAllArticle()
	})
	c.AddFunc("* * * * * *", func() {
		logging.Info("run models.CleanAllTag...")
		models.CleanAllTag()
	})
	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
