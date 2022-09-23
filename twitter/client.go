package twitter

import (
	"os"

	twitter "github.com/michimani/gotwi"
)

// Credentials for a twitter application
type Credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
	BearerToken       string
}

// Read credentials from environment variables
func GetCredentials() Credentials {
	return Credentials{
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
		ConsumerKey:       os.Getenv("GOTWI_API_KEY"),
		ConsumerSecret:    os.Getenv("GOTWI_API_KEY_SECRET"),
	}
}

// Construct twitter client from credentials
func GetClient(credentials Credentials) (*twitter.GotwiClient, error) {
	in := &twitter.NewGotwiClientInput{
		AuthenticationMethod: twitter.AuthenMethodOAuth1UserContext,
		OAuthToken:           credentials.AccessToken,
		OAuthTokenSecret:     credentials.AccessTokenSecret,
	}
	c, err := twitter.NewGotwiClient(in)
	if err != nil {
		return nil, err
	}
	return c, nil
}
