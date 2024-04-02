package entities

import "time"

type Product struct {
	Product_ID   string
	Product_Name string
	Price        float64
	Quantity     int
	Created_At   time.Time
	Updated_At   time.Time
}
