package simpleKeychain

import "testing"

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
	err := Save(testGroup, testName, testAccount, testPassword)
	if err != nil {
		t.Logf("Save(): %s", err.Error())
		t.Fail()
	}
}

func TestLoad(t *testing.T) {
	pass, err := Load(testGroup, testName, testAccount)
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
	err := Delete(testGroup, testName, testAccount)
	if err != nil {
		t.Logf("Delete(): %s", err.Error())
		t.Fail()
		return
	}
	_, err = Load(testGroup, testName, testAccount)
	if err != ErrKeyChainItemNotFound {
		t.Logf("Item still exists after deletion")
		t.Fail()
		return
	}
}
