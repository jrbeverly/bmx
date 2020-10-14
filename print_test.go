package bmx_test

import (
	"strings"
	"testing"

	"github.com/jrbeverly/bmx"
	"github.com/jrbeverly/bmx/mocks"
	awsmocks "github.com/jrbeverly/bmx/saml/serviceProviders/aws/mocks"
)

func assertAwsTokenEnv(t *testing.T, output string) {
	awsStsVars := [3]string{
		"AWS_ACCESS_KEY_ID",
		"AWS_SECRET_ACCESS_KEY",
		"AWS_SESSION_TOKEN",
	}

	for _, envVar := range awsStsVars {
		if !strings.Contains(output, envVar) {
			t.Errorf("Environment variable %s missing, got: %s", envVar, output)
		}
	}
}

func TestMonkey(t *testing.T) {
	options := bmx.PrintCmdOptions{
		Org: "myorg",
	}

	oktaClient := &mocks.Mokta{}

	consolerw := mocks.ConsoleReaderMock{}
	awsProvider := awsmocks.AwsServiceProviderMock{}

	output := bmx.Print(oktaClient, awsProvider, consolerw, options)

	assertAwsTokenEnv(t, output)
}

func TestPShellPrint(t *testing.T) {
	options := bmx.PrintCmdOptions{
		Org:    "myorg",
		Output: bmx.Powershell,
	}

	oktaClient := &mocks.Mokta{}

	consolerw := mocks.ConsoleReaderMock{}
	awsProvider := awsmocks.AwsServiceProviderMock{}

	output := bmx.Print(oktaClient, awsProvider, consolerw, options)

	assertAwsTokenEnv(t, output)
	if !strings.Contains(output, "$env:") {
		t.Errorf("Shell command was incorrect, got: %s, expected powershell", output)
	}
}

func TestBashPrint(t *testing.T) {
	options := bmx.PrintCmdOptions{
		Org:    "myorg",
		Output: bmx.Bash,
	}

	oktaClient := &mocks.Mokta{}

	consolerw := mocks.ConsoleReaderMock{}
	awsProvider := awsmocks.AwsServiceProviderMock{}

	output := bmx.Print(oktaClient, awsProvider, consolerw, options)

	assertAwsTokenEnv(t, output)
	if !strings.Contains(output, "export ") {
		t.Errorf("Shell command was incorrect, got: %s, expected bash", output)
	}
}
