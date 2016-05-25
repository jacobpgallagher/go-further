package main

// START OMIT
import (
	"github.com/Jeffail/gabs"
	"github.com/SocialCodeInc/go-common/loggersdk"
)

func main() {
	k := loggersdk.NewKinesisWriter(
		"alpha-coredata",
		"us-east-1",
		"no-real-key-here",
		"no-real-secret-here",
	).WithMaxAttempts(10)

	record := gabs.New()
	record.Set("http://schemas.socialcodedev.com/analytics/label_labels.1.json", "$schema")
	record.SetP("2016-05-25T12:00:00", "metadata.crawl_time")
	record.SetP("e918b138-2200-11e6-b67b-9e71128cae77", "metadata.guid")

	// Just blank data, as an example.
	record.Set(map[string]interface{}{}, "data")

	k.Write([]interface{}{record})
	//END OMIT
}
