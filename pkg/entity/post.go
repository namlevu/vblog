package entity

import (
  "time"
)

type Post struct {
  ID ID `json:"id"`
  Title string `json:"title"`
  Slug string `json:"slug"`
  Tags []string `json:"tags"`
  Content string `json:"content"`
  AuthorId ID `json:"author_id"`
  CreatedAt time.Time `json:"created_at"`
  CreatedBy ID `json:"created_by"`
  UpdatedAt time.Time `json:"updated_at"`
  UpdatedBy ID `json:"updated_by"`
}
