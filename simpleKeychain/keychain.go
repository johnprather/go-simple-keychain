// Package simpleKeychain provides a few simple functions
// for working with generic passwords in the local OSX keychain.
// It is basically a very simple wrapper for
// github.com/keybase/go-keychain which does all the heavy
// lifting.
package simpleKeychain

import (
	"errors"

	keychain "github.com/keybase/go-keychain"
)

// ErrKeyChainItemNotFound is the error returned by Load when no match is found
var ErrKeyChainItemNotFound = errors.New("keychain item not found")

// Save will save a generic password item with the specified params.  Group is
// the unique access group determining which apps can access the data.  Name is
// the keychain item name.  Account is the account for which we are saving a
// password.  Password is the actual password to save.
func Save(group string, name string, account string, password string) error {
	// populate enough to delete any existing value
	item := keychain.NewItem()
	item.SetSecClass(keychain.SecClassGenericPassword)
	item.SetService(name)
	item.SetAccount(account)
	keychain.DeleteItem(item)

	// populate the rest of the object and save
	item.SetLabel(name)
	item.SetAccessGroup(group)
	item.SetData([]byte(password))
	item.SetSynchronizable(keychain.SynchronizableNo)
	item.SetAccessible(keychain.AccessibleAccessibleAlwaysThisDeviceOnly)

	err := keychain.AddItem(item)
	return err
}

// Load will attempt to dig up the password from the keychain for given params.
// Group is the unique access group.  Name is the keychain item name.  Account
// is the account for which we want the password.
func Load(group string, name string, account string) (pass string, err error) {
	item := keychain.NewItem()
	item.SetSecClass(keychain.SecClassGenericPassword)
	item.SetService(name)
	item.SetAccount(account)
	item.SetAccessGroup(group)
	item.SetMatchLimit(keychain.MatchLimitOne)
	item.SetReturnData(true)
	results, err := keychain.QueryItem(item)
	if err != nil {
		return
	}
	if len(results) != 1 {
		err = ErrKeyChainItemNotFound
		return
	}
	pass = string(results[0].Data)
	return
}

// Delete will attempt to delete the password item from the keychain.  Group is
// the access group.  Name is the keychain item name.  Account is the account
// for which we wish to delete the password.
func Delete(group string, name string, account string) (err error) {
	// populate enough to delete any existing value
	item := keychain.NewItem()
	item.SetSecClass(keychain.SecClassGenericPassword)
	item.SetService(name)
	item.SetAccount(account)
	err = keychain.DeleteItem(item)
	return
}
