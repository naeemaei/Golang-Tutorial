package db

import "http-mini-project/models"

// set up a database dummy
var (
	MovieDb = make(map[string]models.Movie)
)
