package bmx

import (
	"encoding/json"
	"log"
	"time"

	"github.com/jrbeverly/bmx/console"
	"github.com/jrbeverly/bmx/saml/identityProviders"
	"github.com/jrbeverly/bmx/saml/serviceProviders/aws"

	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/jrbeverly/bmx/saml/serviceProviders"
)

type CredentialProcessCmdOptions struct {
	Org      string
	User     string
	Account  string
	NoMask   bool
	Password string
	Role     string
	Output   string
}

type CredentialProcessResult struct {
	Version         int
	AccessKeyId     string
	SecretAccessKey string
	SessionToken    string
	Expiration      time.Time
}

func GetUserInfoFromCredentialProcessCmdOptions(printOptions CredentialProcessCmdOptions) serviceProviders.UserInfo {
	user := serviceProviders.UserInfo{
		Org:      printOptions.Org,
		User:     printOptions.User,
		Account:  printOptions.Account,
		NoMask:   printOptions.NoMask,
		Password: printOptions.Password,
		Role:     printOptions.Role,
	}
	return user
}

func CredentialProcess(idProvider identityProviders.IdentityProvider, consolerw console.ConsoleReader, printOptions CredentialProcessCmdOptions) string {
	printOptions.User = getUserIfEmpty(consolerw, printOptions.User)
	user := GetUserInfoFromCredentialProcessCmdOptions(printOptions)

	saml, err := authenticate(user, idProvider, consolerw)
	if err != nil {
		log.Fatal(err)
	}

	aws := aws.NewAwsServiceProvider(consolerw)
	creds := aws.GetCredentials(saml, printOptions.Role)
	command := credentialProcessCommand(printOptions, creds)
	return command
}

func credentialProcessCommand(printOptions CredentialProcessCmdOptions, creds *sts.Credentials) string {
	result := &CredentialProcessResult{
		Version:         1,
		AccessKeyId:     *creds.AccessKeyId,
		SecretAccessKey: *creds.SecretAccessKey,
		SessionToken:    *creds.SessionToken,
		Expiration:      *creds.Expiration,
	}
	b, err := json.Marshal(result)
	if err != nil {
		return ""
	}
	return string(b)
}
