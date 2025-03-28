package main

import (
	"github.com/cloudwego/kitex/server"
	"link_shorten_server/init_db"
	link "link_shorten_server/link/kitex_gen/link/linkservice"
	"link_shorten_server/redis_client"
	"log"
	"net"
)

func main() {
	//1.连接数据库
	err := init_db.ConnectDB()
	if err != nil {
		log.Fatalf("dao.ConnectDB error: %v", err)
	}
	//2.连接redis
	err = redis_client.InitRedis()
	if err != nil {
		log.Fatalf("redis_client.InitRedis error: %v", err)
	}
	//3.启动服务
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8890")
	if err != nil {
		log.Fatalf("net.ResolveTCPAddr error: %v", err)
	}
	svr := link.NewServer(new(LinkServiceImpl), server.WithServiceAddr(addr))
	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
