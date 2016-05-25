package main

import (
	mw "github.com/SocialCodeInc/go-common/echo"
	"github.com/labstack/echo"
)

func main() {
	// START OMIT
	e := echo.New()
	e.Use(echo.WrapMiddleware(mw.Bouncer("https://alpha.socialcodedev.com",
		"SecretsInPresentationsWouldBeBad")))
	// END OMIT
}
