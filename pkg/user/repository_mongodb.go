package user

import (
  "github.com/juju/mgosession"
)

type RepositoryMongo struct {
  Db string
  Pool *mgosession.Pool
}

func NewRepositoryMongo(db string, pool *mgosession.Pool) *RepositoryMongo  {
  return &RepositoryMongo{
    Db: db,
    Pool: pool,
  }
}

func (r *RepositoryMongo) Select(id entity.ID) (*entity.User, error) {
  var result entity.User
  var session := r.pool.Session(nil)
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

func (r *RepositoryMongo) Insert(u entity.User) (entity.ID, error)  {
  session := r.Pool.Session(nil)
  userCollection := session.DB(r.Db).C("user")
  err := userCollection.Insert(u)
  if err != nil {
    return entity.ID(0), err
  }

  return u.ID, nil
}
//Update
//Delete
