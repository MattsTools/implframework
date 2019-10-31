package LambdaFramework

import (
	"errors"
	"flowkey.io/packages/go/implframework/Common"
	"flowkey.io/packages/go/weberrors/WebErrors"
	"github.com/aws/aws-lambda-go/events"
)

func AuthoriserUnauthorisedAccessResponse() (events.APIGatewayCustomAuthorizerResponse, error) {
	return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Unauthorized")
}

func AuthoriserErrorResponse(err error, principalId string, resource string) (events.APIGatewayCustomAuthorizerResponse, error) {
	Common.ProcessError(err)

	switch t := err.(type) {
	case WebErrors.WebError:
		if t.StatusCode == 403 {
			policy := AuthoriserPolicyResponse(principalId, "deny",
				resource, "", 0, 0, "", "")
			return policy, nil
		}

		if t.StatusCode == 401 {
			return AuthoriserUnauthorisedAccessResponse()
		}

		return events.APIGatewayCustomAuthorizerResponse{}, t
	default:
		return AuthoriserUnauthorisedAccessResponse()
	}
}

func AuthoriserPolicyResponse(principalId, effect, resource string,
	userID string, expirationTime int64, issuedAt int64,
	jwtID string, subject string) events.APIGatewayCustomAuthorizerResponse {
	authResponse := events.APIGatewayCustomAuthorizerResponse{PrincipalID: principalId}

	if effect != "" && resource != "" {
		authResponse.PolicyDocument = events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action:   []string{"execute-api:Invoke"},
					Effect:   effect,
					Resource: []string{resource},
				},
			},
		}
	}

	// Optional output with custom properties of the String, Number or Boolean type.
	authResponse.Context = map[string]interface{}{
		"UserID":         userID,
		"ExpirationTime": expirationTime,
		"IssuedAt":       issuedAt,
		"JWTID":          jwtID,
		"Subject":        subject,
	}
	return authResponse
}
