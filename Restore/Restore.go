package Remitter

import (
  "fmt"
	"google.golang.org/api/option"
	"golang.org/x/net/context"
	"log"
	//Utils "../Utils"
	"firebase.google.com/go"
	"firebase.google.com/go/db"
	"os/user"
)


func Start(targetName string, serverUrl string) bool{
	ctx := context.Background()
	conf := &firebase.Config{
		DatabaseURL: serverUrl,
	}

	usr, _ := user.Current()
	opt := option.WithCredentialsFile(usr.HomeDir + "/.ssh/serviceAccountKey-taskbook.json")
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalln("Error initializing app:", err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}

	var ref *db.Ref = client.NewRef("/" + targetName)
  fmt.Println("----")
  fmt.Print(ref)


	return true

}

