package main

import (
	"user-service/user"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hashicorp/consul/api"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connectDB() *gorm.DB {
	// 格式：用户名:密码@协议(主机:端口)/数据库名?参数
	dsn := "root:123456@tcp(localhost:3306)/tiktok_ecommerce?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}
	return db
}
func main() {
	h := server.Default()
	db := connectDB() // 初始化数据库连接

	// 注册服务到Consul
	consulConfig := api.DefaultConfig()
	consulClient, _ := api.NewClient(consulConfig)
	registration := &api.AgentServiceRegistration{
		Name:    "user-service",
		Port:    8080,
		Address: "localhost",
	}
	consulClient.Agent().ServiceRegister(registration)

	// 绑定用户服务路由
	userHandler := user.NewUserHandler(db)
	h.POST("/user/register", userHandler.Register)
	h.POST("/user/login", userHandler.Login)

	h.Spin()
}
