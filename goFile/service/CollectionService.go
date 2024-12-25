package service

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"os"
)

type CollectionService struct{}

type ConnCollection struct {
	ID             int16  `db:"id"`
	CollectionName string `db:"collection"`
}

type ConnInfo struct {
	ID            int16          `db:"id"`
	Collection_ID sql.NullInt16  `db:"collection_id"`
	ConnName      string         `db:"conn_name"`
	IP            string         `db:"conn_ip"`
	Port          int16          `db:"conn_port"`
	UserName      sql.NullString `db:"user_name"`
	Password      sql.NullString `db:"password"`
	Key           sql.NullString `db:"key"`
}

func init() {
	// 设置日志格式为JSON（可选）
	log.SetFormatter(&log.JSONFormatter{})

	// 设置日志级别（例如：debug, info, warn, error, fatal, panic）
	log.SetLevel(log.InfoLevel)

	// 输出到标准输出（默认），也可以设置为文件等其他输出
	log.SetOutput(os.Stdout)
}

func (s *CollectionService) QueryConnInfos() []*ConnInfo {
	var err error
	db, err := sqlx.Connect("sqlite3", "./literm")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	var connInfos []ConnInfo
	sqlQuery := `select * from conn_info;`
	err = db.Select(&connInfos, sqlQuery)
	if err != nil {
		log.Println(err)
	}
	log.Println("查询结果：", connInfos)
	// 如果你需要将 connInfos 中的元素转换为指针
	connInfoPointers := make([]*ConnInfo, len(connInfos))
	for i := range connInfos {
		connInfoPointers[i] = &connInfos[i]
	}
	log.Println("查询结果：", connInfoPointers)
	return connInfoPointers
}
