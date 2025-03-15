# Software Requirements

In order to successfully build, run, debug, and develop for Authorization Service, you will need to have the following software installed in your development environment. Not all software is required for all platforms. Please review each requirement and note the platforms where the software is required.

1. [Homebrew]()
1. [Git]()
1. [GitHub CLI]()

## Homebrew

- __Apple macOS__
- __Linux__

[Homebrew](https://brew.sh) is a package management tool for Apple macOS and Linux operating systems. Homebrew can be used to install many popular open source or commercial software packages, libraries, and other tools. Naked Software uses Homebrew to install and keep up to date common software development tools used in our software development activities.

To install Homebrew, open a terminal and run:

    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

After installing Homebrew, you will need to close and restart your terminal before environment changes take effect.

## Git

[Git](https://git-scm.com) is a popular distributed version control management system. Using Git, each developer maintains and works out of their own copy, or _clone_, of the Authorization Service repository and can work in a stable environment either offline or online. Git makes it easy to share changes with other developers either directly or through a shared repository such as [GitHub](https://github.com). Git also supports a powerful branching model that allows developers to work on new features in stable isolation and supports a code review process through _pull requests_ using a service such as GitHub.

- __Apple macOS or Linux__: Git can be installed using [Homebrew](#homebrew). In a terminal, run:

```shell
brew install git
```

- __Microsoft Windows__: Git can be installed using [WinGet](https://learn.microsoft.com/en-us/windows/package-manager/winget/). In a Command Prompt window, run:

```batch
winget install --id Git.Git -e --source winget
```

## GitHub CLI

[GitHub CLI](https://cli.github.com) is a command line tool for cloning Git repositories hosted on [GitHub](https://github.com), performing one-off tasks, or automating common administrative actions on GitHub repositories and projects. Naked Software uses GitHub CLI to clone our GitHub-hosted repositories, automate administrative tasks, and manually start GitHub Actions workflows.

- __Apple macOS or Linux__: GitHub CLI can be installed using [Homebrew](#homebrew). In a terminal, run:

```shell
brew install gh
```

- __Microsoft Windows__: GitHub CLI can be installed using [WinGet](https://learn.microsoft.com/en-us/windows/package-manager/winget/). In a Command Prompt window, run:

```batch
winget install --id GitHub.cli
```
