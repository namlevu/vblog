package user

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

func (r *RepositoryMongo) Select(id entity.ID) (*entity.User, error) {
	var result entity.User
	session := r.Pool.Session(nil)
	userCollection := session.DB(r.Db).C("user")

	err := userCollection.Find(bson.M{"_id": id}).One(&result)

	switch err {
	case nil:
		return &result, nil
	case mgo.ErrNotFound:
		return nil, entity.ErrNotFound
	default:
		return nil, err

	}
}

func (r *RepositoryMongo) Search(query string) ([]*entity.User, error) {
	var users []*entity.User
	session := r.Pool.Session(nil)
	coll := session.DB(r.Db).C("user")
	err := coll.Find(bson.M{"username": &bson.RegEx{Pattern: query, Options: "i"}}).Limit(10).Sort("name").All(&users)
	switch err {
	case nil:
		return users, nil
	case mgo.ErrNotFound:
		return nil, entity.ErrNotFound
	default:
		return nil, err
	}
}

func (r *RepositoryMongo) Insert(u *entity.User) (entity.ID, error) {
	log.Println("RepositoryMongo Insert called");
	session := r.Pool.Session(nil)
	userCollection := session.DB(r.Db).C("user")
	err := userCollection.Insert(u)
	if err != nil {
		log.Println("RepositoryMongo Insert Insert failed");
		return entity.ID(0), err
	}

	return u.ID, nil
}

//Update
func (r *RepositoryMongo)Update(u *entity.User) (entity.ID, error) {
  // TODO:
  return entity.NewID(), nil
}
//Delete
func (r *RepositoryMongo)Delete(id entity.ID) error {
  // TODO:
  return nil
}
//
func (r *RepositoryMongo)SelectAll()([]*entity.User, error) {
  // TODO:
	log.Println("RepositoryMongo SelectAll called");
	var users []*entity.User
	session := r.Pool.Session(nil)
	log.Println(session)
	coll := session.DB(r.Db).C("user")
	log.Println(coll)
	err := coll.Find(nil).Sort("username").All(&users)
	switch err {
	case nil:
		return users, nil
	case mgo.ErrNotFound:
		log.Println("RepositoryMongo SelectAll ErrNotFound");
		return nil, entity.ErrNotFound
	default:
		log.Println("RepositoryMongo SelectAll default");
		return nil, err
	}
}
