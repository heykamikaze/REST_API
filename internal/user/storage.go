package user

import "context"

type Storage interface {
	Create(ctx context.Context, user User) (string, error) //метод создания пользователя
	FindOne(ctx context.Context, id string) (User, error)  //метод поиска пользователя
	Update(ctx context.Context, user User) error           //метод смены данных пользователя
	Delete(ctx context.Context, id string) error           //метод удаления пользователя
}
