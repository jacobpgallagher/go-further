package main

import (
	mw "github.com/SocialCodeInc/go-common/echo"
	"github.com/SocialCodeInc/go-common/echo/oauth1storage"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

func main() {
	db := sqlx.MustConnect("mysql", "root:@(localhost:3306)/test_db")
	dbFetcher := dbstorage.NewDBFetcher(db, "oauth_consumers")

	// START OMIT
	e := echo.New()
	e.Use(mw.MultipleAuthentication(
		mw.OAuth1(dbFetcher.GetSecretForKey),
		mw.Bouncer("https://alpha.socialcodedev.com", "SecretsInPresentationsWouldBeBad")))
	// END OMIT
}
