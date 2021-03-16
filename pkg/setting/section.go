package setting

import "time"

//服务
type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

//应用
type AppSettingS struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
}

//数据库
type DatabaseSettingS struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

//JWT
type JwtAuthSettings struct {
	Issuer string
	Secret string
	Expire time.Duration
}

//Zap Log
type ZapLogSettings struct {
	Level         string
	Format        string
	Prefix        string
	Director      string
	LinkName      string
	ShowLine      bool
	EncodeLevel   string
	StacktraceKey string
	LogInConsole  bool
}

//读取对应节点的配置
func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)

	if err != nil {
		return err
	}

	return nil
}
