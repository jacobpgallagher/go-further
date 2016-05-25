package main

// START OMIT
import (
	commonMW "github.com/SocialCodeInc/go-common/echo"
	"github.com/SocialCodeInc/go-common/log"
	"github.com/labstack/echo"
)

func main() {
	err := log.InitializeSentry("http://c1ec9f2d2e6c560692709158d48260cd:e3678739d2994a24b9f517a12a1372b8@sentry.socialcodedev.com:9000/85")
	if err != nil {
		log.Errorf("Error iniaitlizing Sentry: %s", err.Error())
	}
	e := echo.New()
	e.Use(commonMW.Logger())
}

//END OMIT
