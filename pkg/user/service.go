package user

import (
  "time"
  "log"
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


func (s *Service) Insert(u *entity.User) (entity.ID, error) {
	u.ID = entity.NewID()
	u.CreatedAt = time.Now()
	return s.repo.Insert(u)
}
//FindAll bookmarks
func (s *Service) SelectAll() ([]*entity.User, error) {
  log.Println("Service SelectAll called");
	return s.repo.SelectAll()
}

func (s *Service) Search(query string) ([]*entity.User, error) {
	return s.repo.Search(query)
}

//Update
func (s *Service)Update(u *entity.User) (entity.ID, error) {
  // TODO:
  return entity.NewID(), nil
}
//Delete
func (s *Service)Delete(id entity.ID) error {
  // TODO:
  return nil
}
//Select
func (s *Service)Select(id entity.ID) (*entity.User, error) {
  // TODO:
  return nil,nil
}
