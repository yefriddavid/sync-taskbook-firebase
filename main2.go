import (
  "log"

  firebase "firebase.google.com/go"
  "google.golang.org/api/option"
)

// Use a service account
sa := option.WithCredentialsFile("path/to/serviceAccount.json")
app, err := firebase.NewApp(ctx, nil, sa)
if err != nil {
  log.Fatalln(err)
}

client, err := app.Firestore(ctx)
if err != nil {
  log.Fatalln(err)
}


_, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
        "first": "Ada",
        "last":  "Lovelace",
        "born":  1815,
})
if err != nil {
  log.Fatalf("Failed adding alovelace: %v", err)
}



defer client.Close()
