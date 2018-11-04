package entity

import (
  "time"
)

type User struct {
  ID ID `json:"id"`
  Username string `json:"username"`
  Password string `json:"password"`
  Email string `json:"email"`
  Enabled bool `json:"enabled"`
  CreatedAt time.Time `json:"created_at"`
}
