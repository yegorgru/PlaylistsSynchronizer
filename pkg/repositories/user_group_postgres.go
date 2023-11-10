package repositories

import (
	"PlaylistsSynchronizer/pkg/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserGroupPostgres struct {
	db *sqlx.DB
}

func NewUserGroupPostgres(db *sqlx.DB) *UserGroupPostgres {
	return &UserGroupPostgres{db: db}
}

func (r *UserGroupPostgres) Create(userGroup models.UserGroup) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (userid, groupid, roleid) values ($1, $2, $3) RETURNING id", userGroupTable)
	row := r.db.QueryRow(query, userGroup.UserID, userGroup.GroupID, userGroup.RoleID)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *UserGroupPostgres) GetAll() ([]models.UserGroup, error) {
	var groups []models.UserGroup
	query := fmt.Sprintf("SELECT * FROM %s", userGroupTable)
	err := r.db.Select(&groups, query)
	return groups, err
}

func (r *UserGroupPostgres) GetById(id int) (models.UserGroup, error) {
	var group models.UserGroup
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", userGroupTable)
	err := r.db.Get(&group, query, id)
	return group, err
}

func (r *UserGroupPostgres) Update(id int, group models.UpdateUserGroupInput) error {
	query := fmt.Sprintf("UPDATE %s SET userID=$1, groupID=$2, roleid=$3 WHERE id=$4", userGroupTable)
	_, err := r.db.Exec(query, group.UserID, group.GroupID, group.RoleID, id)
	return err
}

func (r *UserGroupPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", userGroupTable)
	_, err := r.db.Exec(query, id)
	return err
}
