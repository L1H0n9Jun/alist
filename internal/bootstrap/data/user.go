package data

import (
	"os"

	"alist/cmd/flags"
	"alist/internal/db"
	"alist/internal/model"
	"alist/internal/op"
	"alist/pkg/utils"
	"alist/pkg/utils/random"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func initUser() {
	admin, err := op.GetAdmin()
	adminPassword := random.String(8)
	envpass := os.Getenv("ALIST_ADMIN_PASSWORD")
	if flags.Dev {
		adminPassword = "admin"
	} else if len(envpass) > 0 {
		adminPassword = envpass
	}
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			salt := random.String(16)
			admin = &model.User{
				Username: "admin",
				Salt:     salt,
				PwdHash:  model.TwoHashPwd(adminPassword, salt),
				Role:     model.ADMIN,
				BasePath: "/",
				Authn:    "[]",
			}
			if err := op.CreateUser(admin); err != nil {
				panic(err)
			} else {
				utils.Log.Infof("Successfully created the admin user and the initial password is: %s", adminPassword)
			}
		} else {
			utils.Log.Fatalf("[init user] Failed to get admin user: %v", err)
		}
	}
	guest, err := op.GetGuest()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			salt := random.String(16)
			guest = &model.User{
				Username:   "guest",
				PwdHash:    model.TwoHashPwd("guest", salt),
				Salt:       salt,
				Role:       model.GUEST,
				BasePath:   "/",
				Permission: 0,
				Disabled:   true,
				Authn:      "[]",
			}
			if err := db.CreateUser(guest); err != nil {
				utils.Log.Fatalf("[init user] Failed to create guest user: %v", err)
			}
		} else {
			utils.Log.Fatalf("[init user] Failed to get guest user: %v", err)
		}
	}
	hashPwdForOldVersion()
	updateAuthnForOldVersion()
}

func hashPwdForOldVersion() {
	users, _, err := op.GetUsers(1, -1)
	if err != nil {
		utils.Log.Fatalf("[hash pwd for old version] failed get users: %v", err)
	}
	for i := range users {
		user := users[i]
		if user.PwdHash == "" {
			user.SetPassword(user.Password)
			user.Password = ""
			if err := db.UpdateUser(&user); err != nil {
				utils.Log.Fatalf("[hash pwd for old version] failed update user: %v", err)
			}
		}
	}
}

func updateAuthnForOldVersion() {
	users, _, err := op.GetUsers(1, -1)
	if err != nil {
		utils.Log.Fatalf("[update authn for old version] failed get users: %v", err)
	}
	for i := range users {
		user := users[i]
		if user.Authn == "" {
			user.Authn = "[]"
			if err := db.UpdateUser(&user); err != nil {
				utils.Log.Fatalf("[update authn for old version] failed update user: %v", err)
			}
		}
	}
}
