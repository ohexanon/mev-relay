package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/kataras/iris/v12"
)

func handlePost(ctx iris.Context) {
	if ctx.Request().Header.Get("Content-Type") != "application/json" {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.Write([]byte("invalid content type"))
		return
	}
	defer ctx.Request().Body.Close()
	client := http.Client{}
	payload := new(Payload)
	if err := json.NewDecoder(ctx.Request().Body).Decode(payload); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.Write([]byte("bad payload format"))
		fmt.Println("could not submit payload", "error", err)
		return
	}

	if payload.Method != "eth_sendBundle" && payload.Method != "eth_sendPrivateRawTransaction" {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.Write([]byte("method not supported"))
		fmt.Println("could not submit payload", "payload", payload)
		return
	}

	body, err := json.Marshal(payload)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.Write([]byte("bad payload format"))
		fmt.Println("could not submit payload", "error", err)
		return
	}

	res, err := client.Post(os.Getenv("BUILDER_IP"), "application/json", bytes.NewReader(body))
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.Write([]byte("error submitting payload"))
		fmt.Println("could not submit payload", "error", err)
		return
	}
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.Write([]byte("error submitting payload"))
		fmt.Println("could not submit payload", "error", err)
		return
	}
	ctx.ContentType("application/json")
	ctx.Write(b)
}
