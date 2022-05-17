package ctrl

import "github.com/go-pg/pg/v10"

type Ctrl struct {
	db *pg.DB
}

func New(db *pg.DB) *Ctrl {
	c := &Ctrl{
		db: db,
	}

	return c
}
