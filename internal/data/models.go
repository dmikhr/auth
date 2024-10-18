package data

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
)

// UserModel - модель данных пользователя
// контекст используется для передачи данных о соединении с БД
type UserModel struct {
	Ctx context.Context
}

// UserNew - структура для создания нового пользователя
type UserNew struct {
	Name      string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// UserRetrieve - структура для получения информации о существующем пользователе
// скрыта информация о пароле и доступен Id, т.к. пользователь уже существует
type UserRetrieve struct {
	ID        int64
	Name      string
	Email     string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// UserEdit - структура для внесения правок
// присутствует поле UpdatedAt для сохранения времени изменений
type UserEdit struct {
	ID        int64
	Name      string
	Email     string
	Role      string
	UpdatedAt time.Time
}

// Models - wrapper на случай, если в будущем появятся новые модели данных
type Models struct {
	User UserModel
}

// Create - создание нового пользователя, возвращает его id
func (m UserModel) Create(user *UserNew) (int64, error) {
	if !isValidEmail(user.Email) {
		return -1, fmt.Errorf("invalid email")
	}
	query := `INSERT INTO users (name, email, password, role, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	// достаем БД соеднинение из контекста
	conn := m.Ctx.Value(DBConn).(*pgx.Conn)
	var userID int64
	err := conn.QueryRow(m.Ctx, query,
		user.Name, user.Email, user.Password, user.Role,
		user.CreatedAt, user.UpdatedAt).Scan(&userID)
	if err != nil {
		return -1, err
	}
	return userID, nil
}

// Get - получить данные о пользователе по id
func (m UserModel) Get(id int64) (UserRetrieve, error) {
	query := `SELECT id, name, email, role, created_at, updated_at FROM users WHERE id = $1`
	conn := m.Ctx.Value(DBConn).(*pgx.Conn)
	var user UserRetrieve
	err := conn.QueryRow(m.Ctx, query, id).Scan(&user.ID,
		&user.Name, &user.Email, &user.Role,
		&user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return UserRetrieve{}, err
	}
	return user, nil
}

// Update - обновить данные о пользователе, UpdatedAt обновляется до time.Now()
func (m UserModel) Update(user *UserEdit) error {
	if !isValidEmail(user.Email) {
		return fmt.Errorf("invalid email")
	}
	query := `UPDATE users SET id=$1, name=$2, email=$3, role=$4, updated_at=$5 WHERE id=$1`
	conn := m.Ctx.Value(DBConn).(*pgx.Conn)
	cmdTag, err := conn.Exec(m.Ctx, query, user.ID, user.Name, user.Email, user.Role, user.UpdatedAt)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("nothing was updated")
	}
	return nil
}

// Delete - удаление пользователя по id
func (m UserModel) Delete(id int64) error {
	query := `DELETE FROM users WHERE id = $1`
	conn := m.Ctx.Value(DBConn).(*pgx.Conn)
	cmdTag, err := conn.Exec(m.Ctx, query, id)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("nothing was deleted")
	}
	return nil
}

// InitModels - инициализация моделей данных
func InitModels(ctx context.Context) Models {
	return Models{
		User: UserModel{Ctx: ctx},
	}
}
