package matrix

import (
	"fmt"
	"matrixhook/env"

	"maunium.net/go/mautrix"
)

var MatrixClient *mautrix.Client

func ConnectBot() {
	var err error
	MatrixClient, err = mautrix.NewClient(env.Homeserver, "", "")
	if err != nil {
		panic(err)
	}

	_, err = MatrixClient.Login(&mautrix.ReqLogin{
		Type:             "m.login.password",
		Identifier:       mautrix.UserIdentifier{Type: mautrix.IdentifierTypeUser, User: env.Username},
		Password:         env.Password,
		StoreCredentials: true,
	})
	if err != nil {
		panic(err)
	}

	if env.DebugMode {
		fmt.Println("[Matrix] Logged in!")
	}
}
