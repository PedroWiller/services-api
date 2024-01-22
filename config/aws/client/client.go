package awsSession

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	"face-id-detection/config/env"
)

var Sess *session.Session

var AWSSession *session.Session

func Init() {
	AWSSession = session.Must(
		session.NewSessionWithOptions(
			session.Options{Config: aws.Config{Region: aws.String(env.AWSRegion)}}))
}
