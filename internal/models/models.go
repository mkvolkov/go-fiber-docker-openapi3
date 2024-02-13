package models

import "github.com/guregu/null"

type FEmployee struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Gender  string `json:"gender"`
	Age     int    `json:"age"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Vdays   int    `json:"vdays"`
}

type PEmployee struct {
	ID      int      `db:"id"`
	Name    string   `db:"name"`
	Phone   string   `db:"phone"`
	Gender  string   `db:"gender"`
	Age     int      `db:"age"`
	Email   string   `db:"email"`
	Address string   `db:"address"`
	Vdays   null.Int `db:"vdays"`
}

type Vdays struct {
	Vdays int `json:"vdays"`
}
