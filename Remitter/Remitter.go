package Remitter



import (
	"google.golang.org/api/option"
	"golang.org/x/net/context"
	"log"
	Utils "../Utils"
	"firebase.google.com/go"
	"firebase.google.com/go/db"
	"os/user"
)


func SendDataToServer(data map[string]Utils.Resource, targetName string, serverUrl string) bool{
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
	ref.Delete(ctx)

  for key, resource := range data{
		var usersRef *db.Ref= ref.Child(key)
		err = usersRef.Set(ctx, resource)
	}

	if err != nil {
		log.Fatalln("Error setting value:", err)
	}
	return true

}

