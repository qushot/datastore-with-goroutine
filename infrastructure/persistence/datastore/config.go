package datastore

import (
	"context"

	"cloud.google.com/go/datastore"
)

const namespace = "goroutinetest"

// *datastore.Client はシングルトンにすべきなのかどうかよくわからない。
// いろんなサンプルコード見てもシングルトンにはしていない。
// でもシングルトンにして計測したらクッソ早くなった気がする。
var client *datastore.Client

func init() {
	var err error
	client, err = datastore.NewClient(context.Background(), datastore.DetectProjectID)
	if err != nil {
		panic(err)
	}
}
