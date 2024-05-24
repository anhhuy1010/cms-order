package service

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	// pbStore "github.com/Diqit-A1-Branch/cpos-microservice-store/grpc/proto/store"

	pb "github.com/anhhuy1010/cms-order/grpc/proto/user"

	// "github.com/Diqit-A1-Branch/cpos-microservice-tenant/helpers/util"
	"github.com/anhhuy1010/cms-order/models"
)

type UserService struct {
}

func NewUserServer() pb.UserServer {
	return &UserService{}
}

func (s *UserService) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	conditions := bson.M{}

	result, err := new(models.Users).Find(conditions)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	var respData []*pb.DetailResponse
	for _, user := range result {

		res := &pb.DetailResponse{
			Uuid:     user.Uuid,
			Username: user.Username,
			IsActive: int32(user.IsActive),
		}
		respData = append(respData, res)
	}

	return &pb.ListResponse{
		Users: respData,
	}, nil
}
