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

package serviceProviders

import (
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/rtkwlf/bmx/saml/serviceProviders/aws"
)

type ServiceProvider interface {
	GetCredentials(saml string, role aws.AwsRole) *sts.Credentials
	ListRoles(saml string) ([]aws.AwsRole, error)
	AssumeRole(creds sts.Credentials, targetRole string, sessionName string) (*sts.Credentials, error)
}
