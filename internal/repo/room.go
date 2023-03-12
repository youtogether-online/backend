package repo

/*

import (
	"you_together/ent"
	"you_together/internal/controller/dto"
	"context"
	"golang.org/x/crypto/bcrypt"
)

type RoomStorage struct {
	roomClient *ent.RoomClient
}

func (r RoomStorage) CreateRoom(ctx context.Context, dto dto.CreateRoomDTO) (int, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(dto.Password), 12)

	save, err := r.roomClient.Create().SetEmail(dto.Email).SetPasswordHash(hashedPassword).Save(ctx)
	if err != nil {
		return 0, err
	}
	return save.ID, nil
}

func (r RoomStorage) FindRoomByUUID(ctx context.Context, id int) (*ent.User, error) {
	return r.userClient.Get(ctx, id)
}

func (r RoomStorage) FindAllRooms(ctx context.Context, limit int) ([]*ent.User, error) {
	return r.userClient.Query().Limit(limit).All(ctx)
}

func (r RoomStorage) UpdateRoom(ctx context.Context, user *ent.User) error {
	return r.userClient.UpdateOne(user).Exec(ctx)
}

func (r RoomStorage) DeleteRoom(ctx context.Context, id int) error {
	return r.userClient.DeleteOneID(id).Exec(ctx)
}

func NewRoomStorage(roomClient *ent.RoomClient) RoomStorage {
	return RoomStorage{roomClient: roomClient}
}
*/
