package simpleKeychain

import (
	"errors"

	keychain "github.com/keybase/go-keychain"
)

// ErrKeyChainItemNotFound is the returned by Load when no match is found
var ErrKeyChainItemNotFound = errors.New("keychain item not found")

// Save will save a generic password item with the specified params
//   group: the AccessGroup, used for having several apps in the same group
//          to share the same keychain data.
//   name: the name of the keychain item
//   user: the account for the keychain item
//   pass: the password to save
//   sync: the keychain item may be synced (icloud?), use false to keep local
func Save(group string, name string, user string, pass string, sync bool) error {
	// populate enough to delete any existing value
	item := keychain.NewItem()
	item.SetSecClass(keychain.SecClassGenericPassword)
	item.SetService(name)
	item.SetAccount(user)
	keychain.DeleteItem(item)

	// populate the rest of the object and save
	item.SetLabel(name)
	item.SetAccessGroup(group)
	item.SetData([]byte(pass))
	if sync {
		item.SetSynchronizable(keychain.SynchronizableYes)
		item.SetAccessible(keychain.AccessibleAlways)
	} else {
		item.SetSynchronizable(keychain.SynchronizableNo)
		item.SetAccessible(keychain.AccessibleAccessibleAlwaysThisDeviceOnly)
	}

	item.SetAccessible(keychain.AccessibleAlways)
	err := keychain.AddItem(item)
	return err
}

// Load will attempt to dig up the password from the keychain for given params
//    group: the access group, used by a suite of apps to share the same
//           keychain data
//    name: the keychain item name
//    user: the keychain item account
func Load(group string, account string, user string) (pass string, err error) {
	item := keychain.NewItem()
	item.SetSecClass(keychain.SecClassGenericPassword)
	item.SetService(account)
	item.SetAccount(user)
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

// Delete will attempt to delete the password item from the keychain
//   group: the access group, used by suite of apps to share same keychain data
//   name: the item name
//   account: the account name
func Delete(group string, name string, account string) (err error) {
	// populate enough to delete any existing value
	item := keychain.NewItem()
	item.SetSecClass(keychain.SecClassGenericPassword)
	item.SetService(name)
	item.SetAccount(account)
	err = keychain.DeleteItem(item)
	return
}
