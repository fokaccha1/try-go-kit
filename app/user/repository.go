package user

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type UserRepository interface {
	Find(ctx context.Context, id int) (user User, err error)
	FindAll(context.Context) ([]User, error)
	Store(context.Context, UserDraft) (id int, err error)
}

type userRepository struct {
	conn *sql.DB
}

func NewUserRepository() UserRepository {
	connection, err := sql.Open("mysql", "go-kit:pass@tcp(db:3306)/go-kit")
	if err != nil {
		panic(err.Error)
	}
	return &userRepository{conn: connection}
}

func (repo *userRepository) Find(ctx context.Context, id int) (user User, err error) {
	user = User{}
	row := repo.conn.QueryRow("SELECT id, name, age FROM users WHERE id = ?", id)
	err = row.Scan(&user.Id, &user.Name, &user.Age)
	return
}

func (repo *userRepository) FindAll(context.Context) (users []User, err error) {
	rows, err := repo.conn.Query("SELECT id, name, age FROM users")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		user := User{}
		if err := rows.Scan(&user.Id, &user.Name, &user.Age); err != nil {
			continue
		}
		users = append(users, user)
	}
	return
}

func (repo *userRepository) Store(ctx context.Context, ud UserDraft) (id int, err error) {
	res, err := repo.conn.Exec("INSERT INTO users (name, age) VALUES (?,?)", ud.Name, ud.Age)
	if err != nil {
		id = 0
		return
	}
	var id64 int64
	id64, err = res.LastInsertId()
	id = int(id64)
	return
}
