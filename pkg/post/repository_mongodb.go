package post

import (
	"log"

	"github.com/juju/mgosession"
	mgo "gopkg.in/mgo.v2"
	//bson "gopkg.in/mgo.v2/bson"
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
  var post *entity.Post
  return post,nil
}
func (r *RepositoryMongo)SelectAll()([]*entity.Post, error) {
  // TODO:
	log.Println("Post Repository SelectAll")
	var posts []*entity.Post
	session := r.Pool.Session(nil)

	coll := session.DB(r.Db).C("post")

	err := coll.Find(nil).Sort("created_at").All(&posts)
	switch err {
	case nil:
		return posts, nil
	case mgo.ErrNotFound:
		log.Println("RepositoryMongo SelectAll ErrNotFound");
		return nil, entity.ErrNotFound
	default:
		log.Println("RepositoryMongo SelectAll default");
		return nil, err
	}

}
func (r *RepositoryMongo)Search(queryString string) ([]*entity.Post, error) {
  // TODO:
  return nil,nil
}
func (r *RepositoryMongo)Insert(p *entity.Post) (entity.ID, error) {
	log.Println("RepositoryMongo Insert called");
	session := r.Pool.Session(nil)
	postCollection := session.DB(r.Db).C("post")
	err := postCollection.Insert(p)
	if err != nil {
		log.Println("RepositoryMongo Insert Post failed");
		return entity.ID(0), err
	}

	return p.ID, nil
}
func (r *RepositoryMongo)Update(p *entity.Post) (entity.ID, error) {
  // TODO:
  return entity.ID(0), nil
}
func (r *RepositoryMongo)Delete(id entity.ID) error {
  // TODO:
  return nil
}
