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

package main

import (
	"fmt"
	"log"

	"github.com/jrbeverly/bmx/config"

	"github.com/jrbeverly/bmx/saml/identityProviders/okta"

	"github.com/jrbeverly/bmx"
	"github.com/spf13/cobra"
)

var processOptions = bmx.CredentialProcessCmdOptions{}

func init() {
	processCmd.Flags().StringVar(&processOptions.Org, "org", "", "the okta org api to target")
	processCmd.Flags().StringVar(&processOptions.User, "user", "", "the user to authenticate with")
	processCmd.Flags().StringVar(&processOptions.Account, "account", "", "the account name to auth against")
	processCmd.Flags().StringVar(&processOptions.Role, "role", "", "the desired role to assume")
	processCmd.Flags().BoolVar(&processOptions.NoMask, "nomask", false, "set to not mask the password. this helps with debugging.")
	processCmd.Flags().StringVar(&processOptions.Output, "output", "", "the output format [bash|powershell]")

	if userConfig.Org == "" {
		processCmd.MarkFlagRequired("org")
	}

	rootCmd.AddCommand(processCmd)
}

var processCmd = &cobra.Command{
	Use:   "credential-process",
	Short: "Credentials to awscli",
	Long:  `Supply the credentials in compatible format`,
	Run: func(cmd *cobra.Command, args []string) {
		mergedOptions := mergeProcessOptions(userConfig, processOptions)

		oktaClient, err := okta.NewOktaClient(mergedOptions.Org)
		if err != nil {
			log.Fatal(err)
		}

		command := bmx.CredentialProcess(oktaClient, mergedOptions)
		fmt.Println(command)
	},
}

func mergeProcessOptions(uc config.UserConfig, pc bmx.CredentialProcessCmdOptions) bmx.CredentialProcessCmdOptions {
	if pc.Org == "" {
		pc.Org = uc.Org
	}
	if pc.User == "" {
		pc.User = uc.User
	}
	if pc.Account == "" {
		pc.Account = uc.Account
	}
	if pc.Role == "" {
		pc.Role = uc.Role
	}

	return pc
}
