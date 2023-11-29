package repositories

import (
	"PlaylistsSynchronizer/pkg/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type TrackPostgres struct {
	db *sqlx.DB
}

func NewTrackPostgres(db *sqlx.DB) *TrackPostgres {
	return &TrackPostgres{db: db}
}

func (t *TrackPostgres) Create(playListId int, track models.Track) (int, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return 0, err
	}
	var ID int
	query := fmt.Sprintf("INSERT INTO %s (spotifyuri, youtubemusicid) values ($1, $2) RETURNING id", tracksTable)
	row1 := t.db.QueryRow(query, track.SpotifyUri, track.YouTubeMusicID)
	if err := row1.Scan(&ID); err != nil {
		tx.Rollback()
		return 0, err
	}
	var playListTrackID int
	playListTrack := fmt.Sprintf("INSERT INTO %s (trackID, playListID) values ($1, $2) RETURNING ID", playlistTrackTable)
	row2 := t.db.QueryRow(playListTrack, ID, playListId)
	if err := row2.Scan(&playListTrackID); err != nil {
		tx.Rollback()
		return 0, err
	}
	return ID, nil
}

func (t *TrackPostgres) GetAll() ([]models.Track, error) {
	var tracks []models.Track
	query := fmt.Sprintf("SELECT * FROM %s", tracksTable)
	err := t.db.Select(&tracks, query)
	return tracks, err
}

func (t *TrackPostgres) GetByPlayListID(playListID int) ([]models.Track, error) {
	var tracks []models.Track
	query := fmt.Sprintf("SELECT t.id, t.spotifyUri, t.youTubeMusicID "+
		"FROM %s t INNER JOIN %s pt ON t.id = pt.trackID WHERE pt.playListID=$1", tracksTable, playlistTrackTable)
	err := t.db.Select(&tracks, query, playListID)
	return tracks, err
}

func (t *TrackPostgres) GetByPlayListIDAndTrackApiID(playListID int, apiID models.ApiTrackID) (*models.Track, error) {
	var track models.Track
	query := fmt.Sprintf("SELECT t.id, t.spotifyUri, t.youTubeMusicID "+
		"FROM %s t INNER JOIN %s pt ON t.id = pt.trackID AND t.spotifyUri=$1 AND t.youTubeMusicID=$2 "+
		"WHERE pt.playListID=$3", tracksTable, playlistTrackTable)
	err := t.db.Get(&track, query, apiID.SpotifyUri, apiID.YouTubeMusic, playListID)
	if track == (models.Track{}) {
		return nil, nil
	}
	return &track, err
}

func (t *TrackPostgres) GetByTrackApiID(apiID models.ApiTrackID) (*models.Track, error) {
	var track models.Track
	query := fmt.Sprintf("SELECT t.id, t.spotifyUri, t.youTubeMusicID "+
		"FROM %s t WHERE t.spotifyUri=$1 AND t.youTubeMusicID=$2", tracksTable)
	err := t.db.Get(&track, query, apiID.SpotifyUri, apiID.YouTubeMusic)
	if track == (models.Track{}) {
		return nil, nil
	}
	return &track, err
}

func (t *TrackPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", tracksTable)
	_, err := t.db.Exec(query, id)
	return err
}

func (t *TrackPostgres) DeleteFromPlayList(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", playlistTrackTable)
	_, err := t.db.Exec(query, id)
	return err
}
