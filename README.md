# GitGud CLI

[![GitGud](https://img.shields.io/badge/GitGud-v1.0-red?style=flat-square)](https://github.com/HRKings/GitGud/tree/stable)

This repository is cross-platform CLI (Command Line Interface) for the [GitGud](https://github.com/HRKings/GitGud/tree/stable) modular git model. It contains a series of commands that help you use git more quickly and following the model.

- [GitGud CLI](#gitgud-cli)
  - [Installation](#installation)
  - [Usage](#usage)
    - [Commit Module](#commit-module)
      - [Flags](#flags)
      - [Full commit](#full-commit)
  - [Compiling](#compiling)
  - [Contributing](#contributing)

## Installation

Just download and drop the latest release into a folder, add it to your path and call the executable in your terminal of preference.

## Usage

Using it is rather simple, after you have it on your path you can call it using: `gitgud <module> <subcommand>`

The commands are divided into modules, just like the original model, the ones available at the moment are:

### Commit Module

This module is equivalent to the [Commit submodel](https://github.com/HRKings/GitGud/blob/stable/Git/Commit.md)

This command is a wrapper around the `git commit -m`, which means that you can use the flags `-a` and `--amend`

```Bash
gitgud commit -m "Commit subject"
gitgud c -m "Commit subject"
```

#### Flags

The following flags can be used:

- `-a` or `--all` : Equivalent to the same flag in git
- `--amend` : Equivalent to the same flag in git
- `-m` : Equivalent to the same flag in git
- `-d` or `--domain` : Add a domain to the commit (an impacted area of the code)
- `-q` or `--quick` : Don't ask for missing parts of the commit if none is provided (Only the domain in that case)

#### Full commit

The full commit subcommand is a variation of the base command that will generate a complete commit message (with body and footer)

```Bash
gitgud commit -m "Commit subject" full
gitgud commit -m "Commit subject" f
```

The following flags can be used after the subcommand:

- `-b` or `--body` :  Specify the text of the body of the commit
- `-c` or `--closes` : Specify a list of closed issues by the commit
- `-s` or `--see` : Specify a list of referenced issues by the commit

## Compiling

You are welcome to clone and compile this repository. For this you will need Go on the latest version and compile it from the terminal using:

```Bash
go -o gitgud
```

## Contributing

You are more than welcome to contribute to this repository opening issues and pull requests, just remember to follow the specifications of the GitGud model, as this repository follows all of it (obviously).