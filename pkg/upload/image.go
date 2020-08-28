/*
 * @Author: wtf
 * @Date: 2020-08-27 21:31:00
 * @LastEditors: wtf
 * @LastEditTime: 2020-08-28 15:31:16
 * @Description: plase write Description
 */
package upload

import (
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"github.com/tengfei31/website/pkg/file"
	"github.com/tengfei31/website/pkg/logging"
	"github.com/tengfei31/website/pkg/setting"
	"github.com/tengfei31/website/pkg/util"
)

func GetImageFullUrl(name string) string {
	return fmt.Sprintf("%s/%s%s", setting.AppSetting.ImagePrefixUrl, GetImagePath(), name)
}

func GetImageName(name string) string {
	ext := path.Ext(name)
	filename := strings.TrimSuffix(name, ext)
	filename = util.EncodeMd5(filename)
	return filename + ext
}

func GetImagePath() string {
	return setting.AppSetting.ImageSavePath
}

func GetImageFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetImagePath()
}

func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allow := range setting.AppSetting.ImageAllowExts {
		if strings.ToUpper(allow) == strings.ToUpper(ext) {
			return true
		}
	}
	return false
}

func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		logging.Warn(err)
		return false
	}
	return size <= setting.AppSetting.ImageMaxSize
}

func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}
	err = file.IsNotExistMkDir(fmt.Sprintf("%s/%s", dir, src))
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}
	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}
	return nil
}

