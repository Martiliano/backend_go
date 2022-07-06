package config

//
// microservices => config => config.go
//

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/viper"
	"go.uber.org/zap"

	logger "BackEnd_Api/logger"
)

type FileInformation struct {
	Path string
	Name string
}

type Config struct {
	Grpc               GrpcServerConfig        `mapstructure:"GRPC"`
	Graphql            GraphqlServerConfig     `mapstructure:"GRAPHQL"`
	Rest               RestServerConfig        `mapstructure:"REST"`
	HealthCheck        HealthCheckServerConfig `mapstructure:"HEALTH_CHECK"`
	Logger             LoggerConfig            `mapstructure:"LOGGER"`
	Promthesius        PromthesiusConfig       `mapstructure:"PROMTHESIUS"`
	Mongo              MongoConfig             `mapstructure:"MONGO"`
	Postgres           PostgresConfig          `mapstructure:"POSTGRES"`
	Metrics            MetricsConfig           `mapstructure:"METRICS"`
	Jaeger             JaegerServerConfig      `mapstructure:"JAEGER"`
	Auth0              Auth0Config             `mapstructure:"AUTH0"`
	Auth               AuthConfig              `mapstructure:"AUTH"`
	ServerCertificate  ServerCertificateConfig `mapstructure:"SERVER_CERTIFICATE"`
	ClientCertificate  ClientCertificateConfig `mapstructure:"CLIENT_CERTIFICATE"`
	Email              EmailConfig             `mapstructure:"EMAIL"`
	GithubVCSConfig    VCSSConfig              `mapstructure:"GITHUB_VCS_CONFIG"`
	CadenceConfig      CadenceConfig           `mapstructure:"CADENCE_CONFIG"`
	SupportedVcsConfig []string
}

type EmailConfig struct {
	SmtpServer  string `mapstructure:"SMTP_SERVER"`
	SmtpSecure  string `mapstructure:"SMTP_SECURE"`
	SmtpPort    string `mapstructure:"SMTP_PORT"`
	PopServer   string `mapstructure:"POP_SERVER"`
	PopSecure   string `mapstructure:"POP_SECURE"`
	PopPort     string `mapstructure:"POP_PORT"`
	ImapServer  string `mapstructure:"IMAP_SERVER"`
	ImapSecure  string `mapstructure:"IMAP_SECURE"`
	ImapPort    string `mapstructure:"IMAP_PORT"`
	User        string `mapstructure:"USER"`
	Password    string `mapstructure:"PASSWORD"`
	Account     string `mapstructure:"ACCOUNT"`
	AccountName string `mapstructure:"ACCOUNT_NAME"`
}

type ServerCertificateConfig struct {
	ServerCertFile string `mapstructure:"SERVER_CERT_FILE"`
	ServerKeyFile  string `mapstructure:"SERVER_KEY_FILE"`
	CaCertFile     string `mapstructure:"CA_CERT_FILE"`
}

type ClientCertificateConfig struct {
	ClientCertFile string `mapstructure:"CLIENT_CERT_FILE"`
	ClientKeyFile  string `mapstructure:"CLIENT_KEY_FILE"`
	CaCertFile     string `mapstructure:"CA_CERT_FILE"`
}

type AuthConfig struct {
	Secret            string `mapstructure:"SECRET"`
	DurationInMinutes string `mapstructure:"DURATION_IN_MINUTES"`
	CreateSecret      string `mapstructure:"CREATE_SECRET"`
}

type PromthesiusConfig struct {
	Port string `mapstructure:"PORT"`
	Host string `mapstructure:"HOST"`
}

type CadenceConfig struct {
	Domain  string `mapstructure:"DOMAIN"`
	Service string `mapstructure:"SERVICE"`
	Port    string `mapstructure:"PORT"`
	Host    string `mapstructure:"HOST"`
}

type GrpcServerConfig struct {
	Port           string `mapstructure:"PORT"`
	Host           string `mapstructure:"HOST"`
	RequestTimeout int    `mapstructure:"REQUEST_TIMEOUT"`
}

type GraphqlServerConfig struct {
	Port           string `mapstructure:"PORT"`
	Host           string `mapstructure:"HOST"`
	RequestTimeout int    `mapstructure:"REQUEST_TIMEOUT"`
}

type RestServerConfig struct {
	Port           string `mapstructure:"PORT"`
	Host           string `mapstructure:"HOST"`
	RequestTimeout int    `mapstructure:"REQUEST_TIMEOUT"`
}

type HealthCheckServerConfig struct {
	Port string `mapstructure:"PORT"`
	Host string `mapstructure:"HOST"`
}

type LoggerConfig struct {
	LogLevel string `mapstructure:"LOG_LEVEL"`
}

type MongoConfig struct {
	Host   string `mapstructure:"HOST"`
	Port   string `mapstructure:"PORT"`
	DbName string `mapstructure:"DB_NAME"`
}

type PostgresConfig struct {
	Host     string `mapstructure:"HOST"`
	Port     string `mapstructure:"PORT"`
	User     string `mapstructure:"USER"`
	Password string `mapstructure:"PASSWORD"`
	DbName   string `mapstructure:"DB_NAME"`
	SslMode  string `mapstructure:"SSL_MODE"`
	Driver   string `mapstructure:"PG_DRIVER"`
}

type MetricsConfig struct {
	URL         string `mapstructure:"URL"`
	ServiceName string `mapstructure:"SERVICE_NAME"`
}

type JaegerServerConfig struct {
	Host        string `mapstructure:"HOST"`
	Port        string `mapstructure:"PORT"`
	ServiceName string `mapstructure:"SERVICE_NAME"`
	LogSpans    string `mapstructure:"LOG_SPANS"`
}

type Auth0Config struct {
	Auth0ClientID     string `mapstructure:"AUTH0_CLIENT_ID"`
	Auth0Domain       string `mapstructure:"AUTH0_DOMAIN"`
	Auth0ClientSecret string `mapstructure:"AUTH0_CLIENT_SECRET"`
	Auth0CallBackURL  string `mapstructure:"AUTH0_CALL_BACK_URL"`
}

type VCSSConfig struct {
	IType        int    `mapstructure:"ITYPE"`
	Provider     string `mapstructure:"PROVIDER"`
	URLTemplate  string `mapstructure:"URL_TEMPLATE"`
	ClientID     string `mapstructure:"CLIENT_ID"`
	RedirectURI  string `mapstructure:"REDIRECT_URI"`
	State        string `mapstructure:"STATE"`
	Scope        string `mapstructure:"SCOPE"`
	ResponseType string `mapstructure:"RESPONSE_TYPE"`
	ClientSecret string `mapstructure:"CLIENT_SECRET"`
	Name         string `mapstructure:"NAME"`
}

func GetVcsConfig(name string, vcsConfig []VCSSConfig) *VCSSConfig {
	for _, v := range vcsConfig {
		if v.Name == name {
			return &v
		}
	}
	return nil
}

func LoadConfig(filename, path string) (*viper.Viper, error) {
	v := viper.New()
	v.AddConfigPath(path)
	v.SetConfigName(filename)
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config
	err := v.Unmarshal(&c)
	if err != nil {
		logger.Log.Fatal("incapaz de decodificar em struct", zap.Error(err))
		return nil, err
	}
	return &c, nil
}

func GetConfigName() string {
	fileName := os.Getenv("CONFIG_NAME")
	if fileName != "" {
		return fileName
	}
	return "config"
}

func GetConfigDirectory() string {
	filePath := os.Getenv("CONFIG_DIRECTORY")

	if filePath != "" {
		return filePath
	}
	return RootDir()
}
func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

func GetConfig() *Config {
	configFileName := GetConfigName()
	configFileDirectory := GetConfigDirectory()

	logger.Log.Info("Detalhes da configuração", zap.String("configFileDirectory", configFileDirectory), zap.String("configFileName", configFileName))

	cfgFile, configFileLoadError := LoadConfig(configFileName, configFileDirectory)
	if configFileLoadError != nil {
		logger.Log.Fatal("não foi possível obter a configuração", zap.Error(configFileLoadError))
		panic(configFileLoadError)
	}

	cfg, parseError := ParseConfig(cfgFile)
	if parseError != nil {
		logger.Log.Fatal("não foi possível obter a configuração", zap.Error(parseError))
		panic(parseError)
	}

	cfg.SupportedVcsConfig = supportedVcsConfig()
	return cfg
}

func supportedVcsConfig() []string {
	return []string{"github"}
}
