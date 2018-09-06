package Remitter



import (
	"fmt"
	"google.golang.org/api/option"
	"golang.org/x/net/context"
	"log"
	tt "../Reader"
	"firebase.google.com/go"
)

type User struct {
        Description string `json:"description"`
        DateOfBirth string `json:"date_of_birth,omitempty"`
        FullName    string `json:"full_name,omitempty"`
        Nickname    string `json:"nickname,omitempty"`
}

func SendDataToServer(data []tt.Resource) bool{
	ctx := context.Background()
	conf := &firebase.Config{
		DatabaseURL: "https://taskbook-18035.firebaseio.com",
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

	ref := client.NewRef("/tasks")
	/*var data map[string]interface{}
	if err := ref.Get(ctx, &data); err != nil {
		log.Fatalln("Error reading from database:", err)
	}
	fmt.Println(data)*/


	fmt.Println(data[0].Id)
	//usersRef := ref.Child("xxxx_2")
	//usersRef := ref.Child("_id_" + string(data[0].Id))
	err = usersRef.Set(ctx, map[string]*User{
		"alanisawesome": {
			Description: data[0].Description,
			DateOfBirth: "June 23, 1912",
			FullName:    "Alan Turing",
		},
	})
	if err != nil {
		log.Fatalln("Error setting value:", err)
	}
	return true

}

