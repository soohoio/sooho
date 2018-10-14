@sooho/cli
==========

CLI tool to interact with Sooho

[![oclif](https://img.shields.io/badge/cli-oclif-brightgreen.svg)](https://oclif.io)
[![Version](https://img.shields.io/npm/v/@sooho/cli.svg)](https://npmjs.org/package/@sooho/cli)
[![Downloads/week](https://img.shields.io/npm/dw/@sooho/cli.svg)](https://npmjs.org/package/@sooho/cli)
[![License](https://img.shields.io/npm/l/@sooho/cli.svg)](https://github.com/soohoio/sooho/blob/master/package.json)

<!-- toc -->
* [Usage](#usage)
* [Commands](#commands)
<!-- tocstop -->
# Usage
<!-- usage -->
```sh-session
$ npm install -g @sooho/cli
$ sooho COMMAND
running command...
$ sooho (-v|--version|version)
@sooho/cli/0.0.1 darwin-x64 node-v9.11.1
$ sooho --help [COMMAND]
USAGE
  $ sooho COMMAND
...
```
<!-- usagestop -->
# Commands
<!-- commands -->
* [`sooho encrypt [FILEPATH]`](#sooho-encrypt-filepath)
* [`sooho help [COMMAND]`](#sooho-help-command)

## `sooho encrypt FILEPATH`

Encrypt source code into hash file

```
USAGE
  $ sooho encrypt FILEPATH

ARGUMENTS
  FILEPATH  entry path

OPTIONS
  -f, --force
  -h, --help   show CLI help

EXAMPLE
  $ sooho encrypt FILE_PATH
```

_See code: [src/commands/encrypt.ts](https://github.com/soohoio/sooho/blob/v0.0.1/src/commands/encrypt.ts)_

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

_See code: [@oclif/plugin-help](https://github.com/oclif/plugin-help/blob/v2.1.3/src/commands/help.ts)_
<!-- commandsstop -->
