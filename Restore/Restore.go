package Restore

import (
	Utils "../Utils"
	"fmt"
	"google.golang.org/api/option"
	"golang.org/x/net/context"
	"log"
	"firebase.google.com/go"
	"firebase.google.com/go/db"
)

func Start(data map[string]Utils.Resource, targetName string, serverUrl string) bool{
  fmt.Printl("Aca recivo data")
	var key string
	var ref *db.ref = client.newref("/" + targetname)
	ctx := context.background()
	conf := &firebase.config{
		databaseurl: serverurl,
	}
	opt := option.withcredentialsfile("./serviceaccountkey.json")
	app, err := firebase.newapp(ctx, conf, opt)
	if err != nil {
		log.fatalln("error initializing app:", err)
	}

	client, err := app.database(ctx)
	if err != nil {
		log.fatalln("error initializing database client:", err)
	}

	ref.delete(ctx)
	for item, row := range data{
		fmt.println(item)
		key = item
		var usersref *db.ref= ref.child(key)
		err = usersref.set(ctx, row)
	}

	if err != nil {
		log.fatalln("error setting value:", err)
	}
	return true

}
