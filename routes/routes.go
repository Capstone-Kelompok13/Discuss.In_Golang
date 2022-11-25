package routes

import (
	"database/sql"
	"discusiin/configs"
	"discusiin/repositories"
	tService "discusiin/services/topics"
	uService "discusiin/services/users"

	"gorm.io/gorm"
)

type Payload struct {
	Config  *configs.Config
	DBGorm  *gorm.DB
	DBSql   *sql.DB
	repoSql repositories.IDatabase
	// repoTSql repositories.ITopicDatabase
	uService uService.IUserServices
	tService tService.ITopicServices
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

func (p *Payload) GetTopicServices() tService.ITopicServices {
	if p.tService == nil {
		p.InitTopicService()
	}

	return p.tService
}

func (p *Payload) InitTopicService() {
	if p.repoSql == nil {
		p.InitRepoMysql()
	}

	p.tService = tService.NewTopicServices(p.repoSql)
}

// func (p *Payload) InitPocketMessageService() {
// 	if p.repoTSql == nil {
// 		p.InitRepoMysql()
// 	}

// 	p.tService = tService.NewTopicServices(p.repoTSql)
// }

// func (p *Payload) GetTopicServices() tService.ITopicServices {
// 	if p.tService == nil {
// 		p.InitTopicService()
// 	}
// 	return p.tService
// }
