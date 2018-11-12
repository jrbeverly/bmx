[![Build Status](https://travis-ci.com/Brightspace/bmx.svg?token=XBuHJueJZM92zaxjesN6&branch=go)](https://travis-ci.com/Brightspace/bmx)

# BMX

BMX grants you API access to your AWS accounts, based on Okta credentials that you already own.  It uses your Okta identity to create short-term AWS STS tokens, as an alternative to long-term IAM access keys.  BMX manages your STS tokens with five commands:

1. `bmx aws` wraps the AWS CLI, calling the CLI on your behalf and updating tokens as necessary.  Example: `bmx aws cloudformation describe-stacks`.
1. `bmx print` writes your short-term tokens to `stdout` as AWS environment variables.  You can execute `bmx print`'s output to make the environment variables available to your shell.
1. `bmx write` writes your short-term tokens to `~/.aws/credentials`.
1. `bmx renew` requests a new token with a one-hour TTL.
1. `bmx remove` forgets a token.

BMX prints detailed usage information when you run `bmx -h` or `bmx <cmd> -h`.

[A BMX demo](https://internal.desire2learncapture.com/1/Watch/6371.aspx) is on Capture.

## Installation

BMX installs with Pip from D2L's Artifactory repository.

### Credentials

You will be asked for Artifactory credentials while installing BMX.  Your credentials are your D2L network credentials.  For example, your username is \<first initial> + \<last name>.  The next section shows the Artifactory password prompt in Bash.

### Bash

```
vagrant@ubuntu-artful:~$ python3 --version
Python 3.6.3

vagrant@ubuntu-artful:~$ pip3 install --user --upgrade --extra-index-url https://d2lartifacts.artifactoryonline.com/d2lartifacts/api/pypi/pypi-local/simple bmx
Collecting bmx
User for d2lartifacts.artifactoryonline.com: credekop
Password:
```

### PowerShell

```
PS C:\Users\credekop> py -3 --version
Python 3.6.2

PS C:\Users\credekop> py -3 -m pip install --user --upgrade --extra-index-url https://d2lartifacts.artifactoryonline.com/d2lartifacts/api/pypi/pypi-local/simple bmx
```

## System Requirements

* [Python 3.6+](https://www.python.org/downloads/)
* Pip, the Python installer.

## Features
1. BMX is multi-platform: it runs on Linux and Windows.
1. BMX maintains your Okta session for 12 hours: you enter your Okta password once a day, and BMX takes care of the rest.
1. BMX supports TOTP and SMS MFA.
1. BMX manages its own AWS STS tokens and never modifies `~/.aws/credentials` without explicit direction from you.  (See `bmx write`.)

## Development

BMX is designed to be extensible and easily rolled out.

* It's a command-driven utility (think of Git, Terraform, or the AWS CLI) where new commands can be added to the base system.
* It's on our private Artifactory repo and can be easily installed.

BMX is written in Python, like the AWS CLI.

* It introduces no new language dependencies.
* `bmx aws` runs in the same process as the AWS CLI, reducing overhead.

### Developer Setup

```bash
git clone git@github.com:Brightspace/bmx.git
cd bmx
pip install -e .
bmx -h
```

### Pylint

BMX uses [Pylint](https://www.pylint.org/) to enforce styling and run quality checkers.

**Install**: `pip install pylint`

**Lint**: `pylint bmx`

### Slated development

BMX has [issues](https://github.com/Brightspace/bmx/issues).

## Usage Examples

### Getting Help

```bash
$ python3 --version
Python 3.6.2

$ bmx -h

usage: bmx {aws,write,print} ...

Okta time-out helper for AWS CLI

commands:
  {aws,write,print}
    aws                awscli with automatic STS token renewal
    write              create new AWS credentials and write them to ~/.aws/credentials
    print              create new AWS credentials and print them to stdout

Copyright 2017 D2L Corporation
```

## `bmx aws` in Bash
```
$ bmx aws cloudformation describe-stacks
{
    "Stacks": [
        {
...
        }
    ]
}

$ bmx aws cloudformation describe-stacks
Your AWS STS token has expired.  Renewing...
Okta username: credekop
Okta password:

Available AWS Accounts:
 1: DEV-BroadcastEventService
 2: Dev-AnalyticsInegration
 3: Dev-BDP
 4: Dev-CI
 5: Dev-IPA-EDU
 6: Dev-LMS
 7: Dev-PD-Tools
 8: Dev-ServiceDashboard
 9: Dev-Staging
10: Dev-Translation
11: Lrn-NimbusToronto
12: PRD-BroadcastEventService
13: Prd-BDP
14: Prd-CDN
15: Prd-NA
16: Prd-ServiceDashboard
17: Prd-Totem
18: Service Dashboard
AWS Account Index: 11

Available Roles in Lrn-NimbusToronto:
 1: Lrn-NimbusToronto-Owner
 2: Lrn-NimbusToronto-User
Role Index: 2

{
    "Stacks": [
        {
...
        }
    ]
}
```

## `bmx` via Docker

You can build and run bmx inside of Docker:

1. Clone this repo
1. Run `docker build . -t bmx` in the cloned directory

Optionally, add an alias/helper in your `.bashrc`, for example:

```bash
bmx() {
  docker run \
    --rm -ti \
    -e "BMX_USERNAME=${BMX_USERNAME-$USER}" \
    -e "BMX_ACCOUNT=${BMX_ACCOUNT-$AWS_PROFILE}" \
    -e "BMX_PROFILE=${BMX_PROFILE-$AWS_PROFILE}" \
    -e "BMX_ROLE=${BMX_ROLE-}" \
    -v "${HOME}/.aws:/home/bmx/.aws" \
    -v "${HOME}/.bmx:/home/bmx/.bmx" \
    bmx $@
}
```