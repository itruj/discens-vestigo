package api

import (
	// "database/sql"
	"itruj/discens-vestigo/api/store/db"
	"net/http"
)

// type DB interface {
// 	Query(query string, args ...interface{}) (*sql.Rows, error)
// 	Close() error
// }
//
// type sqliteDB struct {
// 	*sql.DB
// }
//
// func NewSQLiteDB(dsn string) (DB, error) {
// 	db, err := sql.Open("sqlite3", dsn)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	if err := db.Ping(); err != nil {
// 		db.Close()
// 		return nil, err
// 	}
//
// 	return &sqliteDB{DB: db}, nil
// }
//
// func (d *sqliteDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
// 	return d.DB.Query(query, args...)
// }
//
// func (d *sqliteDB) Close() error {
// 	return d.DB.Close()
// }

type Server struct {
	DB db.DB
}

func NewServer(db db.DB) *Server {
	return &Server{DB: db}
}

func (s *Server) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleHomeView)
	mux.HandleFunc("/actions", s.handleActionsView)

	return mux
}
