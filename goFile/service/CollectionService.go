package service

import (
	"database/sql"
	"log"
)

type CollectionService struct{}

type ConnInfo struct {
	id        int
	conn_name string
	ip        string
	port      int
	userName  string
	password  string
	key       string
}

func QueryConnInfo(s *CollectionService) ([]*ConnInfo, error) {
	var err error
	db, err := sql.Open("sqlite3", "./literm")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	sqlQuery := `
select * from conn_info;
`
	rows, err := db.Query(sqlQuery)
	defer rows.Close()

	var connInfos []*ConnInfo

	for rows.Next() {
		var info *ConnInfo
		err = rows.Scan(&info.id, &info.conn_name, &info.ip, &info.port, &info.userName, &info.password, &info.key)
		if err != nil {
			log.Println(err)
		}
		connInfos = append(connInfos, info)
	}
	log.Println(connInfos)
	return connInfos, err
}
