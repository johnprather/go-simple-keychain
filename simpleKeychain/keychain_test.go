package simpleKeychain_test

import (
	"fmt"
	"testing"

	"github.com/johnprather/go-simple-keychain/simpleKeychain"
)

type TestSet struct {
	TestName string
	Group    string
	Name     string
	Account  string
	Password string
	Sync     bool
}

const (
	testGroup    = "group.com.github.johnprather.go-simple-keychain"
	testName     = "a test item"
	testAccount  = "testuser"
	testPassword = "a test password string"
)

func TestSave(t *testing.T) {
	err := simpleKeychain.Save(testGroup, testName, testAccount, testPassword)
	if err != nil {
		t.Logf("Save(): %s", err.Error())
		t.Fail()
	}
}

func TestLoad(t *testing.T) {
	pass, err := simpleKeychain.Load(testGroup, testName, testAccount)
	if err != nil {
		t.Logf("Load(): %s", err.Error())
		t.Fail()
		return
	}
	if pass != testPassword {
		t.Logf("Loaded password does not match saved password")
		t.Fail()
		return
	}
}

func TestDelete(t *testing.T) {
	err := simpleKeychain.Delete(testGroup, testName, testAccount)
	if err != nil {
		t.Logf("Delete(): %s", err.Error())
		t.Fail()
		return
	}
	_, err = simpleKeychain.Load(testGroup, testName, testAccount)
	if err != simpleKeychain.ErrKeyChainItemNotFound {
		t.Logf("Item still exists after deletion")
		t.Fail()
		return
	}
}

func ExampleSave() {
	group := "group.com.github.johnprather.go-simple-keychain.example"
	name := "Inventory API"
	account := "root"
	password := "a password string"

	err := simpleKeychain.Save(group, name, account, password)
	if err != nil {
		fmt.Printf("Save(): %s\n", err.Error())
	} else {
		fmt.Println("Password saved.")
	}
}

func ExampleLoad() {
	group := "group.com.github.johnprather.go-simple-keychain.example"
	name := "Inventory API"
	account := "root"

	pass, err := simpleKeychain.Load(group, name, account)
	if err != nil {
		if err == simpleKeychain.ErrKeyChainItemNotFound {
			// handle no-such-keychain-item (save before load?)
			return
		}
		// handle some keychain error err
		return
	}
	fmt.Printf("Read a password of %d characters.\n", len(pass))
}

func ExampleDelete() {
	group := "group.com.github.johnprather.go-simple-keychain.example"
	name := "Inventory API"
	account := "root"

	err := simpleKeychain.Delete(group, name, account)
	if err != nil {
		// handle some keychain error err
		return
	}
	fmt.Println("Password deleted.")
}
