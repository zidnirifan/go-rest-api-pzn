package helper

import (
	"database/sql"
)

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errCallback := tx.Rollback()
		PanicIfError(errCallback)
	} else {
		errCommit := tx.Commit()
		PanicIfError(errCommit)
	}
}
