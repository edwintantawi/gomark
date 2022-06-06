package helper

type sqlTX interface {
	Rollback() error
	Commit() error
}

func HandleDeferTX(tx sqlTX) {
	err := recover()

	if err != nil {
		errRollback := tx.Rollback()
		PanicError(errRollback)
	} else {
		errCommit := tx.Commit()
		PanicError(errCommit)
	}
}
