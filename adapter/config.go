package adapter

import (
	"path"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"config_adapter/utils"
)

var log = logrus.New()

const SAVE_PATH = "maplejava/output/"
const ConfigFilePath = "MobilePerf/"

func ModifySavePath(apkName string) {

	v, err := getViperPoint()
	if err != nil {
		log.Error("err: get viper point failed")
		return
	}

	userPath, _ := utils.Home()
	outPutPath := path.Join(userPath,SAVE_PATH, apkName)
	v.Set("save_path", outPutPath)
	if err := v.WriteConfig(); err != nil {
		log.Error(err)
		return
	}

	log.Infof("Create %s output in path %s", apkName, outPutPath)
	if err := utils.CreatFile(outPutPath); err != nil {
		log.Error(err)
		return
	}
}

func ModifyPackage(pkgName string) {
	v, err := getViperPoint()
	if err != nil {
		log.Error("err: get viper point failed")
		return
	}

	v.Set("package", pkgName)
	if err := v.WriteConfig(); err != nil {
		log.Error(err)
	}
}

func getViperPoint() (*viper.Viper, error) {
	v := viper.New()
	userPath, _ := utils.Home()
	v.SetConfigFile(path.Join(userPath, ConfigFilePath, "config"))
	v.SetConfigType("properties")

	err := v.ReadInConfig()
	if err != nil {
		log.Errorf("read config file failed with error: %s", err)
		return nil, err
	}
	return v, nil
}
