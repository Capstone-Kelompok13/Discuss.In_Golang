package main

import (
	"discusiin/configs"
	"discusiin/routes"
	v1 "discusiin/routes/v1"
)

func main() {

	configs.InitConfig()
	configs.InitDatabase()

	routePayload := &routes.Payload{
		DBGorm: configs.DB,
		Config: configs.Cfg,
	}

	routePayload.InitUserService()

	e, trace := v1.InitRoute(routePayload)
	defer trace.Close()

	e.Logger.Fatal(e.Start(configs.Cfg.APIPort))
}
