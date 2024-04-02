package entities

import "time"

type User struct {
	User_ID    string    `db:"id, omitempty"`
	Name       string    `db:"name, omitempty"`
	Email      string    `db:"email, omitempty"`
	Password   string    `db:"password, omitempty"`
	Created_At time.Time `db:"created_at, omitempty"`
	Updated_At time.Time `db:"updated_at, omitempty"`
	Token      string
}
