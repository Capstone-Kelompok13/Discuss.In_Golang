package routes

import (
	"database/sql"
	"discusiin/configs"
	"discusiin/repositories"
	uService "discusiin/services/users"

	"gorm.io/gorm"
)

type Payload struct {
	Config   *configs.Config
	DBGorm   *gorm.DB
	DBSql    *sql.DB
	repoSql  repositories.IDatabase
	uService uService.IUserServices
}

func (p *Payload) InitUserService() {
	if p.repoSql == nil {
		p.InitRepoMysql()
	}

	p.uService = uService.NewUserServices(p.repoSql)
}
func (p *Payload) InitPocketMessageService() {
	if p.repoSql == nil {
		p.InitRepoMysql()
	}

	p.uService = uService.NewUserServices(p.repoSql)
}

func (p *Payload) InitRepoMysql() {
	p.repoSql = repositories.NewGorm(p.DBGorm)
}

func (p *Payload) GetUserServices() uService.IUserServices {
	if p.uService == nil {
		p.InitUserService()
	}
	return p.uService
}
