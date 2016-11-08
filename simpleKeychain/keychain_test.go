package simpleKeychain

import "testing"

const (
	testGroup    = "group.com.github.johnprather.go-simple-keychain"
	testName     = "test item"
	testAccount  = "testuser"
	testPassword = "a test password string"
)

func TestSave(t *testing.T) {
	err := Save(testGroup, testName, testAccount, testPassword, false)
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
	}
	if pass != testPassword {
		t.Logf("Loaded password does not match saved password")
		t.Fail()
	}
}

func TestDelete(t *testing.T) {
	err := Delete(testGroup, testName, testAccount)
	if err != nil {
		t.Logf("Delete(): %s", err.Error())
		t.Fail()
	}
}
