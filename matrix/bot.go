package matrix

import (
	"fmt"
	"matrixhook/env"

	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/id"
)

var MatrixClient *mautrix.Client

func ConnectBot() {
	MatrixClient, err := mautrix.NewClient(env.Homeserver, "", "")
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

	fmt.Println("[Matrix] Logged in!")
	MatrixClient.SendText(id.RoomID(env.RoomId), "Logged in!")

	err = MatrixClient.Sync()
	if err != nil {
		panic(err)
	}

}
