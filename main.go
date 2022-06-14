package main

import (
	"main/common"
	"main/devices"
)

func main() {
	var config, _ = common.InitializeConfig() // TODO error handing

	devices.Handler(config)
}
