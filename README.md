# wercker api client

go-wercker-api is the official wercker client for the
[wercker API][api-docs].

> Caution: Both the API and this client are under active development. This
> client may introduce breaking changes, so be sure to vendor in this client.

# Usage

To start using this client, import this package and create a new client:

```golang
import "github.com/wercker/go-wercker-api"

client := wercker.NewClient(nil)
```

This will create a new client with the default settings. If you want to override
the default behavious, then you need to create a `wercker.Options` object and
pass this `wercker.NewClient`:

```golang
import "github.com/wercker/go-wercker-api"

options := &wercker.Options{}
client := wercker.NewClient(options)
```

## Authentication

The client uses a `Provider` to get the credentials for the user. By default it
uses the following strategy:

- Retrieve the token from the environment variable`$WERCKER_TOKEN`
- Retrieve the token from file `~/.wercker/credentials`
- Fallback to anonymous user

If you retrieved your wercker token through other means, then you can use the
`credentials.Token` method to user a static `Provider`:

```golang
import "github.com/wercker/go-wercker-api"

token := "... your token ..."
options := wercker.Options{Creds: credentials.Token(token)}
client := wercker.NewClient(options)
```

[More information][auth-docs]

# FAQ

...

[api-docs]: http://devcenter.wercker.com/api/index.html
[auth-docs]: docs/authentication.md
