package token

import "time"

// need to implement all the methods in this interface to implement
type Maker interface {
	// creates new token
	CreateToken(username string, duration time.Duration) (string, *Payload, error)

	//verify the token
	VerifyToken(token string) (*Payload, error)
}
