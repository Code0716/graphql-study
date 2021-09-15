package util

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

// Environment has application env variables.
type Environment struct {
	AppPort     string
	AppHost     string
	AppEnv      string
	AppLocale   string
	AppLogLevel string
	AppDebug    bool

	DBDialect  string
	DBHost     string
	DBPort     string
	DBName     string
	DBAdmin    string
	DBPassword string
	DBCharset  string
	DBTimezone string

	Signingkey string
}

var env Environment
var envMux sync.Mutex
var envOnce sync.Once

// Env returns env variables.
func Env() Environment {
	envMux.Lock()
	defer envMux.Unlock()

	envOnce.Do(func() {
		env = loadEnv()
	})

	return env
}

func loadEnv() Environment {
	envPath := os.Getenv("GO_ENV")
	err := godotenv.Load(fmt.Sprintf("%s.env", envPath))
	if err != nil {
		panic(err)
	}
	var e Environment

	// APP side
	e.AppHost = os.Getenv("APP_HOST")
	e.AppPort = os.Getenv("APP_LISTEN_PORT")
	e.AppEnv = os.Getenv("APP_ENV")
	e.AppLocale = os.Getenv("APP_LOCALE")
	e.AppLogLevel = os.Getenv("APP_LOG_LEVEL")
	isDebug, err := strconv.ParseBool(os.Getenv("APP_DEBUG"))
	if err != nil {
		fmt.Println(err)
	}
	e.AppDebug = isDebug
	// DB side
	e.DBDialect = os.Getenv("DATABASE_DIALECT")
	e.DBHost = os.Getenv("DATABASE_MYSQL_HOST")
	e.DBPort = os.Getenv("DATABASE_MYSQL_PORT")
	e.DBName = os.Getenv("DATABASE_MYSQL_DBNAME")
	e.DBAdmin = os.Getenv("DATABASE_MYSQL_USER")
	e.DBPassword = os.Getenv("DATABASE_MYSQL_PASSWORD")
	e.DBCharset = os.Getenv("DATABASE_MYSQL_CHARSET")
	e.DBTimezone = os.Getenv("DATABASE_MYSQL_TIMEZONE")
	e.Signingkey = os.Getenv("SIGNINGKEY")
	return e
}

// GetAPIPath get path
// TODO:環境によって読み替えが必要だが、後ほど。
func GetAPIPath(e Environment) string {
	addr := fmt.Sprintf("%s:%s", e.AppHost, e.AppPort)
	return addr
}
