package aws_test

import (
	"testing"
	"time"

	"github.com/jrbeverly/bmx/mocks"

	"github.com/aws/aws-sdk-go/service/sts/stsiface"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/service/sts"
	awsService "github.com/jrbeverly/bmx/saml/serviceProviders/aws"
)

type stsMock struct {
	stsiface.STSAPI
}

func (s *stsMock) AssumeRoleWithSAML(input *sts.AssumeRoleWithSAMLInput) (*sts.AssumeRoleWithSAMLOutput, error) {
	out := &sts.AssumeRoleWithSAMLOutput{
		Credentials: &sts.Credentials{
			AccessKeyId:     aws.String("access_key_id"),
			SecretAccessKey: aws.String("secrest_access_key"),
			SessionToken:    aws.String("session_token"),
			Expiration:      aws.Time(time.Now().Add(time.Hour * 8)),
		},
	}

	return out, nil
}

func TestMonkey(t *testing.T) {
	consolerw := mocks.ConsoleReaderMock{}
	provider := awsService.NewAwsServiceProvider(consolerw)
	provider.StsClient = &stsMock{}
	provider.InputReader = mocks.ConsoleReaderMock{}

	// This is a base64 encoded minimal SAML input
	saml := "PHNhbWxwOlJlc3BvbnNlPgogIDxzYW1sOkFzc2VydGlvbj4KICAgIDxzYW1sOkF0dHJpYnV0ZVN0YXRlbWVudD4KICAgICAgPHNhbWw6QXR0cmlidXRlIE5hbWU9Imh0dHBzOi8vYXdzLmFtYXpvbi5jb20vU0FNTC9BdHRyaWJ1dGVzL1JvbGUiPgogICAgICAgIDxzYW1sOkF0dHJpYnV0ZVZhbHVlIHhzaTp0eXBlPSJ4czpzdHJpbmciPkFybixyb2xlL1JvbGVBcm48L3NhbWw6QXR0cmlidXRlVmFsdWU+CiAgICAgIDwvc2FtbDpBdHRyaWJ1dGU+CiAgICA8L3NhbWw6QXR0cmlidXRlU3RhdGVtZW50PgogIDwvc2FtbDpBc3NlcnRpb24+Cjwvc2FtbHA6UmVzcG9uc2U+"

	creds := provider.GetCredentials(saml, "RoleArn")
	if creds == nil {
		panic("fail")
	}
}
