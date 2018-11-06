package auth

import (
	"log"
	"time"
  "errors"

	"github.com/juju/mgosession"
	//mgo "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
	"golang.org/x/crypto/bcrypt"

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
//
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//-----------------------------------------------------------------------------
func (r *RepositoryMongo) Login(u *entity.User) (entity.Auth, error) {
  var auth entity.Auth
  var user entity.User
  session := r.Pool.Session(nil)
	userCollection := session.DB(r.Db).C("user")
	err := userCollection.Find(bson.M{"username": u.Username}).One(&user)
  if !CheckPasswordHash(u.Password, user.Password) {
    // password incorrect
    return entity.Auth{}, errors.New("Password incorrect")
  }
  auth.ID = entity.NewID()
  auth.UserId = user.ID
  auth.SessionId = entity.NewID()
  auth.ClientId = entity.NewID()
	auth.CreatedAt = time.Now()

  authCollection := session.DB(r.Db).C("auth")
	err = authCollection.Insert(auth)
	if err != nil {
		log.Println("RepositoryMongo Insert Auth failed");
		return entity.Auth{}, err
	}

  return auth, nil
}
func (r *RepositoryMongo) Logout(id entity.ID) error {
	log.Println("RepositoryMongo auth logout")
	log.Println(id)
  session := r.Pool.Session(nil)
	coll := session.DB(r.Db).C("auth")
	log.Println(coll)
	return coll.Remove(bson.M{"_id": id})
}
