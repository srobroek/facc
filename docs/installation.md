# Quick Start - Install facc

> [!TIP]
> facc is installable via [instl.sh](https://instl.sh).\
> You just have to run the following command and you're ready to go!

<!-- tabs:start -->

#### ** Windows **

### Windows Command

```powershell
iwr instl.sh/srobroek/facc/windows | iex
```

#### ** Linux **

### Linux Command

```bash
curl -sSL instl.sh/srobroek/facc/linux | bash
```

#### ** macOS **

### macOS Command

```bash
curl -sSL instl.sh/srobroek/facc/macos | bash
```

#### ** Compile from source **

### Compile from source with Golang

?> **NOTICE**
To compile facc from source, you have to have [Go](https://golang.org/) installed.

Compiling facc from source has the benefit that the build command is the same on every platform.\
It is not recommended to install Go only for the installation of facc.

```command
go install github.com/srobroek/facc@latest
```

<!-- tabs:end -->
