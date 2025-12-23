package helper

import "database/sql"

// membuat function untuk melakukan commit atau rollback transaction
func CommitOrRollback(tx *sql.Tx) {
	// melakukan recover error
	err := recover()

		// mengecek jika terjadi error
		if err != nil {
			// maka semua proses transaction di database akan dibatalkan
			errorRollback := tx.Rollback()

			// mengeceke error kembali pada saat rollback (overlapping)
			PanicIfError(errorRollback)

			// kemudian akan mengembalikan error untuk diterima di kode atas
			panic(err)
		} else {
			// tapi kalau misalnya tidak error, maka lakukan commit
			errorCommit := tx.Commit()

			// mengecek error kembali pada saat commit (overlapping)
			PanicIfError(errorCommit)
		}
}