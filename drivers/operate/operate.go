package operate

import (
	"github.com/Xhofe/alist/drivers/base"
	"github.com/Xhofe/alist/model"
	"github.com/Xhofe/alist/utils"
	log "github.com/sirupsen/logrus"
)

func MakeDir(driver base.Driver, account *model.Account, path string, clearCache bool) error {
	log.Debugf("mkdir: %s", path)
	err := driver.MakeDir(path, account)
	if err == nil && clearCache {
		_ = base.DeleteCache(utils.Dir(path), account)
	}
	if err != nil {
		log.Errorf("mkdir error: %s", err.Error())
	}
	return err
}

func Move(driver base.Driver, account *model.Account, src, dst string, clearCache bool) error {
	log.Debugf("move %s to %s", src, dst)
	rename := false
	if utils.Dir(src) == utils.Dir(dst) {
		rename = true
	}
	var err error
	if rename {
		err = driver.Rename(src, dst, account)
	} else {
		err = driver.Move(src, dst, account)
	}
	if err == nil && clearCache {
		_ = base.DeleteCache(utils.Dir(src), account)
		if !rename {
			_ = base.DeleteCache(utils.Dir(dst), account)
		}
	}
	if err != nil {
		log.Errorf("move error: %s", err.Error())
	}
	return err
}

func Copy(driver base.Driver, account *model.Account, src, dst string, clearCache bool) error {
	log.Debugf("copy %s to %s", src, dst)
	err := driver.Copy(src, dst, account)
	if err == nil && clearCache {
		_ = base.DeleteCache(utils.Dir(dst), account)
	}
	if err != nil {
		log.Errorf("copy error: %s", err.Error())
	}
	return err
}

func Delete(driver base.Driver, account *model.Account, path string, clearCache bool) error {
	log.Debugf("delete %s", path)
	err := driver.Delete(path, account)
	if err == nil && clearCache {
		_ = base.DeleteCache(utils.Dir(path), account)
	}
	if err != nil {
		log.Errorf("delete error: %s", err.Error())
	}
	return err
}

func Upload(driver base.Driver, account *model.Account, file *model.FileStream, clearCache bool) error {
	defer func() {
		_ = file.Close()
	}()
	err := driver.Upload(file, account)
	if err == nil && clearCache {
		_ = base.DeleteCache(file.ParentPath, account)
	}
	if err != nil {
		log.Errorf("upload error: %s", err.Error())
	}
	return err
}
