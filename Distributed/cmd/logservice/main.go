package main

import (
	"context"
	"fmt"
	"github.com/MicroService-Su/Distributed/log"
	"github.com/MicroService-Su/Distributed/service"
	stlog "log"
)
func main() {
	log.Run("./distributed.log")
	host,port := "localhost","4000"
	ctx,err := service.Start(context.Background(),"Log Server", host, port, log.RegisterHandlers)
	if err != nil {
		stlog.Fatalln(err)
	}
	<- ctx.Done()
	fmt.Println("日志服务关闭")
}
