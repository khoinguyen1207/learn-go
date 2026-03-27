package handler

import (
	"context"
	"encoding/json"

	"lynox/service"

	lynoxpb "proto-demo-repo/gen/lynox"

	"google.golang.org/protobuf/types/known/structpb"
)

type UserHandler struct {
	lynoxpb.UnimplementedLynoxServer
	svc *service.UserService
}
type GetUserRequest struct {
	Id int64
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func structToProto[T any](s *structpb.Struct, out *T) error {
	b, err := json.Marshal(s.AsMap())
	if err != nil {
		return err
	}
	return json.Unmarshal(b, out)
}

func (h *UserHandler) Index(ctx context.Context, req *lynoxpb.LynoxRequest) (*lynoxpb.LynoxResponse, error) {
	switch req.Service {
	case "get_user":
		var getUserReq GetUserRequest
		if err := structToProto(req.Data, &getUserReq); err != nil {
			return nil, err
		}
		return h.GetUser(ctx, &getUserReq)
	default:
		return &lynoxpb.LynoxResponse{
			Code: 1,
		}, nil
	}

}

func (h *UserHandler) GetUser(ctx context.Context, req *GetUserRequest) (*lynoxpb.LynoxResponse, error) {
	user, err := h.svc.GetUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &lynoxpb.LynoxResponse{
		Code:    0,
		Balance: int64(user.Balance),
	}, nil
}
