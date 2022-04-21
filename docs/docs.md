# facc

## Usage
> Facc your firewall

facc
## Examples

```bash
facc 
facc

```

## Flags
|Flag|Usage|
|----|-----|
|`--debug`|enable debug messages|
|`--disable-update-checks`|disables update checks|
|`--raw`|print unstyled raw output (set it if output is written to a file)|

## Commands
|Command|Usage|
|-------|-----|
|`facc completion`|Generate the autocompletion script for the specified shell|
|`facc help`|Help about any command|
# ... completion
`facc completion`

## Usage
> Generate the autocompletion script for the specified shell

facc completion

## Description

```
Generate the autocompletion script for facc for the specified shell.
See each sub-command's help for details on how to use the generated script.

```

## Commands
|Command|Usage|
|-------|-----|
|`facc completion bash`|Generate the autocompletion script for bash|
|`facc completion fish`|Generate the autocompletion script for fish|
|`facc completion powershell`|Generate the autocompletion script for powershell|
|`facc completion zsh`|Generate the autocompletion script for zsh|
# ... completion bash
`facc completion bash`

## Usage
> Generate the autocompletion script for bash

facc completion bash

## Description

```
Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(facc completion bash)

To load completions for every new session, execute once:

#### Linux:

	facc completion bash > /etc/bash_completion.d/facc

#### macOS:

	facc completion bash > /usr/local/etc/bash_completion.d/facc

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion fish
`facc completion fish`

## Usage
> Generate the autocompletion script for fish

facc completion fish

## Description

```
Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

	facc completion fish | source

To load completions for every new session, execute once:

	facc completion fish > ~/.config/fish/completions/facc.fish

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion powershell
`facc completion powershell`

## Usage
> Generate the autocompletion script for powershell

facc completion powershell

## Description

```
Generate the autocompletion script for powershell.

To load completions in your current shell session:

	facc completion powershell | Out-String | Invoke-Expression

To load completions for every new session, add the output of the above command
to your powershell profile.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion zsh
`facc completion zsh`

## Usage
> Generate the autocompletion script for zsh

facc completion zsh

## Description

```
Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions for every new session, execute once:

#### Linux:

	facc completion zsh > "${fpath[1]}/_facc"

#### macOS:

	facc completion zsh > /usr/local/share/zsh/site-functions/_facc

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... help
`facc help`

## Usage
> Help about any command

facc help [command]

## Description

```
Help provides help for any command in the application.
Simply type facc help [path to command] for full details.
```


---
> **Documentation automatically generated with [PTerm](https://github.com/pterm/cli-template) on 21 April 2022**
