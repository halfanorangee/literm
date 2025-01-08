package service

import (
	"database/sql"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"literm/goFile/types"
	"os"
)

type CollectionService struct {
	db *gorm.DB
}

type Collection struct {
	gorm.Model
	ID             int64  `db:"id"`
	CollectionName string `db:"collection_name"`
}
type ConnInfo struct {
	gorm.Model
	ID            int64  `db:"id"`
	Collection_ID int64  `db:"collection_id"`
	ConnName      string `db:"conn_name"`
	IP            string `db:"conn_ip"`
	Port          string `db:"conn_port"`
	UserName      string `db:"user_name"`
	Password      string `db:"password"`
	Key           string `db:"key"`
}

type CollectionConnRel struct {
	ID             int64  `db:"id"`
	CollectionName string `db:"collection_name"`
	ConnInfos      []*ConnInfo
}
type CollectionAndConnInfo struct {
	CollectionName string
	ConnId         sql.NullInt64 `gorm:"column:id"`
	CollectionId   int64
	ConnName       sql.NullString
	ConnIP         sql.NullString `gorm:"column:"`
	ConnPort       sql.NullString
	UserName       sql.NullString
	Password       sql.NullString
	Key            sql.NullString
}

func init() {
	// 设置日志格式为JSON（可选）
	log.SetFormatter(&log.JSONFormatter{})

	// 设置日志级别（例如：debug, info, warn, error, fatal, panic）
	log.SetLevel(log.InfoLevel)

	// 输出到标准输出（默认），也可以设置为文件等其他输出
	log.SetOutput(os.Stdout)
}

// 查询所有连接集合及对应的连接信息
func (s *CollectionService) QueryAllConnectionRel() []*CollectionConnRel {
	var err error
	s.db, err = gorm.Open(sqlite.Open("./literm"), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	var connAllInfos []CollectionAndConnInfo
	s.db.Debug().Table("conn_collection").Joins("left join conn_info on conn_collection.id = conn_info.collection_id").
		Select("conn_collection.collection_name,conn_info.id,conn_collection.id as collection_id,conn_info.conn_name,conn_info.conn_ip,conn_info.conn_port,conn_info.user_name,conn_info.password,conn_info.key").
		Unscoped().
		Scan(&connAllInfos)
	if err != nil {
		log.Println(err)
	}

	//数据库查询结果转成map，map转成一对多的对象传递给前端
	connMap := make(map[int64][]CollectionAndConnInfo)
	for _, connInfo := range connAllInfos {
		connMap[connInfo.CollectionId] = append(connMap[connInfo.CollectionId], connInfo)
	}
	rels := make([]*CollectionConnRel, 0)
	for _, connInfos := range connMap {
		rel := &CollectionConnRel{
			ID:             connInfos[0].CollectionId,
			CollectionName: connInfos[0].CollectionName,
			ConnInfos:      make([]*ConnInfo, 0),
		}

		for _, connInfo := range connInfos {
			if connInfo.ConnId.Valid {
				info := &ConnInfo{
					ID:       connInfo.ConnId.Int64,
					ConnName: connInfo.ConnName.String,
					IP:       connInfo.ConnIP.String,
					Port:     connInfo.ConnPort.String,
					UserName: connInfo.UserName.String,
					Password: connInfo.Password.String,
					Key:      connInfo.Key.String,
				}
				rel.ConnInfos = append(rel.ConnInfos, info)
			}
		}
		rels = append(rels, rel)
	}
	return rels
}

// 插入新的连接集合
func (s *CollectionService) InsertCollection(collectionName string) *types.Response {
	var err error
	s.db, err = gorm.Open(sqlite.Open("./literm"), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}

	// 插入数据
	collection := Collection{CollectionName: collectionName}
	result := s.db.Create(&collection)
	if result.Error != nil {
		log.Println(result.Error)
		return &types.Response{
			ResponseCode: 500,
			ResponseMsg:  "插入失败",
		}
	}
	return &types.Response{
		ResponseCode: 200,
		ResponseMsg:  "插入成功",
	}
}

func (s *CollectionService) InsertConnection(info ConnInfo) *types.Response {
	var err error
	s.db, err = gorm.Open(sqlite.Open("./literm"), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	insertInfo := ConnInfo{}
	insertInfo.ConnName = info.ConnName
	insertInfo.IP = info.IP
	insertInfo.Port = info.Port
	insertInfo.UserName = info.UserName
	insertInfo.Password = info.Password
	insertInfo.Key = info.Key
	result := s.db.Create(&insertInfo)

	if result.Error != nil {
		log.Println(result.Error)
		return &types.Response{
			ResponseCode: 500,
			ResponseMsg:  "插入失败",
		}
	}
	return &types.Response{
		ResponseCode: 200,
		ResponseMsg:  "插入成功",
	}
}
