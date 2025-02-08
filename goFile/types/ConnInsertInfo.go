package types

type ConnInsertInfo struct {
	CollectionId int64
	Name         string
	Comment      string
	IP           string
	Port         string
	UserName     string
	Password     string
	Authkey      string
	AuthType     string
}
