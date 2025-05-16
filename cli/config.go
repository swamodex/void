package main

import (
	void "void/app"

	"github.com/cosmos/cosmos-sdk/version"
)

func CliConfig() {
	version.Name = void.AppName
	version.AppName = void.AppName
	version.Version = void.AppVersion
}
