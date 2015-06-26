# Authentication

The current wercker api supports two types of authentication: using a token or a
username/password combination. The username/password combination however only
works for a few endpoints. This is the reason why most Providers only return a
token credential.

By default this client tries to fetch the credentials from
`~/.wercker/credentials`, the environment variable `$WERCKER_TOKEN` and falls
back to anonymous user.

If you retrieved your wercker token through other means, then you can use the
`credentials.Token` method to pass this value:

```golang
import "github.com/wercker/go-wercker-api"
import "github.com/wercker/go-wercker-api/credentials"

token := "... your token ..."
options := wercker.Options{Creds: credentials.Token(token)}
client := wercker.NewClient(options)
```

# credentials.Provider

This client retrieves the token by using the
[`Provider`](../credentials/provider.go) interface. This interface exposes a
single function: `GetCredentials() (*Creds, error)`.

`Creds` is a simple struct containg a `Token`, `Username`, and `Password`. A
provider should return a `Creds` containing either a `Token` or a `Username` and
a `Password`.

It is possible to create your own implementation of the `Provider` interface to
retrieve the credentials from a different backend.

## credentials.MultiProvider

The `MultiProvider` wraps an array of other `Provider` implementations. It will
go through all the implementations, and return the first `Creds` file. It will
ignore any errors returned from the wrapped implementations. If all wrapped
implementations fail to provide a `Creds` object, then the `MultiProvider` will
fail.

## credentials.EnvProvider

The `EnvProvider` retrieves the wercker token from a environment variable. By
default this is `WERCKER_TOKEN`. To overwrite the environment variable, change
the `key` property.

## credentials.StaticProvider

The `StaticProvider` returns a static `Creds` object, that is based on
the properties of the `StaticProvider`. This `Provider` is recommended for
situations wherer a user passes in the token through commandline arguments.
