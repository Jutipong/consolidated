package model

//##Server
type ServerConfiguration struct {
	Env       string
	RunMode   string
	Port      string
	SecretKey string
}

//##Database
type DatabaseConfiguration struct {
	Server       string
	Port         int
	User         string
	Pass         string
	DatabaseName string
}

//##Logger File
type LoggerFile struct {
	RootPath string
}
