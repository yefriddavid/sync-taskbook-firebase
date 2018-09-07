package main

import (
  "fmt"
	"google.golang.org/api/option"
	"golang.org/x/net/context"
	"log"
	"firebase.google.com/go"
	//"firebase.google.com/go/db"
	"os/user"
  Utils "../../Utils"
)

func main(){
	usr, _ := user.Current()
  ctx := context.Background()
  conf := &firebase.Config{
          DatabaseURL: "https://taskbook-18035.firebaseio.com",
  }
  // Fetch the service account key JSON file contents
  opt := option.WithCredentialsFile(usr.HomeDir + "/.ssh/serviceAccountKey-taskbook.json")

  // Initialize the app with a service account, granting admin privileges
  app, err := firebase.NewApp(ctx, conf, opt)
  if err != nil {
          log.Fatalln("Error initializing app:", err)
  }

  client, err := app.Database(ctx)
  if err != nil {
          log.Fatalln("Error initializing database client:", err)
  }

  // As an admin, the app has access to read and write all data, regradless of Security Rules
  ref := client.NewRef("/archive_Microvoz")
  //var data map[string]interface{}
  var data map[string]Utils.Resource
  //var data map[Utils.Resource]interface{}
  if err := ref.Get(ctx, &data); err != nil {
          log.Fatalln("Error reading from database:", err)
  }
  fmt.Println(data)

}
