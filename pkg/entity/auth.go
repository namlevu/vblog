package entity

import (
  "time"
)

type Auth struct {
  ID ID `json:"id"`
  UserId ID `json:"user_id"`
  SessionId ID `json:"session_id"`
  ClientId ID `json:"client_id"`
  CreatedAt time.Time `json:"created_at"`
}
