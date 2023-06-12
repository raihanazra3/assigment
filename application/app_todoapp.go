package application

import (
	"assigment/domain_todocore/controller/restapi"
	"assigment/domain_todocore/gateway/withgorm"
	"assigment/domain_todocore/usecase/getalltodo"
	"assigment/domain_todocore/usecase/runtodocheck"
	"assigment/domain_todocore/usecase/runtodocreate"
	"assigment/shared/config"
	"assigment/shared/gogen"
	"assigment/shared/infrastructure/logger"
	"assigment/shared/infrastructure/token"
)

type todoapp struct{}

func NewTodoapp() gogen.Runner {
	return &todoapp{}
}

func (todoapp) Run() error {

	const appName = "todoapp"

	cfg := config.ReadConfig()

	appData := gogen.NewApplicationData(appName)

	log := logger.NewSimpleJSONLogger(appData)

	jwtToken := token.NewJWTToken(cfg.JWTSecretKey)

	datasource := withgorm.NewGateway(log, appData, cfg)

	primaryDriver := restapi.NewController(appData, log, cfg, jwtToken)

	primaryDriver.AddUsecase(
		//
		getalltodo.NewUsecase(datasource),
		runtodocheck.NewUsecase(datasource),
		runtodocreate.NewUsecase(datasource),
	)

	primaryDriver.RegisterRouter()

	primaryDriver.Start()

	return nil
}
