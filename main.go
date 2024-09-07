package main

import (
	"fmt"
	_ "server/models"
	_ "server/routers"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
)

// initalize database connection
func init() {
	// Register PostgreSQL driver
	orm.RegisterDriver("postgres", orm.DRPostgres)
	// check db connection
	connStr, err := beego.AppConfig.String("pgurl")

	if err != nil {
		fmt.Println("Failed to connect DB")
	}
	orm.RegisterDataBase("default", "postgres", connStr)
}

func main() {
	orm.RunSyncdb("default", true, true)
	beego.Run()

}
