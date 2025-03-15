# Authorization Service

Authorization Service provides identity management and authorization services to client applications and resource servers. Authorization Service implements the OAuth 2, OpenID Connect, and WebAuthn protocols to securely identify users and issue tokens that confirm the user or client identity or authorizes client applications to perform actions on resources on behalf of users.

Authorization Service supports managing the user base for an application or service. Authorization Service can support both new user sign up and existing user log in operations. Authorization Service provides secure support for managing _local users_ using passwords and multi-factor authentication, or integration with external trusted identity providers using the OpenID Connect or other trusted protocols. Authorization Service supports the use of secure passkeys based on private key cryptography to authenticate user identities using the WebAuthn protocol.

Authorization Service supports authorization by issuing cryptographically secure and verifiable access tokens. Resource servers can verify that the access tokens were issued by the authorization service and can be trusted. Access tokens can further be used by resource servers to implement authorization based on the user's or client's identity and permissions granted to that identity.

## Getting Started

Before beginning to work with the Authorization Service source code, please read the [software requirements](docs/software_requirements.md) and ensure that you have installed the necessary software into your development environment. After verifying that your development environment is ready, you can begin by cloning the [GitHub repository](https://github.com/nakedsoftware/authorization-service) into your development environment. In a terminal (Apple macOS or Linux) or Command Prompt window (Microsoft Windows), navigate to the location in your file system where you want to store the repository and run:

    gh repo clone nakedsoftware/authorization-service

This command will clone the repository into the `authorization-service` subdirectory and will check out the `main` branch.
