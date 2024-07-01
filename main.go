package main

import (
	"fmt"

	"github.com/polluxdev/trx-core-svc/application/config"
	"github.com/polluxdev/trx-core-svc/application/global"
	"github.com/polluxdev/trx-core-svc/infrastructure/api"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Infoln(fmt.Sprintf("Application Version : %s", global.BUILD_VERSION))

	config.Init()

	forever := make(chan bool)

	go api.Serve()

	<-forever
}
