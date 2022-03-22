package matrix

import (
	"fmt"
	"matrixhook/env"

	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/event"
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

	fmt.Println("Logged in!")

	syncer := MatrixClient.Syncer.(*mautrix.DefaultSyncer)
	syncer.OnEventType(event.EventMessage, func(source mautrix.EventSource, evt *event.Event) {
		fmt.Printf("<%[1]s> %[4]s (%[2]s/%[3]s)\n", evt.Sender, evt.Type.String(), evt.ID, evt.Content.AsMessage().Body)
	})

	err = MatrixClient.Sync()
	if err != nil {
		panic(err)
	}

}
