package service

import (
	"context"
	"github.com/wtkeqrf0/you-together/ent"
	"github.com/wtkeqrf0/you-together/internal/controller/dto"
)

type RoomPostgres interface {
	Create(ctx context.Context, rm dto.Room, creatorId int) (*ent.Room, error)
}

type RoomService struct {
	postgres RoomPostgres
}

func NewRoomService(postgres RoomPostgres) *RoomService {
	return &RoomService{postgres: postgres}
}

func (r *RoomService) Create(rm dto.Room, creatorId int) (*ent.Room, error) {
	return r.postgres.Create(context.Background(), rm, creatorId)
}
