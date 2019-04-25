package poststats

import (
	"database/sql"
	"errors"
	"log"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

var (
	errNotFound   = errors.New("post stats not found")
	errNotCreated = errors.New("post stats not created")
)

// PostStats describes post statistics
type PostStats struct {
	UID         uuid.UUID
	NumLikes    int32
	NumDislikes int32
	NumViews    int32
}

type datastore interface {
	get(uuid.UUID) (*PostStats, error)
	create(uuid.UUID) (*PostStats, error)
	like(uuid.UUID, uuid.UUID) (bool, bool, error)
	dislike(uuid.UUID, uuid.UUID) (bool, bool, error)
	view(uuid.UUID) error
	delete(uuid.UUID) error
}

type db struct {
	*sql.DB
}

func newDB(connString string) (*db, error) {
	postgres, err := sql.Open("postgres", connString)
	return &db{postgres}, err
}

func (db *db) get(uid uuid.UUID) (*PostStats, error) {
	query := `SELECT
	(SELECT num_views FROM post_views WHERE post_uid=$1),
	(SELECT COUNT(*) FROM post_votes WHERE post_uid=$1 AND vote = 1) "num_likes",
	(SELECT COUNT(*) FROM post_votes WHERE post_uid=$1 AND vote = -1) "num_dislikes"`
	row := db.QueryRow(query, uid.String())
	result := new(PostStats)
	switch err := row.Scan(&result.NumViews, &result.NumLikes, &result.NumDislikes); err {
	case nil:
		result.UID = uid
		return result, nil
	case sql.ErrNoRows:
		return nil, errNotFound
	default:
		return nil, err
	}
}

func (db *db) create(uid uuid.UUID) (*PostStats, error) {
	query := "INSERT INTO post_views (post_uid, num_views) VALUES ($1, 0)"
	ps := new(PostStats)
	ps.UID = uid
	ps.NumLikes = 0
	ps.NumDislikes = 0
	ps.NumViews = 0
	result, err := db.Exec(query, uid.String())
	nRows, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if nRows == 0 {
		return nil, errNotCreated
	}

	return ps, nil
}

func (db *db) like(postUID, userUID uuid.UUID) (bool, bool, error) {
	return db.updateVote(postUID, userUID, 1)
}

func (db *db) dislike(postUID, userUID uuid.UUID) (bool, bool, error) {
	return db.updateVote(postUID, userUID, -1)
}

func (db *db) updateVote(postUID, userUID uuid.UUID, vote int) (bool, bool, error) {
	query := "SELECT vote FROM post_votes WHERE post_uid=$1 AND user_uid=$2"
	row := db.QueryRow(query, postUID.String(), userUID.String())
	var userVote int
	switch err := row.Scan(&userVote); err {
	case nil:
		if vote == userVote {
			return false, false, nil
		}
		log.Println(vote, userVote)

		query = "UPDATE post_votes SET vote=$1 WHERE post_uid=$2 AND user_uid=$3"
		result, err := db.Exec(query, vote, postUID.String(), userUID.String())
		nRows, err := result.RowsAffected()
		if err != nil {
			return false, false, err
		}

		if nRows == 0 {
			return false, false, errNotFound
		}

		return true, false, nil
	case sql.ErrNoRows:
		query = "INSERT INTO post_votes(post_uid, user_uid, vote) VALUES ($1, $2, $3)"
		result, err := db.Exec(query, postUID.String(), userUID.String(), vote)
		nRows, err := result.RowsAffected()
		if err != nil {
			return false, false, err
		}

		if nRows == 0 {
			return false, false, errNotFound
		}

		return true, true, nil
	default:
		return false, false, err
	}
}

func (db *db) view(uid uuid.UUID) error {
	query := "UPDATE post_views SET num_views = num_views + 1 WHERE post_uid=$1"
	result, err := db.Exec(query, uid.String())
	nRows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if nRows == 0 {
		return errNotFound
	}

	return nil
}

func (db *db) delete(uid uuid.UUID) error {
	query := "DELETE FROM post_views WHERE post_uid=$1"
	result, err := db.Exec(query, uid.String())
	nRows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if nRows == 0 {
		return errNotFound
	} else {
		query := "DELETE FROM post_votes WHERE post_uid=$1"
		result, err := db.Exec(query, uid.String())
		_, err = result.RowsAffected()
		if err != nil {
			return err
		}
	}

	return nil
}
