package main

import (
	mw "github.com/SocialCodeInc/go-common/echo"
	"github.com/SocialCodeInc/go-common/echo/oauth1storage"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

func main() {
	// START OMIT
	db := sqlx.MustConnect("mysql", "root:@(localhost:3306)/test_db")
	dbFetcher := dbstorage.NewDBFetcher(db, "oauth_consumers")

	e := echo.New()
	e.Use(echo.WrapMiddleware(mw.OAuth1(dbFetcher.GetSecretForKey)))
	// END OMIT
}
