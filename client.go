package main

import (
	"context"
	"fmt"
	"github.com/rsocket/rsocket-go/extension"
	"log"

	"github.com/rsocket/rsocket-go"
	_ "github.com/rsocket/rsocket-go/extension"
	"github.com/rsocket/rsocket-go/payload"
)

func main() {
	au, err := extension.NewAuthentication("bearer", []byte("eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJSU29ja2V0QnJva2VyIiwic3ViIjoicmVxdWVzdCIsImF1ZCI6WyJhIiwiYiIsInIiLCJyIiwieSJdLCJpYXQiOjE3MDQ4MTIxMjEsImlkIjoiMWJjNGIzMzEtYTQ5ZC00YjU2LWFiMDAtNWI2YzMwNzEzNTkzIiwic2FzIjpbImIiLCJhIiwiciIsInIiLCJ5Il0sIm9yZ3MiOlsiMSJdLCJyb2xlcyI6WyJhIiwiZCIsImkiLCJtIiwibiJdLCJhdXRob3JpdGllcyI6WyJhIiwiZCIsIm0iLCJpIiwibiJdfQ.nDcbnS0f9Vct7M3HTMtWT6i0-MdQ6rF5xjeWQeUzXMLUMN1Cr2GxbkR67KEynbQvK-u1xwrdWzQMV_cZutz6SiK76QGU2sUaF2AYhTHD_99-jMrHXdG9rfTR0fjFTEavDovfrg378-3xShrOan1943m-0TB4JiRgnrzM4AqUOufOV14-hmjR0Is7hwoqQRZQToV2wE630-_KfZS6iWKizB3dqy560A6YLlE2sQNjKmcfUiJ35A9em9Bs8-aCtzxAGrSXPGCsy5PJeAf1Itkbumx31uXn4h2bOCDyEOEaNeTG1qYDZXYdIXiPNXSiebjxg4dXSBlqvecYuMznhTP4UA"))
	if err != nil {
		panic(err)
	}

	cm, err := extension.NewCompositeMetadataBuilder().
		PushWellKnown(extension.MessageAuthentication, au.Payload()).
		Build()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(cm))

	// Connect to server
	cli, err := rsocket.Connect().
		DataMimeType("message/x.rsocket.application+json").
		MetadataMimeType(extension.MessageCompositeMetadata.String()).
		SetupPayload(payload.New([]byte("{}"), cm)).
		Transport(rsocket.TCPClient().SetHostAndPort("127.0.0.1", 9999).Build()).
		Start(context.Background())
	if err != nil {
		panic(err)
	}

	defer cli.Close()

	// Send request
	result, err := cli.RequestResponse(payload.NewString("你好", "世界")).Block(context.Background())
	if err != nil {
		panic(err)
	}
	log.Println("response:", string(result.Data()))
}

//eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJSU29ja2V0QnJva2VyIiwic3ViIjoicmVxdWVzdCIsImF1ZCI6WyJhIiwiYiIsInIiLCJyIiwieSJdLCJpYXQiOjE3MDQ4MTIxMjEsImlkIjoiMWJjNGIzMzEtYTQ5ZC00YjU2LWFiMDAtNWI2YzMwNzEzNTkzIiwic2FzIjpbImIiLCJhIiwiciIsInIiLCJ5Il0sIm9yZ3MiOlsiMSJdLCJyb2xlcyI6WyJhIiwiZCIsImkiLCJtIiwibiJdLCJhdXRob3JpdGllcyI6WyJhIiwiZCIsIm0iLCJpIiwibiJdfQ.nDcbnS0f9Vct7M3HTMtWT6i0-MdQ6rF5xjeWQeUzXMLUMN1Cr2GxbkR67KEynbQvK-u1xwrdWzQMV_cZutz6SiK76QGU2sUaF2AYhTHD_99-jMrHXdG9rfTR0fjFTEavDovfrg378-3xShrOan1943m-0TB4JiRgnrzM4AqUOufOV14-hmjR0Is7hwoqQRZQToV2wE630-_KfZS6iWKizB3dqy560A6YLlE2sQNjKmcfUiJ35A9em9Bs8-aCtzxAGrSXPGCsy5PJeAf1Itkbumx31uXn4h2bOCDyEOEaNeTG1qYDZXYdIXiPNXSiebjxg4dXSBlqvecYuMznhTP4UA
