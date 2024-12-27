package service

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"literm/goFile/types"
	"os"
)

type CollectionService struct {
}

type ConnCollection struct {
	ID             int16  `db:"id"`
	CollectionName string `db:"collection_name"`
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

func (s *CollectionService) QueryCollections() []*ConnCollection {
	var err error
	db, err := sqlx.Connect("sqlite3", "./literm")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	var collections []ConnCollection
	sqlQuery := `select * from conn_collection;`
	err = db.Select(&collections, sqlQuery)
	if err != nil {
		log.Println(err)
	}
	log.Println("查询结果：", collections)
	collectionPointer := make([]*ConnCollection, len(collections))
	for i := range collections {
		collectionPointer[i] = &collections[i]
	}
	log.Println("查询结果：", collectionPointer)
	return collectionPointer
}

func (s *CollectionService) QueryConnInfos(collectionId int16) []*ConnInfo {
	var err error
	db, err := sqlx.Connect("sqlite3", "./literm")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	var connInfos []ConnInfo
	sqlQuery := `select * from conn_info where collection_id = ?;`
	err = db.Select(&connInfos, sqlQuery, collectionId)
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

func (s *CollectionService) InsertCollection(collectionName string) *types.Response {
	var err error
	db, err := sqlx.Connect("sqlite3", "./literm")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	var ids []int16
	sqlQuery := `select id from conn_collection order by id desc limit 1;`
	err = db.Select(&ids, sqlQuery)
	if err != nil {
		log.Println(err)
	}
	var id int16
	if ids != nil {
		id = ids[0]
	}
	log.Println("最大id为：", id)
	id++

	// 插入数据
	sqlInsert := `insert into conn_collection(collection_name) values (?);`
	_, err = db.Exec(sqlInsert, collectionName)
	if err != nil {
		log.Println(err)
	}
	return &types.Response{
		ResponseCode: 200,
		ResponseMsg:  "插入成功",
	}
}
