@sooho/cli
==========

CLI tool to interact with Sooho

[![Version](https://img.shields.io/npm/v/@sooho/cli.svg)](https://npmjs.org/package/@sooho/cli)
[![Downloads/week](https://img.shields.io/npm/dw/@sooho/cli.svg)](https://npmjs.org/package/@sooho/cli)
[![License](https://img.shields.io/npm/l/@sooho/cli.svg)](https://github.com/soohoio/sooho/blob/master/package.json)

* [Usage](#usage)
* [Commands](#commands)

# Usage

```sh-session
$ npm install -g @sooho/cli
$ sooho COMMAND
running command...
$ sooho (-v|--version|version)
@sooho/cli/0.0.2-alpha.0 darwin-x64 node-v9.11.1
$ sooho --help [COMMAND]
USAGE
  $ sooho COMMAND
...
```

# Commands

* [`sooho encrypt INPUT_PATH`](#sooho-encrypt-input-path)
* [`sooho audit INPUT_PATH`](#sooho-audit-input-path)
* [`sooho help [COMMAND]`](#sooho-help-command)

## `sooho encrypt INPUT_PATH`

Encrypt source code into hash file

```
USAGE
  $ sooho encrypt INPUT_PATH

ARGUMENTS
  INPUT_PATH  entry path

OPTIONS
  -a, --abstract  turn on abstraction mode
  -s, --save      save encrypted file
  -h, --help      show CLI help

EXAMPLE
  $ sooho encrypt INPUT_PATH
```

_See code: [src/commands/encrypt.ts](https://github.com/soohoio/sooho/blob/master/packages/sooho-cli/src/commands/encrypt.ts)

## `sooho audit INPUT_PATH`

Audit smart contract

```
USAGE
  $ sooho audit INPUT_PATH

ARGUMENTS
  INPUT_PATH  entry path

EXAMPLE
  $ sooho audit INPUT_PATH
```

_See code: [src/commands/audit.ts](https://github.com/soohoio/sooho/blob/master/packages/sooho-cli/src/commands/audit.ts)

## `sooho help [COMMAND]`

display help for sooho

```
USAGE
  $ sooho help [COMMAND]

ARGUMENTS
  COMMAND  command to show help for

OPTIONS
  --all  see all commands in CLI
```

_See code: [@oclif/plugin-help](https://github.com/oclif/plugin-help/blob/v2.1.3/src/commands/help.ts)
