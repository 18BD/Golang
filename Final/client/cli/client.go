package cli

import (
	"encoding/base64"
	"google.golang.org/grpc/metadata"
)

type Client struct {
	Metadata metadata.MD
	AuthString string
}

func New(login, password string) *Client {
	src := []byte(login + ":" + password)
	encoded := make([]byte, base64.StdEncoding.EncodedLen(len(src)))

	base64.StdEncoding.Encode(encoded, src)

	md := metadata.New(map[string]string{"Authorization": "Basic " + string(encoded)})

	cli := Client{
		Metadata: md,
		AuthString: string(encoded),
	}

	return &cli
}
