package service

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	SatHelloEndpoint endpoint.Endpoint
}

func NewEndPoint(svc Service) (ep Endpoints) {

	ep.SatHelloEndpoint = makeSayHelloService(svc)

	return ep

}

func makeSayHelloService(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SayHelloRequest)
		v, err := svc.sayHello(req.content)
		if err != nil {
			return SayHelloResponse{v, err.Error()}, nil
		}
		return SayHelloResponse{v, ""}, nil
	}
}
