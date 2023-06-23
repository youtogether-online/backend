package service

import (
	"context"
	"github.com/wtkeqrf0/you-together/internal/controller/dto"
)

type RoomPostgres interface {
	UpsertRoom(ctx context.Context, rm dto.Room, creatorId int) error
}

type RoomService struct {
	postgres RoomPostgres
}

func NewRoomService(postgres RoomPostgres) *RoomService {
	return &RoomService{postgres: postgres}
}

func (r *RoomService) UpsertRoom(rm dto.Room, creatorId int) error {
	return r.postgres.UpsertRoom(context.Background(), rm, creatorId)
}
