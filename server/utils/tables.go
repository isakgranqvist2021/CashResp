package utils

func CreateTablesIfNotExists() {
	db := Connect()
	defer db.Close()

	querys := []string{
		`
			CREATE TABLE IF NOT EXISTS cashresp.users (
			ID INT NOT NULL AUTO_INCREMENT,
			Email VARCHAR(499) NOT NULL,
			Password VARCHAR(499) NOT NULL,
			AuthType VARCHAR(45) NOT NULL,
			VerifyCode VARCHAR(49) NULL,
			CreatedAt DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			UpdatedAt DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			EmailVerified TINYINT NOT NULL DEFAULT 0,
			Admin TINYINT NOT NULL DEFAULT 0,
			PRIMARY KEY (ID),
			UNIQUE INDEX Email_UNIQUE (Email ASC) VISIBLE,
			UNIQUE INDEX ID_UNIQUE (ID ASC) VISIBLE);
		`,
		`
			CREATE TABLE IF NOT EXISTS cashresp.images (
			ID INT NOT NULL AUTO_INCREMENT,
			filename VARCHAR(200) NOT NULL,
			originalFilename VARCHAR(200) NOT NULL,
			type VARCHAR(45) NOT NULL,
			alt VARCHAR(45) NOT NULL,
			CreatedAt DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			UpdatedAt DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			size INT NOT NULL,
			PRIMARY KEY (ID));
		`,
		`
			CREATE TABLE IF NOT EXISTS cashresp.offers (
			ID int NOT NULL AUTO_INCREMENT,
			OfferID varchar(45) NOT NULL,
			PublisherID varchar(45) NOT NULL,
			AppID varchar(45) NOT NULL,
			Description varchar(999) NOT NULL,
			CreatedAt datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
			UpdatedAt datetime DEFAULT CURRENT_TIMESTAMP,
			Href varchar(45) NOT NULL,
			ImageID int NOT NULL,
			Draft TINYINT NOT NULL DEFAULT 0,
			Provider varchar(45) NOT NULL,
			PRIMARY KEY (ID),
			UNIQUE KEY ID_UNIQUE (ID),
			FOREIGN KEY (ImageID) REFERENCES images(ID));
		`,
	}

	for i := 0; i < len(querys); i++ {
		_, err := db.Exec(querys[i])

		if err != nil {
			panic(err)
		}
	}
}
