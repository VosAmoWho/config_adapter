package adapter

import (
	"os/exec"
	"path"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"config_adapter/utils"
)

var log = logrus.New()

const SAVE_PATH = "~/maplejava/output/"
const ConfigFilePath = "/Users/huawei/Test/"

func ModifySavePath(apkName string) {
	if err := modifyConfigSuffix(); err != nil {
		log.Error(err)
		return
	}

	v, err := getViperPoint()
	if err != nil {
		log.Error("err: get viper point failed")
		return
	}

	outPutPath := path.Join(SAVE_PATH, apkName)
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

	if err := recoverConfigSuffix(); err != nil {
		log.Error(err)
	}
}

func ModifyPackage(pkgName string) {
	if err := modifyConfigSuffix(); err != nil {
		log.Error(err)
		return
	}

	v, err := getViperPoint()
	if err != nil {
		log.Error("err: get viper point failed")
		return
	}

	pkgPath := path.Join(SAVE_PATH, pkgName)
	v.Set("package", pkgPath)
	if err := v.WriteConfig(); err != nil {
		log.Error(err)
	}

	if err := recoverConfigSuffix(); err != nil {
		log.Error(err)
	}
}

func modifyConfigSuffix() error {
	cmd := exec.Command(
		"mv",
		path.Join(ConfigFilePath, "config.conf"),
		path.Join(ConfigFilePath, "config"))
	return cmd.Run()
}

func recoverConfigSuffix() error {
	cmd := exec.Command(
		"mv",
		path.Join(ConfigFilePath, "config"),
		path.Join(ConfigFilePath, "config.conf"))
	return cmd.Run()
}

func getViperPoint() (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigFile(path.Join(ConfigFilePath, "config"))
	v.SetConfigType("properties")

	err := v.ReadInConfig()
	if err != nil {
		log.Errorf("read config file failed with error: %s", err)
		return nil, err
	}
	return v, nil
}
