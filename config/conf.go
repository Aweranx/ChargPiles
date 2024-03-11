package config

import (
	"github.com/spf13/viper"
	"os"
)

var Config *Conf

type Conf struct {
	System *System `toml:"system"`
	//Oss           *Oss                    `toml:"oss"`
	MySql *MySql `toml:"mysql"`
	//Email         *Email                  `toml:"email"`
	Redis *Redis `toml:"redis"`
	//EncryptSecret *EncryptSecret          `toml:"encryptSecret"`
	//Cache         *Cache                  `toml:"cache"`
	//KafKa         map[string]*KafkaConfig `toml:"kafka"`
	//RabbitMq      *RabbitMq               `toml:"rabbitMq"`
	//Es            *Es                     `toml:"es"`
	PhotoPath *LocalPhotoPath `toml:"photoPath"`
	Sms       *Sms            `toml:"sms"`
}

type System struct {
	AppEnv      string `toml:"appEnv"`
	Domain      string `toml:"domain"`
	Version     string `toml:"version"`
	HttpPort    string `toml:"httpPort"`
	Host        string `toml:"host"`
	UploadModel string `toml:"uploadModel"`
}

type Oss struct {
	BucketName      string `toml:"bucketName"`
	AccessKeyId     string `toml:"accessKeyId"`
	AccessKeySecret string `toml:"accessKeySecret"`
	Endpoint        string `toml:"endPoint"`
	EndpointOut     string `toml:"endpointOut"`
	QiNiuServer     string `toml:"qiNiuServer"`
}

type MySql struct {
	Dialect  string `toml:"dialect"`
	DbHost   string `toml:"dbHost"`
	DbPort   string `toml:"dbPort"`
	DbName   string `toml:"dbName"`
	UserName string `toml:"userName"`
	Password string `toml:"password"`
	Charset  string `toml:"charset"`
}

type Email struct {
	ValidEmail string `toml:"validEmail"`
	SmtpHost   string `toml:"smtpHost"`
	SmtpEmail  string `toml:"smtpEmail"`
	SmtpPass   string `toml:"smtpPass"`
}

type Redis struct {
	RedisHost     string `toml:"redisHost"`
	RedisPort     string `toml:"redisPort"`
	RedisUsername string `toml:"redisUsername"`
	RedisPassword string `toml:"redisPwd"`
	RedisDbName   int    `toml:"redisDbName"`
	RedisNetwork  string `toml:"redisNetwork"`
	RedisEmpires  int    `toml:"redisEmpires"`
}

type KafkaConfig struct {
	DisableConsumer bool   `toml:"disableConsumer"`
	Debug           bool   `toml:"debug"`
	Address         string `toml:"address"`
	RequiredAck     int    `toml:"requiredAck"`
	ReadTimeout     int64  `toml:"readTimeout"`
	WriteTimeout    int64  `toml:"writeTimeout"`
	MaxOpenRequests int    `toml:"maxOpenRequests"`
	Partition       int    `toml:"partition"`
}

type EncryptSecret struct {
	JwtSecret   string `toml:"jwtSecret"`
	EmailSecret string `toml:"emailSecret"`
	PhoneSecret string `toml:"phoneSecret"`
	MoneySecret string `toml:"moneySecret"`
}

type LocalPhotoPath struct {
	PhotoHost   string `toml:"photoHost"`
	ProductPath string `toml:"productPath"`
	AvatarPath  string `toml:"avatarPath"`
}

//type Cache struct {
//	CacheType    string `toml:"cacheType"`
//	CacheExpires int64  `toml:"cacheExpires"`
//	CacheWarmUp  bool   `toml:"cacheWarmUp"`
//	CacheServer  string `toml:"cacheServer"`
//}

type Es struct {
	EsHost  string `toml:"esHost"`
	EsPort  string `toml:"esPort"`
	EsIndex string `toml:"esIndex"`
}

type RabbitMq struct {
	RabbitMQ         string `toml:"rabbitMq"`
	RabbitMQUser     string `toml:"rabbitMqUser"`
	RabbitMQPassWord string `toml:"rabbitMqPassWord"`
	RabbitMQHost     string `toml:"rabbitMqHost"`
	RabbitMQPort     string `toml:"rabbitMqPort"`
}

type Sms struct {
	EndPoint        string `toml:"endPoint"`
	AccessKeyId     string `toml:"accessKeyId"`
	AccessKeySercet string `toml:"accessKeySercet"`
}

func Init() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(workDir + "/config/")
	viper.AddConfigPath(workDir)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Config)
	if err != nil {
		panic(err)
	}
}
