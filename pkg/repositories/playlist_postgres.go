package repositories

import (
	"PlaylistsSynchronizer/pkg/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type PlayListPostgres struct {
	db *sqlx.DB
}

func NewPlayListPostgres(db *sqlx.DB) *PlayListPostgres {
	return &PlayListPostgres{db: db}
}

func (r *PlayListPostgres) Create(playlist models.PlayList) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, description, groupID) values ($1, $2, $3) RETURNING id", playlistsTable)
	row := r.db.QueryRow(query, playlist.Name, playlist.Description, playlist.GroupID)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *PlayListPostgres) GetAll() ([]models.PlayList, error) {
	var playLists []models.PlayList
	query := fmt.Sprintf("SELECT * FROM %s", playlistsTable)
	err := r.db.Select(&playLists, query)
	return playLists, err
}

func (r *PlayListPostgres) GetById(id int) (models.PlayList, error) {
	var playList models.PlayList
	query := fmt.Sprintf("SELECT * FROM %s WHERE id= $1", playlistsTable)
	err := r.db.Get(&playList, query, id)
	return playList, err
}

func (r *PlayListPostgres) GetByGroupId(id int) (models.PlayList, error) {
	var playList models.PlayList
	query := fmt.Sprintf("SELECT * FROM %s WHERE groupID=$1", playlistsTable)
	err := r.db.Get(&playList, query, id)
	return playList, err
}

func (r *PlayListPostgres) Update(id int, playList models.UpdatePlayListInput) error {
	query := fmt.Sprintf("UPDATE %s SET name=$1, description=$2, groupID=$3 WHERE id=$4", playlistsTable)
	_, err := r.db.Exec(query, playList.Name, playList.Description, playList.GroupID, id)
	return err
}

func (r *PlayListPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id= $1", playlistsTable)
	_, err := r.db.Exec(query, id)
	return err
}
