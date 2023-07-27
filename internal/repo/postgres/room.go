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

func (r *RoomStorage) UpsertRoom(ctx context.Context, rm dto.Room, creatorId int) (int, error) {
	cl := r.roomClient.Create().
		SetOwnerID(creatorId).
		SetNillablePrivacy(rm.Privacy).
		SetNillableDescription(rm.Description).
		SetTitle(rm.Title)

	if rm.Password != nil {
		cl.SetPasswordHash([]byte(*rm.Password))
	}

	return cl.OnConflict().UpdateNewValues().ID(ctx)
}

func (r *RoomStorage) GetRoomById(ctx context.Context, roomId int) (*ent.Room, error) {
	return r.roomClient.Get(ctx, roomId)
}
