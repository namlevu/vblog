package post

import (
	"log"

	"github.com/juju/mgosession"
	mgo "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
	"vblog/pkg/entity"
)

type RepositoryMongo struct {
	Db   string
	Pool *mgosession.Pool
}

func NewRepositoryMongo(db string, pool *mgosession.Pool) *RepositoryMongo {
	return &RepositoryMongo{
		Db:   db,
		Pool: pool,
	}
}

//----------------------------------------------------------------------------
// Implement interface
//----------------------------------------------------------------------------

func (r *RepositoryMongo)Select(id entity.ID) (*entity.Post, error) {
  // TODO:
  var post entity.Post{}
  return post,nil
}
func (r *RepositoryMongo)SelectAll()([]*entity.Post, error) {
  // TODO:
  return nil,nil
}
func (r *RepositoryMongo)Search(queryString string) ([]*entity.Post, error) {
  // TODO:
  return nil,nil
}
func (r *RepositoryMongo)Insert(p *entity.Post) (entity.ID, error) {
  // TODO:
  return entity.ID(0), nil
}
func (r *RepositoryMongo)Update(p *entity.Post) (entity.ID, error) {
  // TODO:
  return entity.ID(0), nil
}
func (r *RepositoryMongo)Delete(id entity.ID) error {
  // TODO:
  return nil
}
