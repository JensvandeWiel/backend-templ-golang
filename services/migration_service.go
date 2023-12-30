package services

import (
	"backend-templ-golang/app"
	"backend-templ-golang/models"
)

type MigrationService struct {
	a *app.App
}

func NewMigrationService(a *app.App) *MigrationService {
	return &MigrationService{a: a}
}

func (s *MigrationService) Migrate() {
	err := s.a.Db.AutoMigrate(&models.User{})
	if err != nil {
		s.a.Logger.Error(err)
	}
}
