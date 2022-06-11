package model

type Transactions struct {
	UserID             uint
	TransactID         uint
	Email              string
	Sum                float32
	Currency           string
	DateTimeCreate     string
	DateTimeLastChange string
	Status             string
}
