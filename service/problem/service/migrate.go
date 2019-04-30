package service

import (
	"context"
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

const sqlFresh = `
CREATE TABLE problems (id VARCHAR(255) PRIMARY KEY, title VARCHAR(255), statement BLOB);
`

func (s *serv) Migrate(prevVersion string) error {
	// Alter config to support multi statements
	mscfg, err := mysql.ParseDSN(s.config.MySQL)
	if err != nil {
		return err
	}
	mscfg.MultiStatements = true
	db, err := sql.Open("mysql", mscfg.FormatDSN())
	if err != nil {
		return err
	}
	defer db.Close()

	if prevVersion == "" {
		_, err := db.ExecContext(context.Background(), sqlFresh)
		return err
	}
	return nil
}