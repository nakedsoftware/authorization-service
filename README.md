# Authorization Service

Authorization Service provides identity management and authorization services to client applications and resource servers. Authorization Service implements the OAuth 2, OpenID Connect, and WebAuthn protocols to securely identify users and issue tokens that confirm the user or client identity or authorizes client applications to perform actions on resources on behalf of users.

Authorization Service supports managing the user base for an application or service. Authorization Service can support both new user sign up and existing user log in operations. Authorization Service provides secure support for managing _local users_ using passwords and multi-factor authentication, or integration with external trusted identity providers using the OpenID Connect or other trusted protocols. Authorization Service supports the use of secure passkeys based on private key cryptography to authenticate user identities using the WebAuthn protocol.

Authorization Service supports authorization by issuing cryptographically secure and verifiable access tokens. Resource servers can verify that the access tokens were issued by the authorization service and can be trusted. Access tokens can further be used by resource servers to implement authorization based on the user's or client's identity and permissions granted to that identity.

## Getting Started

Before beginning to work with the Authorization Service source code, please read the [software requirements](docs/software_requirements.md) and ensure that you have installed the necessary software into your development environment. After verifying that your development environment is ready, you can begin by cloning the [GitHub repository](https://github.com/nakedsoftware/authorization-service) into your development environment. In a terminal (Apple macOS or Linux) or Command Prompt window (Microsoft Windows), navigate to the location in your file system where you want to store the repository and run:

    gh repo clone nakedsoftware/authorization-service

This command will clone the repository into the `authorization-service` subdirectory and will check out the `main` branch.

After cloning the repository, there are additional dependencies such as tools and libraries that need to be downloaded and configured before you can begin working with the source code. These steps have been automated and stored in the [`setup.sh`](setup.sh) or [`Setup.bat`](Setup.bat) programs. There programs need to be run first.

On Apple macOS or Linux, in the terminal, run:

```shell
cd authorization-service
./setup.sh
```

On Microsoft Windows, in the Command Prompt window, run:

```batch
cd authorization-service
Setup.bat
```

After `setup.sh` or `Setup.bat` completes, the repository is ready for development.

## Using the Development Container

It is recommended that most development be performed using the [development container](https://containers.dev). The development container is pre-configured with all development tools and libraries that you will need to build, run, debug, and develop with the source code for Authorization Service.

To use the development container with [Visual Studio Code](https://code.visualstudio.com), open the root directory for the Authorization Service repository in Visual Studio Code. Once loaded, Visual Studio Code will prompt you to re-open within the development container.

Using a [JetBrains](https://jetbrains.com) IDE such as [Goland](https://jetbrains.com/go), you can launch the development container from within the IDE. From the Welcome screen, use the __Dev Containers__ tab to locate the local repository and start the development container.
