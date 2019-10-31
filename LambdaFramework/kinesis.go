package LambdaFramework

import "flowkey.io/packages/go/implframework/Common"

func KinesisErrorResponse(err error, alert bool) error {
	Common.ProcessError(err)

	if alert {
		return err
	}

	return nil
}
