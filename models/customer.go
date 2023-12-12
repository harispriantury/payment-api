package models

import "time"

type Customer struct {
  Name  string `json:"name"`
  Username string `json:"username" gorm:"primary_key; not null; unique"`
  Password string `json:"password" gorm:"not null"`
  Phone string `json:"phone" gorm:"not null"`
  Token string `json:"token"`
  TokenExpiredAt time.Time `json:"token_expired_at"`
  Updated   int64 `gorm:"autoUpdateTime:milli"`// Use unix milli seconds as updating time
  Created   int64 `gorm:"autoCreateTime"`

}

type InputLogin struct {
  Username string `json:"username"`
  Password string `json:"password"`
}

type InputRegister struct {
  Name  string `json:"name"`
  Username string `json:"username"`
  Password string `json:"password"`
  Phone string `json:"phone" gorm:"not null"`
}