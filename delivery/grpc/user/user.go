package user

import (
	"context"
	"github.com/google/uuid"
	"github.com/sergey4qb/mf1-test/dto"
	"github.com/sergey4qb/mf1-test/services/user"

	"github.com/sergey4qb/mf1-test/model"
	pb "github.com/sergey4qb/mf1-test/proto/pb"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
	userService user.User
}

func NewUserServer(userService user.User) *UserServiceServer {
	return &UserServiceServer{
		userService: userService,
	}
}

func (s *UserServiceServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	u := &model.User{
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}

	if err := s.userService.Create(ctx, u); err != nil {
		return nil, err
	}

	resp := &pb.CreateUserResponse{
		User: &pb.User{
			Id:    u.ID.String(),
			Name:  u.Name,
			Email: u.Email,
		},
	}
	return resp, nil
}

func (s *UserServiceServer) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	users, err := s.userService.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var pbUsers []*pb.User
	for _, u := range users {
		pbUsers = append(pbUsers, &pb.User{
			Id:    u.ID.String(),
			Name:  u.Name,
			Email: u.Email,
		})
	}

	resp := &pb.ListUsersResponse{
		Users: pbUsers,
	}

	return resp, nil
}

func (s *UserServiceServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}
	u, err := s.userService.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	resp := &pb.GetUserResponse{
		User: &pb.User{
			Id:    u.ID.String(),
			Name:  u.Name,
			Email: u.Email,
		},
	}
	return resp, nil
}
func (s *UserServiceServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}
	var namePtr *string
	if req.GetName() != "" {
		name := req.GetName()
		namePtr = &name
	}
	var emailPtr *string
	if req.GetEmail() != "" {
		email := req.GetEmail()
		emailPtr = &email
	}
	dto := &dto.UpdateUserDTO{
		ID:    id,
		Name:  namePtr,
		Email: emailPtr,
	}
	updatedUser, err := s.userService.Update(ctx, dto)
	if err != nil {
		return nil, err
	}
	resp := &pb.UpdateUserResponse{
		User: &pb.User{
			Id:    updatedUser.ID.String(),
			Name:  updatedUser.Name,
			Email: updatedUser.Email,
		},
	}
	return resp, nil
}

func (s *UserServiceServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}
	if err := s.userService.Delete(ctx, id); err != nil {
		return nil, err
	}
	resp := &pb.DeleteUserResponse{}
	return resp, nil
}
