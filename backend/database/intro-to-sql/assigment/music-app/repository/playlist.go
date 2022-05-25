package repository

import (
	"database/sql"

	"github.com/ruang-guru/playground/backend/database/intro-to-sql/assigment/music-app/model"
)

type PlaylistRepositoryInterface interface {
	FetchUserPlaylists(userID int64) ([]model.UserPlaylist, error)
	CreatePlaylist(playlist model.Playlist) error
	UpdateUserPlaylist(userID int64, playlist model.Playlist) error
	FetchPlaylistTrack(playlistID int64) ([]model.PlaylistTrack, error)
}

type PlaylistRepository struct {
	db *sql.DB
}

func NewPlaylistRepository(db *sql.DB) PlaylistRepositoryInterface {
	return &PlaylistRepository{db}
}

func (p *PlaylistRepository) FetchUserPlaylists(userID int64) ([]model.UserPlaylist, error) {
	// var sqlStatement string
	var playlists []model.UserPlaylist

	// Task 1: Buat query untuk mengambil playlist yang dimiliki user
	// TODO: answer here
	sqlStatement := `SELECT u.id as user_id, u.name as user_name, p.id as playlist_id,
	p.name as playlist_name FROM playlists p INNER JOIN users u ON u.id = p.user_id WHERE user_id = ?;`

	// Task 2: Buat execute statement untuk mengambil playlist yang dimiliki user
	// TODO: answer here
	rows, err := p.db.Query(sqlStatement, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// Task 3: Buat looping untuk mengambil data playlist yang dimiliki user
	// TODO: answer here
	for rows.Next() {
		var playlist model.UserPlaylist
		err := rows.Scan(&playlist.UserID, &playlist.UserName, &playlist.PlaylistID, &playlist.PlaylistName)
		if err != nil {
			return nil, err
		}
		playlists = append(playlists, playlist)
	}

	return playlists, nil
}

func (p *PlaylistRepository) CreatePlaylist(playlist model.Playlist) error {
	// var sqlStatement string

	// Task 1 : Buat query untuk menambahkan playlist baru
	// TODO: answer here
	sqlStatement := `INSERT INTO playlists (name, created_at, user_id) VALUES (?, ?, ?);`

	// Task 2 := Buat execute statement untuk menambahkan playlist baru
	// TODO: answer here
	_, err := p.db.Exec(sqlStatement, playlist.Name, playlist.CreatedAt, playlist.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (p *PlaylistRepository) UpdateUserPlaylist(userID int64, playlist model.Playlist) error {
	// var sqlStatement string

	// Task 1 : Buat query untuk mengupdate playlist name dengan id playlist tertentu yang dimiliki user tertentu
	// TODO: answer here
	sqlStatement := `UPDATE playlists SET name = ? WHERE id = ? AND user_id = ?;`

	// Task 2 : Buat execute statement untuk mengupdate playlist name dengan id playlist tertentu yang dimiliki user tertentu
	// TODO: answer here
	_, err := p.db.Exec(sqlStatement, playlist.Name, playlist.ID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (p *PlaylistRepository) FetchPlaylistTrack(playlistID int64) ([]model.PlaylistTrack, error) {
	// var sqlStatement string

	// Task 1 : Buat query untuk mengambil track yang dimiliki playlist tertentu
	// TODO: answer here
	sqlStatement := `SELECT pt.playlist_id, p.name as playlist_name, pt.track_id, t.name as track_name, t.artist as track_artist FROM playlist_tracks pt INNER JOIN playlists p ON p.id = pt.playlist_id
	INNER JOIN tracks t ON t.id = pt.track_id WHERE pt.playlist_id = ?`

	// Task 2 : Buat execute statement untuk mengambil track yang dimiliki playlist tertentu
	// TODO: answer here
	row, err := p.db.Query(sqlStatement, playlistID)
	if err != nil {
		return nil, err
	}

	defer row.Close()

	var playlistTracks []model.PlaylistTrack
	// Task 3 : Buat looping untuk mengambil track yang dimiliki playlist tertentu
	// TODO: answer here
	for row.Next() {
		var playlistTrack model.PlaylistTrack
		err := row.Scan(&playlistTrack.PlaylistID, &playlistTrack.PlaylistName, &playlistTrack.TrackID, &playlistTrack.TrackName, &playlistTrack.TrackArtist)
		if err != nil {
			return nil, err
		}
		playlistTracks = append(playlistTracks, playlistTrack)
	}

	return playlistTracks, nil
}
