package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func Start(ctx context.Context, serviceName, host, port string, registerHandlers func())(context.Context,error){
	registerHandlers()
	ctx = startService(ctx,serviceName,host,port)
	return ctx,nil
}

func startService(ctx context.Context, serviceName, host, port string) context.Context{
	ctx, cancel := context.WithCancel(ctx)
	var srv http.Server
	srv.Addr = ":"+ port

	go func() {
		// 启动时发生错误就打印
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Printf("%v 启动中，按任意键停止。",serviceName)
		var s string
		_,err := fmt.Scanln(&s)
		if err != nil {
			fmt.Println(err)
		}
		err = srv.Shutdown(ctx)
		if err != nil {
			fmt.Println(err)
		}
		cancel()
	}()

	return ctx
}