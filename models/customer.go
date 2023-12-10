package models

type Customer struct {
  Name  string `json:"name"`
  Username string `json:"username" gorm:"primary_key; not null; unique"`
  Password string `json:"password" gorm:"not null"`
  Token string `json:"token" gorm:"unique"`
  TokenExpiredAt int `json:"token_expired_at"`
}

type InputLogin struct {
  Username string `json:"username"`
  Password string `json:"password"`
}