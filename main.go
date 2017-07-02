package main

import (
	"github.com/astaxie/beego"

	_ "github.com/BitAssetManagement/cms/src/routers"
	_ "github.com/BitAssetManagement/cms/src/service"
)

func main() {
	beego.SetLevel(beego.LevelDebug)
	beego.SetLogger("console", "")
	beego.Run()
}
