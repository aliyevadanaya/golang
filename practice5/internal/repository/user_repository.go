package repository

import (
	"database/sql"
	"fmt"
	"practice5/internal/models"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) GetPaginatedUsers(page int, pageSize int, orderBy, name, email, gender string) (models.PaginatedResponse, error) {

	var users []models.User
	offset := (page - 1) * pageSize

	if orderBy == "" {
		orderBy = "id"
	}

	query := "select id, name, email, gender, birth_date from users where 1=1"
	countQuery := "select count(*) from users where 1=1"

	var args []interface{}
	argID := 1

	if name != "" {
		query += fmt.Sprintf(" and name ilike $%d", argID)
		countQuery += fmt.Sprintf(" and name ilike $%d", argID)
		args = append(args, "%"+name+"%")
		argID++
	}

	if email != "" {
		query += fmt.Sprintf(" and email ilike $%d", argID)
		countQuery += fmt.Sprintf(" and email ilike $%d", argID)
		args = append(args, "%"+email+"%")
		argID++
	}

	if gender != "" {
		query += fmt.Sprintf(" and gender = $%d", argID)
		countQuery += fmt.Sprintf(" and gender = $%d", argID)
		args = append(args, gender)
		argID++
	}

	query += fmt.Sprintf(" order by %s limit $%d offset $%d", orderBy, argID, argID+1)

	argsWithPagination := append(args, pageSize, offset)

	var totalCount int
	err := r.db.QueryRow(countQuery, args...).Scan(&totalCount)
	if err != nil {
		return models.PaginatedResponse{}, err
	}

	rows, err := r.db.Query(query, argsWithPagination...)
	if err != nil {
		return models.PaginatedResponse{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Gender, &u.BirthDate); err != nil {
			return models.PaginatedResponse{}, err
		}
		users = append(users, u)
	}

	return models.PaginatedResponse{
		Data:       users,
		TotalCount: totalCount,
		Page:       page,
		PageSize:   pageSize,
	}, nil
}

func (r *Repository) GetCommonFriends(user1 int, user2 int) ([]models.User, error) {

	var friends []models.User

	query := `select u.id, u.name, u.email, u.gender, u.birth_date from users u 
    join user_friends uf1 on u.id = uf1.friend_id join user_friends uf2 on u.id = uf2.friend_id where uf1.user_id = $1 and uf2.user_id = $2`

	rows, err := r.db.Query(query, user1, user2)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var u models.User

		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Gender, &u.BirthDate); err != nil {
			return nil, err
		}

		friends = append(friends, u)
	}

	return friends, nil
}
