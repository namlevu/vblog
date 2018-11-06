package auth

import (
	"vblog/pkg/entity"
)
//Service service interface
type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) Login(u *entity.User) (entity.Auth, error) {
	return s.repo.Login(u)
}
func (s *Service) Logout(id entity.ID) error {
  return s.repo.Logout(id)
}
