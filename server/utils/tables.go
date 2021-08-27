package utils

import "database/sql"

func CreateTablesIfNotExists() {
	db := Connect()
	defer db.Close()

	CreateUsersTable(db)
}

func CreateUsersTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS cashresp.users (
		ID INT NOT NULL AUTO_INCREMENT,
		Email VARCHAR(499) NOT NULL,
		Password VARCHAR(499) NOT NULL,
		AuthType VARCHAR(45) NOT NULL,
		VerifyCode VARCHAR(49) NULL,
		CreatedAt DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		UpdatedAt DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		EmailVerified TINYINT NOT NULL DEFAULT 0,
		PRIMARY KEY (ID),
		UNIQUE INDEX Email_UNIQUE (Email ASC) VISIBLE,
		UNIQUE INDEX ID_UNIQUE (ID ASC) VISIBLE);
	`

	_, err := db.Exec(query)

	if err != nil {
		panic(err)
	}

	return nil
}
