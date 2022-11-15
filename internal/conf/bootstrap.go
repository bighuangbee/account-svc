package conf

type Bootstrap struct {
	Name         string
	Version      string
	Env          string
	Server       *Server
	Trace        *Trace
	Data         *Data
	Metrics      *Metrics
	Discovery    *Discovery
	MicroService *MicroService
	HttpEndpoint *HttpEndpoint
	MQTT         *MQTTServer
	SysConf map[string]map[string]interface{}
	OpLog   *OpLog
}

type Data struct {
	Database  *Database
	Redis     *Redis
	CloudCore *CloudCore
}

type Metrics struct {
	Addr string
}

type Database struct {
	Address, UserName, Password, DBName, Driver string
	Timeout                                     int
}

type Redis struct {
	Address, Password, AutoPrefix string
	DB                            int32
}

type Trace struct {
	Endpoint string
	Fraction float64
}

type Server struct {
	HTTP   *HTTPServer
	Grpc   *GRPCServer
	NodeId int64
}

type GRPCServer struct {
	Network string
	Addr    string
	Timeout int32
}

type HTTPServer struct {
	Addr    string
	Timeout int32
}

type Discovery struct {
	IPAddr  string `json:"ipAddr"`
	Port    uint32
	Timeout uint32
	OnOff   bool
}

type HttpEndpoint struct {
	Addr    string
	NeedAlt bool
}

type MQTTServer struct {
	Server                string
	SubTopicDeviceUpdate  string
	SubTopicChannelUpdate string
	PubTopoicChanCalib    string
}


type OpLog struct {
	// cloud core的地址
	OpLogGrpcAddr string
	// mozi Settings 服务的地址
	SettingsServiceHttpAddr string
}
