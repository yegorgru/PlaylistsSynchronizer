package repositories

import (
	"PlaylistsSynchronizer/pkg/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type GroupPostgres struct {
	db *sqlx.DB
}

func NewGroupPostgres(db *sqlx.DB) *GroupPostgres {
	return &GroupPostgres{db: db}
}

func (r *GroupPostgres) Create(userID, roleID int, group models.Group) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	createGroup := fmt.Sprintf("INSERT INTO %s (name, description) values ($1, $2) RETURNING id", groupsTable)
	row1 := r.db.QueryRow(createGroup, group.Name, group.Description)
	if err := row1.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	var userGroupId int
	userGroup := fmt.Sprintf("INSERT INTO %s (userId, id, roleID) values ($1, $2, $3) RETURNING id", userGroupTable)
	row3 := r.db.QueryRow(userGroup, userID, id, roleID)
	if err := row3.Scan(&userGroupId); err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}

func (r *GroupPostgres) GetAll() ([]models.Group, error) {
	var groups []models.Group
	query := fmt.Sprintf("SELECT * FROM %s", groupsTable)
	err := r.db.Select(&groups, query)
	return groups, err
}

func (r *GroupPostgres) GetById(id int) (models.Group, error) {
	var group models.Group
	query := fmt.Sprintf("SELECT * FROM %s WHERE id= $1", groupsTable)
	err := r.db.Get(&group, query, id)
	return group, err
}

func (r *GroupPostgres) Update(id int, group models.UpdateGroupInput) error {
	query := fmt.Sprintf("UPDATE %s SET name=$1, description=$2 WHERE id=$3", groupsTable)
	_, err := r.db.Exec(query, group.Name, group.Description, id)
	return err
}

func (r *GroupPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id= $1", groupsTable)
	_, err := r.db.Exec(query, id)
	return err
}
