# go-simple-keychain

Just a wrapper for github.com/keybase/go-keychain

# Usage

Import it:

```
import (
  ...

  "github.com/johnprather/go-simple-keychain/simpleKeychain"
)
```

Use it:

```
// a single access group is used by multiple applications to share
// one set of keychain data, this should be a globally unique value
// using your own domain and app/service/package name.

group := "group.com.github.johnprather.go-simple-keychain.example"

// a name for the keychain item (what is this a password for?)

name := "inventory api"

// load/generate these somehow
account := "root"
password := "somepass"

// will this password be only available on this device?

sync := false

// save a password
err := simpleKeychain.Save(group, name, account, password, false)


// load a password

pw, err := simpleKeychain.Load(group, name, account)
if err != nil {
  if err == simpleKeychain.ErrKeyChainItemNotFound {
    // no such keychain item exists (should save pass before loading)
  } else {
    // some other error with keychain api
  }
}
```
