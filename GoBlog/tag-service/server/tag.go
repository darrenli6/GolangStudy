package server

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/darrenli6/GolangStudy/GoBlog/tag-service/pkg/bapi"
	pb "github.com/darrenli6/GolangStudy/GoBlog/tag-service/proto"
)

type TagServer struct {
}

func NewTagServer() *TagServer {
	return &TagServer{}
}

func (t *TagServer) GetTagList(ctx context.Context, r *pb.GetTagListRequest) (*pb.GetTagListReply, error) {

	api := bapi.NewAPI("http://localhost:8001")
	body, err := api.GetTagList(ctx, r.GetName())
	if err != nil {
		return nil, errors.New("error get tag list")
	}
	tagList := pb.GetTagListReply{}
	err = json.Unmarshal(body, &tagList)
	if err != nil {
		return nil, errors.New("failed ")
	}
	return &tagList, nil

}