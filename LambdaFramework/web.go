package LambdaFramework

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/mattstools/implframework/Common"
	"github.com/mattstools/weberrors/WebErrors"
	"net/http"
)

// Helper to automatically process errors
func WebErrorResponse(error error) events.APIGatewayProxyResponse {
	Common.ProcessError(error)
	switch t := error.(type) {
	case WebErrors.WebError:
		return WebBodyResponse(t.MakeSafe(), t.StatusCode)
	default:
		return WebStatusResponse(http.StatusInternalServerError)
	}
}

// Response to simply emit a basic status code
func WebStatusResponse(statusCode int) events.APIGatewayProxyResponse {
	var toRespond struct {
		Message string
	}

	toRespond.Message = http.StatusText(statusCode)
	return WebBodyResponse(toRespond.Message, statusCode)
}

// Response to send back a body with a given status code
func WebBodyResponse(body interface{}, statusCode int) events.APIGatewayProxyResponse {
	bodyToReturn, _ := json.Marshal(body)
	fmt.Println("Returning")
	toReturn := events.APIGatewayProxyResponse{
		Headers:    getHeaders(),
		StatusCode: statusCode,
		Body:       string(bodyToReturn),
	}
	fmt.Println(toReturn)
	return toReturn
}

func getHeaders() map[string]string {
	toReturn := make(map[string]string)
	toReturn["X-Requested-With"] = "*"
	toReturn["Access-Control-Allow-Headers"] = "Content-Type,X-Amz-Date,Authorization,X-Api-Key,x-requested-with"
	toReturn["Access-Control-Allow-Origin"] = "*"
	toReturn["Access-Control-Allow-Methods"] = "POST,GET,OPTIONS"
	return toReturn
}
