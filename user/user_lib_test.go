package user

import (
	"encoding/json"
	"fmt"
	"types"
	"testing"
)

func TestNewUser(t *testing.T) {
	uJson, _ := json.Marshal(types.NewEncryptedUser("tetsUser", "testPass"))
	fmt.Println(string(uJson))
}
