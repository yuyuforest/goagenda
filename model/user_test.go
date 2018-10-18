package model

import (
	"github.com/spf13/viper"
	"testing"
)

func TestUserDB_GetUserByName_Empty(t *testing.T) {
	viper.Set("Database.Path", "../test/userfile")
	viper.Set("Database.UserFile", "empty.json")
	user := UserModel.GetUserByName("MegaShow")
	if user.Name != "" {
		t.Error("Test " + viper.GetString("Database.UserFile") + " fail.")
	}
}
