package migrations

var migartion_v1 []string = []string{
	`CREATE TABLE IF NOT EXISTS users (
		id VARCHAR(40) PRIMARY KEY,
        username VARCHAR(100),
		email VARCHAR(100),
		password VARCHAR(100)
	)`,
}
