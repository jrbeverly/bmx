#!/usr/bin/python3

import base64
import configparser
import getpass
import os
import re
import sys
import requests

import boto3
import lxml
import oktautil
import prompt

from okta.framework.OktaError import OktaError
from requests import HTTPError

def renew_credentials():
    write_credentials(get_credentials())

def get_credentials():
    auth_client = oktautil.create_auth_client()
    sessions_client = oktautil.create_sessions_client()

    while True:
        try:
            username = prompt.prompt_for_value(input, 'Okta username: ')
            password = prompt.prompt_for_value(getpass.getpass, 'Okta password: ')

            authentication = auth_client.authenticate(username, password)
            session = sessions_client.create_session(username, password)

            applink = get_app_selection(
                filter_applinks(
                    oktautil.create_users_client(session.id).
                        get_user_applinks(session.userId)
                )
            )

            saml_assertion = oktautil.connect_to_app(
                applink.linkUrl, 
                authentication.sessionToken
            )

            role = get_role_selection(
                applink.label,
                get_app_roles(base64.b64decode(saml_assertion))
            )
            split_role = role.split(',')

            credentials = sts_assume_role(
                saml_assertion,
                split_role[0],
                split_role[1]
            )

            break
        except OktaError as e:
            print(e)
        except HTTPError as e:
            print(e)

    return credentials

def filter_applinks(applinks):
    return sorted(
        filter(
            lambda x: x.appName == "amazon_aws",
            applinks
        ),
        key=lambda x: x.label
    )

def get_app_selection(applinks):
    return applinks[
        prompt.MinMenu(
            '\nAvailable AWS Accounts: ',
            list(map(lambda x: '{}'.format(x.label), applinks)),
            'AWS Account Index: '
        ).get_selection()
    ]

def get_role_selection(app_name, roles):
    return roles[
        prompt.MinMenu(
            '\nAvailable Roles in {}:'.format(app_name),
            list(map(lambda x: re.sub('.*role/', '', x.split(',')[1]), roles)),
            'Role Index: '
        ).get_selection()
    ]

def get_app_roles(saml_assertion):
    return lxml.etree.fromstring(saml_assertion).xpath(
        '//x:AttributeStatement/x:Attribute[@Name="https://aws.amazon.com/SAML/Attributes/Role"]/x:AttributeValue/text()',
        namespaces={'x': 'urn:oasis:names:tc:SAML:2.0:assertion'}
    )

def sts_assume_role(saml_assertion, principal, role):
    response = boto3.client('sts').assume_role_with_saml(
        PrincipalArn = principal,
        RoleArn = role,
        SAMLAssertion = saml_assertion,
        DurationSeconds = 3600
    )

    return response['Credentials']

def write_credentials(credentials):
    config = configparser.ConfigParser()
    filename = os.path.expanduser('~/.aws/credentials');

    config.read(filename)
    config['default'] = {
        'aws_access_key_id': credentials['AccessKeyId'],
        'aws_secret_access_key': credentials['SecretAccessKey'],
        'aws_session_token': credentials['SessionToken']
    }

    with open(os.path.expanduser('~/.aws/credentials'), 'w') as config_file:
        config.write(config_file)

def main():
    renew_credentials()

    return 0

if __name__ == "__main__":
    sys.exit(main())
