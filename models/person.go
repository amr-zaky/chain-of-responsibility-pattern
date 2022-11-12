package models

type Person struct {
	Age     int
	Gender  string
	Country string
	Visa    Visa
}

type Visa struct {
	Bank         string
	Balance      int
	CreditRating int
}
