package models

import "time"

type Merchant struct {
  Name  string `json:"name"`
  Username string `json:"username" gorm:"primary_key; not null; unique"`
  Password string `json:"password" gorm:"not null"`
  Address string `json:"address"`
  Phone string `json:"phone"`
  Token string `json:"token"`
  TokenExpiredAt time.Time `json:"token_expired_at"`
  Updated   int64 `gorm:"autoUpdateTime:milli"`// Use unix milli seconds as updating time
  Created   int64 `gorm:"autoCreateTime"`

}

type InputMerchant struct {
  Username string `json:"username"`
  Password string `json:"password"`
}

type InputRegisterMerchant struct {
  Name  string `json:"name"`
  Username string `json:"username"`
  Password string `json:"password"`
  Phone string `json:"phone" gorm:"not null"`
}
