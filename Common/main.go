package Common

import (
	"fmt"
	"github.com/getsentry/raven-go"
	"github.com/mattstools/weberrors/WebErrors"
	"os"
)

func ProcessError(err error) {
	dsn := os.Getenv("SENTRY_DSN")
	fmt.Println(err)
	if dsn == "" {
		fmt.Println("Missing SENTRY_DSN")
	} else {
		switch t := err.(type) {
		case WebErrors.WebError:
			if t.StatusCode != 500 && t.StatusCode != 503 {
				// we do not want non-500 error codes sent back
				return
			}
		}

		_ = raven.SetDSN(dsn)
		raven.CaptureErrorAndWait(err, nil)
	}
}
