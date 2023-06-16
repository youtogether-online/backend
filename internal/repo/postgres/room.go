package postgres

import (
	"context"
	"github.com/wtkeqrf0/you-together/ent"
	"github.com/wtkeqrf0/you-together/internal/controller/dto"
)

type RoomStorage struct {
	roomClient *ent.RoomClient
}

func NewRoomStorage(roomClient *ent.RoomClient) *RoomStorage {
	return &RoomStorage{roomClient: roomClient}
}

func (r *RoomStorage) Create(ctx context.Context, rm dto.Room, creatorId int) (*ent.Room, error) {
	cl := r.roomClient.Create().
		SetNillableCustomName(rm.CustomName).
		SetNillableDescription(rm.Description).
		SetNillablePrivacy(rm.Privacy).
		SetNillableSetChat(rm.HasChat).
		SetOwnerID(creatorId)

	if rm.Password != nil {
		cl.SetPasswordHash([]byte(*rm.Password))
	}

	return cl.Save(ctx)
}
