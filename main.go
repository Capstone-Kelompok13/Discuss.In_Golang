package main

import (
	"discusiin/configs"
	"discusiin/routes"
	v1 "discusiin/routes/v1"
)

// @title Diskusi.in API
// @version 1.0
// @description Dokumentasi API Diskusi.in.

// @host 35.78.120.202
// @BasePath /api/v1
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

	err := e.Start(configs.Cfg.APIPort)
	if err != nil {
		panic(err)
	}
}
