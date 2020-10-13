/*
Copyright 2019 D2L Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package bmx

import (
	"encoding/json"
	"log"
	"time"

	"github.com/jrbeverly/bmx/saml/identityProviders"

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

func CredentialProcess(idProvider identityProviders.IdentityProvider, printOptions CredentialProcessCmdOptions) string {
	printOptions.User = getUserIfEmpty(printOptions.User)
	user := GetUserInfoFromCredentialProcessCmdOptions(printOptions)

	saml, err := authenticate(user, idProvider)
	if err != nil {
		log.Fatal(err)
	}

	creds := AwsServiceProvider.GetCredentials(saml, printOptions.Role)
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
