package Remitter



import (
	"fmt"
	"google.golang.org/api/option"
	"golang.org/x/net/context"
	"log"
	Utils "../Utils"
	"firebase.google.com/go"
	"firebase.google.com/go/db"
)


func SendDataToServer(data []Utils.Resource, targetName string) bool{
	//rar tag string = "Personal"
	var serverUrl string = "https://taskbook-18035.firebaseio.com"
	ctx := context.Background()
	conf := &firebase.Config{
		DatabaseURL: serverUrl,
	}
	opt := option.WithCredentialsFile("./serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalln("Error initializing app:", err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}

	//var ref *db.Ref = client.NewRef("/storage_" + tag)
	var ref *db.Ref = client.NewRef("/" + targetName)
	//ref.Update(ctx, nil)
	ref.Delete(ctx)


	var key string

	for i := 0; i < len(data); i++{
		key = fmt.Sprintf("%d", data[i].Id)
		var usersRef *db.Ref= ref.Child(key)
		fmt.Println(key)
		//err = usersRef.Set(ctx, nil)
		err = usersRef.Set(ctx, data[i])
	}

	if err != nil {
		log.Fatalln("Error setting value:", err)
	}
	return true

}

