package core

import (
	"context"
	"fmt"
	systemicdbGrpc "github.com/SamuelBanksTech/SystemicDB-Server/protobuf_files"
	"github.com/SamuelBanksTech/SystemicDB-Server/systemicdb_service"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

type getserver struct {
	systemicdbGrpc.UnimplementedGetServer
	systemicdbService *systemicdb_service.SystemicDBService
}

type setserver struct {
	systemicdbGrpc.UnimplementedSetServer
	systemicdbService *systemicdb_service.SystemicDBService
}

type delserver struct {
	systemicdbGrpc.UnimplementedDelServer
	systemicdbService *systemicdb_service.SystemicDBService
}

func (g *getserver) ByKey(ctx context.Context, in *systemicdbGrpc.GetByKeyRequest) (*systemicdbGrpc.SingleResponse, error) {

	var returnData systemicdbGrpc.SingleResponse

	datum, err := g.systemicdbService.GetByKey(in.Key)
	if err != nil {
		return nil, err
	}

	returnData.Datum = datum

	return &returnData, nil
}

func (g *getserver) ByTag(ctx context.Context, in *systemicdbGrpc.GetByTagRequest) (*systemicdbGrpc.RepeatedResponse, error) {
	var returnData systemicdbGrpc.RepeatedResponse

	fmt.Println(in.Tag)

	returnData.Datum = []string{"hello", "this", "is", "data"}

	return &returnData, nil
}

func (s *setserver) ByKey(ctx context.Context, in *systemicdbGrpc.SetByKeyRequest) (*systemicdbGrpc.BoolResponse, error) {

	var returnData systemicdbGrpc.BoolResponse

	if err := s.systemicdbService.SetByKey(in.Key, in.Datum, in.ExpirySeconds); err != nil {
		return &returnData, err
	}

	returnData.Ok = true

	return &returnData, nil
}

func (s *setserver) ByTag(ctx context.Context, in *systemicdbGrpc.SetByTagRequest) (*systemicdbGrpc.BoolResponse, error) {
	var returnData systemicdbGrpc.BoolResponse

	fmt.Println(in.Datum, in.Tag, in.Key)

	return &returnData, nil
}

func (d *delserver) ByKey(ctx context.Context, in *systemicdbGrpc.GetByKeyRequest) (*systemicdbGrpc.BoolResponse, error) {

	var returnData systemicdbGrpc.BoolResponse

	err := d.systemicdbService.DelByKey(in.Key)
	if err != nil {
		return nil, err
	}

	returnData.Ok = true

	return &returnData, nil
}

func (d *delserver) ByTag(ctx context.Context, in *systemicdbGrpc.GetByTagRequest) (*systemicdbGrpc.BoolResponse, error) {
	var returnData systemicdbGrpc.BoolResponse

	fmt.Println(in.Tag)

	return &returnData, nil
}

func InitGrpcServer(port int, systemicdbService *systemicdb_service.SystemicDBService) {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatalln(err)
	}

	gs := grpc.NewServer()
	systemicdbGrpc.RegisterGetServer(gs, &getserver{systemicdbService: systemicdbService})
	systemicdbGrpc.RegisterSetServer(gs, &setserver{systemicdbService: systemicdbService})
	systemicdbGrpc.RegisterDelServer(gs, &delserver{systemicdbService: systemicdbService})

	if err := gs.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
