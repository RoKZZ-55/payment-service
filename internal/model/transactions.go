package model

type Transactions struct {
	TransactID         uint64  `json:"transact-id"`
	UserID             uint64  `json:"user-id"`
	Email              string  `json:"email"`
	Sum                float64 `json:"sum"`
	Currency           string  `json:"currency"`
	DateTimeCreate     string  `json:"date-time-create"`
	DateTimeLastChange string  `json:"date-time-last-change"`
	Status             string  `json:"status"`
}
