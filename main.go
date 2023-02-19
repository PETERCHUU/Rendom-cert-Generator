package main

import (
	"linuxWebUI/controller"
	"linuxWebUI/model/config"
	"linuxWebUI/model/security/tlsgen"
	"runtime"
)

func main() {
	config := config.OS{}
	config.Type = runtime.GOOS
	config.CPU = runtime.GOARCH
	publicKeyblock, privatekeyblock := tlsgen.GenerateKey()
	controller.Run(publicKeyblock, privatekeyblock)
}
