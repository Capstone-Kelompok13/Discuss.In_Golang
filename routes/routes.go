package routes

import (
	"database/sql"
	"discusiin/configs"
	"discusiin/repositories"
	cService "discusiin/services/comments"
	pService "discusiin/services/posts"
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
	pService pService.IPostServices
	cService cService.ICommentServices
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

// Topic -----------------------------------------------------------------------------------------------------------------
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

// Post -----------------------------------------------------------------------------------------------------------------
func (p *Payload) GetPostServices() pService.IPostServices {
	if p.pService == nil {
		p.InitPostService()
	}

	return p.pService
}

func (p *Payload) InitPostService() {
	if p.repoSql == nil {
		p.InitRepoMysql()
	}

	p.pService = pService.NewPostServices(p.repoSql)
}

// Comment -----------------------------------------------------------------------------------------------------------------
func (p *Payload) GetCommentServices() cService.ICommentServices {
	if p.cService == nil {
		p.InitCommentService()
	}

	return p.cService
}

func (p *Payload) InitCommentService() {
	if p.repoSql == nil {
		p.InitRepoMysql()
	}

	p.cService = cService.NewCommentServices(p.repoSql)
}
