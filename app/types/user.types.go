package types

import "time"

type UserDTO struct {
	Username string    `json:"username"`
	Address  string    `json:"address"`
	TTL      time.Time `json:"ttl"`
}
