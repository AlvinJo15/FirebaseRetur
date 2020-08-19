package firebaseclient

import (
	"context"
	"encoding/json"
	"firebase/internal/config"
	"firebase/pkg/errors"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var (
	sharedClient = &firebase.App{}
)

// Client ...
type Client struct {
	App    *firebase.App
	Client *firestore.Client
}

// NewClient ...
func NewClient(cfg *config.Config, credentials map[string]string) (*Client, error) {
	var c Client
	cb, err := json.Marshal(credentials)
	if err != nil {
		return &c, errors.Wrap(err, "[FIREBASE] Failed to marshal credentials!")
	}
	option := option.WithCredentialsJSON(cb)
	config := &firebase.Config{ProjectID: cfg.Firebase.ProjectID}
	c.App, err = firebase.NewApp(context.Background(), config, option)
	if err != nil {
		return &c, errors.Wrap(err, "[FIREBASE] Failed to initiate firebase client!")
	}
	c.Client, err = c.App.Firestore(context.Background())
	if err != nil {
		return &c, errors.Wrap(err, "[FIREBASE] Failed to initiate firestore client!")
	}
	return &c, err
}
