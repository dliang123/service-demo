package service

import (
	"context"
	"net/http"
	"encoding/json"
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
)

func NewSayHelloHandler(svr Service) *httptransport.Server  {
	return httptransport.NewServer(
		makeSayHelloService(svr),
		DecodeSayHelloRequest,
		EncodeSayHelloResponse,
		//httptransport.ServerBefore(utils.PrintRequestBody),
		//httptransport.ServerErrorLogger(logger),
	)
}


func DecodeSayHelloRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println(r.Header)
	req := &SayHelloRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return req, err
	}
	return req, nil
}

func EncodeSayHelloResponse(_ context.Context, w http.ResponseWriter, response interface{}) (err error) {
	//utils.SetHuaWeiMessageHeader(w, "f37043a59d8a49baaba13b961a542672", "09ae0b71b63d417e8fcb8dde719c7c23")
	if response != nil {
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			return err
		}
	}

	return nil
}
