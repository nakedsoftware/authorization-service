# Software Requirements

In order to successfully build, run, debug, and develop for Authorization Service, you will need to have the following software installed in your development environment. Not all software is required for all platforms. Please review each requirement and note the platforms where the software is required.

1. [Homebrew](#homebrew)
1. [Git](#git)
1. [GitHub CLI](#github-cli)
1. [Fast Node Manager](#fast-node-manager)
1. [Docker Desktop](#docker-desktop)
1. [Visual Studio Code](#visual-studio-code)
1. [Remote Development Extension Pack for Visual Studio Code](#remote-development-extension-pack-for-visual-studio-code)
1. [JetBrains Goland](#jetbrains-goland)

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

## Fast Node Manager

[Fast Node Manager](https://github.com/Schniz/fnm) is a version manager for [Node.js](https://nodejs.org). Fast Node Manager can be used to install and manage multiple versions of Node.js, and allow developers to switch between versions as necessary for different projects. Fast Node Manager can integrate into the developer's shell and automatically select the correct version of Node.js to use for a particular project when a `.node-version` file is present in the root directory of a project. The `.node-version` file contains the version number of Node.js that should be used with the software project.

- __Apple macOS or Linux__: Fast Node Manager can be installed using a script you download. The script will use [Homebrew](#homebrew) to install Fast Node Manager and update your shell profile. In a terminal, run:

```shell
curl -fsSL https://fnm.vercel.app/install | bash
```

- __Microsoft Windows__: Fast Node Manager can be installed using [WinGet](https://learn.microsoft.com/en-us/windows/package-manager/winget/). In a Command Prompt window, run:

```batch
winget install Schniz.fnm
```

## Docker Desktop

- __Apple macOS__
- __Microsoft Windows__

[Docker Desktop](https://www.docker.com/products/docker-desktop/) is a tool that developers can use to build and run Docker containers on their development machine. Docker Desktop hosts and runs containers on the host operating system for developers. For development of Authorization Service, most development will take place in a [development container](https://containers.dev), which is a special kind of Docker container that includes all of the tools and programming languages that are needed to build, run, debug, and develop Authentication Service.

Docker Desktop can be downloaded from the Docker [website](https://www.docker.com/products/docker-desktop/).

## Visual Studio Code

[Visual Studio Code](https://code.visualstudio.com) is an open-source text editor and extensible development environment created by [Microsoft](https://www.microsoft.com). Visual Studio Code can be used to build, run, debug, and test the Authorization Service software components. Naked Software uses Visual Studio Code to do development activities for Authorization Service.

Visual Studio Code can be downloaded and installed from the [website](https://code.visualstudio.com).

## Remote Development Extension Pack for Visual Studio Code

[Remote Development Extension Pack](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack) for [Visual Studio Code](#visual-studio-code) is a set of extensions that are helpful with developing in remote environments. Remote Development Extension Pack includes support for running [development containers](https:/containers.dev). A development container is a Docker container that has been pre-configured to contain development tools and libraries needed to build, run, test, and develop for Authorization Service.

Remote Development Extension Pack can be installed from [here](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack).

## OPTIONAL: JetBrains Goland

[JetBrains Goland](https://www.jetbrains.com/go/) is a professional development environment for software development using the [Go programming language](https://go.dev/). Goland can be used as an alternative to developing with [Visual Studio Code](#visual-studio-code) when the extra features and power of Goland is needed.

Goland can be downloaded and installed from the [Goland website](https://www.jetbrains.com/go/).
