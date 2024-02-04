package sqlexec

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func InitDB(DSN string) (DB *sql.DB, err error) {

	DB, err = sql.Open("postgres", DSN)
	if err != nil {
		return nil, err
	}

	info := fmt.Sprintf("dsn check success")
	log.Println(info)

	err = DB.Ping()
	if err != nil {
		return nil, err
	}

	info = fmt.Sprintf("database connect success")
	log.Println(info)

	return DB, nil
}

func GetHeightAndDelay(DB *sql.DB) (height int64, ts int64, err error) {
	SQL := fmt.Sprintf("SELECT height,timestamp FROM \"explorer\".\"block\" ORDER BY height DESC LIMIT 1 OFFSET 0")
	err = DB.QueryRow(SQL).Scan(&height, &ts)
	if err != nil {
		return 0, 0, err
	}

	return height, ts, nil
}
