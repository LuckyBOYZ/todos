package repository

import "database/sql"

type Todo struct {
	Id          int
	Description string
	Done        bool
	Created     sql.NullTime
}
