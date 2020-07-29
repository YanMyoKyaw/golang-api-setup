package database

func CreateProduct(p []string) error {
	db = db.Exec("INSERT INTO products VALUES(?, ?, ?, ?, ?)", p[0], p[1], p[2], p[3], p[4])
	if db.Error != nil {
		return db.Error
	}
	return nil
}
