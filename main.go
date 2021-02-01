package main

import(
	"gjek/service"
	"gjek/config"
	"gjek/controller"
	"gjek/dao"
	"gorm.io/gorm"
	"fmt"
)

func main(){
	e := config.InitWeb()
	g := initDb()
	service.SetService(g)
	dao.SetDao(g)

	fmt.Println("flag g")
	fmt.Println(g)
	fmt.Println("flag g")

	controller.SetInit(e)
	eg := config.SetJwt(e)
	controller.SetGsend(eg)

	e.Logger.Fatal(e.Start(":1234"))
}

func initDb() *gorm.DB {
	g, err := config.Conn()
	if err == nil {
		config.MigrateSchema(g)
		return g
	}
	panic(err)
}