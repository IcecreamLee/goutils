package goutils

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// ID生成器
type IDGenServ struct {
	Port int
	IDGenerator
}

var idGenServ *IDGenServ

var idHttpSrv *http.Server

var idGenServOnce sync.Once

// 获取 IDGenServ 单例
func IDServSingleton() *IDGenServ {
	idGenServOnce.Do(func() {
		idGenServ = &IDGenServ{}
		idGenServ.init()
		idGenServ.Port = 9999
	})
	return idGenServ
}

// 启用ID生成器Web服务
func (id *IDGenServ) Run() {
	fmt.Printf("ID Service: \n\nPort = %d \nEpoch = %d \nMachineId = %d \nMachineBit = %d \nSequence = %d "+
		"\nSequenceBit = %d", id.Port, id.Epoch, id.MachineId, id.MachineBit, id.Sequence, id.SequenceBit)

	idHttpSrv = &http.Server{Addr: ":" + strconv.Itoa(id.Port)}
	http.HandleFunc("/", id.response) // 设置访问的路由

	err := idHttpSrv.ListenAndServe()
	if err != nil {
		log.Fatal("Start ID Service Failure: ", err)
	}
}

// 停止ID生成器Web服务
func (id *IDGenServ) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := idHttpSrv.Shutdown(ctx); err != nil {
		log.Fatal("Stop ID Service Failure: ", err)
	}
	log.Println("Stopped ID Service")
}

// 生成ID响应输出
func (id *IDGenServ) response(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "%d", id.NextID()) // 这个写入到w的是输出到客户端的
	if err != nil {
		log.Fatal("ID Service Response Failure: ", err)
	}
}
