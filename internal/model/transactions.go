package model

import "time"

type Transactions struct {
	User
	ID                 uint      `json:"id"`
	Sum                float64   `json:"sum"`
	Currency           string    `json:"currency"`
	DateTimeCreate     time.Time `json:"date-time-create"`
	DateTimeLastChange time.Time `json:"date-time-last-change"`
	Status             string    `json:"status"`
}

type User struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}
