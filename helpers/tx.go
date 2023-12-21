package helpers

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil{
		errRollBack := tx.Rollback()
		PanicIfError(errRollBack)
	}else{
		errCommit := tx.Commit()
		PanicIfError(errCommit)
	}
}

