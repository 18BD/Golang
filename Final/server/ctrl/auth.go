package ctrl

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"net/http"
	"strings"
)

const (
	auth_host = "localhost"
	auth_port = ":3000"
)

func authMiddleware(ctx context.Context) (context.Context, error) {
	credsEnc, err := grpc_auth.AuthFromMD(ctx, "Basic")
	if err != nil {
		return nil, err
	}

	// Base64 Standard Decoding
	credsDec, err := base64.StdEncoding.DecodeString(credsEnc)
	if err != nil {
		return nil, err
	}

	arr := strings.Split(string(credsDec), ":")

	values := map[string]string{"login": arr[0], "password": arr[1]}

	json_data, err := json.Marshal(values)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post("http://" + auth_host + auth_port + "/", "application/json",
		bytes.NewBuffer(json_data),
	)

	if err != nil {
		return nil, err
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	return context.Background(), nil
}
