package main

import (
	"config_adapter/adapter"
	"github.com/urfave/cli"


)

func ModifySavePath(context *cli.Context) error {
	apkName := context.Args().First()
	adapter.ModifySavePath(apkName)
	return nil
}

func ModifyPackage(context *cli.Context) error {
	pkgName := context.Args().First()
	adapter.ModifyPackage(pkgName)
	return nil
}