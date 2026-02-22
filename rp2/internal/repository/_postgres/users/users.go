package users

import (
	"fmt"
	"time"

	"rp2/internal/repository/_postgres"
	"rp2/pkg/modules"
)

type Repository struct {
	db               *_postgres.Dialect
	executionTimeout time.Duration
}

func NewUserRepository(db *_postgres.Dialect) *Repository {
	return &Repository{
		db:               db,
		executionTimeout: time.Second * 5,
	}
}

func (r *Repository) GetUsers(limit, offset int) ([]modules.User, error) {
	var users []modules.User
	//err := r.db.DB.Select(&users, "select * from users")
	err := r.db.DB.Select(&users, "select * from users where deleted_at is null order by id limit $1 offset $2", limit, offset)
	if err != nil {
		return nil, err
	}

	//fmt.Println(users)
	return users, nil
}

func (r *Repository) GetUserByID(id int) (*modules.User, error) {
	var user modules.User
	err := r.db.DB.Get(&user, "select * from users where id = $1 and deleted_at is null", id)
	if err != nil {
		return nil, err
	}

	fmt.Println(user)
	return &user, nil
}

func (r *Repository) CreateUser(user modules.User) (int, error) {
	var id int

	query := "insert into users (name, age, gender, city) values ($1, $2, $3, $4) returning id"

	err := r.db.DB.QueryRow(query, user.Name, user.Age, user.Gender, user.City).Scan(&id) // &user.ID неправильно

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Repository) UpdateUser(user modules.User) error {
	query := "update users set name=$1, age=$2, gender=$3, city=$4 where id = $5"

	result, err := r.db.DB.Exec(query, user.Name, user.Age, user.Gender, user.City, user.ID)

	if err != nil {
		fmt.Printf("Error updating user: %v\n", err)
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Error updating user: %v\n", err)
	}
	if rows == 0 {
		fmt.Println("No rows updated.")
	}

	return nil
}

func (r *Repository) DeleteUser(id int) (int64, error) {
	//query := "delete from users where id = $1"
	query := "update users set deleted_at = NOW() where id = $1 and deleted_at is null"

	result, err := r.db.DB.Exec(query, id)
	if err != nil {
		return 0, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	if rows == 0 {
		return 0, fmt.Errorf("user not found or already deleted")
	}
	return rows, nil
}
