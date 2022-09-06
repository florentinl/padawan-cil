# Padawan Cli

Padawan Cli is a command line interface for the Padawan PaaS. It allows you to login using OAuth2, create and manage a single container and access it via SSH.

## Installation

Download a release binary from the [releases page](https://gitlab.viarezo.fr/flow/padawancli/-/releases) and put it in your PATH.

## Usage

The login flow is not very friendly for now and requires you to copy and paste a token from the browser. This will be improved in the future (probably not though).

```bash
padawan login
```

This will prompt you a link to open in your browser. Once you are logged in, you will be redirected to a page with a cookie ('\_forward_auth') set in your browser. You need to copy its value and paste it in the terminal.

```bash
padawan login <cookie value>
```

You can then manage your containers with the `ctr` subcommand.
As an admin you can also manage images with the `img` subcommand.
