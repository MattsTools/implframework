package LambdaFramework

import "github.com/mattstools/implframework/Common"

func KinesisErrorResponse(err error, alert bool) error {
	Common.ProcessError(err)

	if alert {
		return err
	}

	return nil
}
