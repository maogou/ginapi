package setting

import "time"

//ServerSettingS 服务
type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

//AppSettingS 应用
type AppSettingS struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
}

//DatabaseSettingS 数据库
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

//JwtAuthSettings jwt
type JwtAuthSettings struct {
	Issuer string
	Secret string
	Expire time.Duration
}

//ZapLogSettings Zap-Log
type ZapLogSettings struct {
	Development     bool
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
	Level           string
	Format          string
	Prefix          string
	Director        string
	LinkName        string
	ShowLine        bool
	EncodeLevel     string
	StacktraceKey   string
	LogInConsole    bool
	MaxSize         int
	MaxAge          int
	MaxBackups      int
}

//ReadSection 读取对应节点的配置
func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)

	if err != nil {
		return err
	}

	return nil
}
