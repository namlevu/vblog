package post

import (
	"log"
	//"time"

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

//----------------------------------------------------------------------------
// Implement interface
//----------------------------------------------------------------------------

func (s *Service)Select(id entity.ID) (*entity.Post, error) {
  // TODO:
  return s.repo.Select(id)
}
func (s *Service)SelectAll()([]*entity.Post, error) {
  // TODO:
	log.Println("Post Service SelectAll")
  return s.repo.SelectAll()
}
func (s *Service)Search(queryString string) ([]*entity.Post, error) {
  // TODO:
  return s.repo.Search(queryString)
}
func (s *Service)Insert(p *entity.Post) (entity.ID, error) {
  // TODO:
  return s.repo.Insert(p)
}
func (s *Service)Update(p *entity.Post) (entity.ID, error) {
  // TODO:
  return s.repo.Update(p)
}
func (s *Service)Delete(id entity.ID) error {
  // TODO:
  return s.repo.Delete(id)
}
