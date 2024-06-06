package cfg

import (
	"os"
	"sync"
)

var lock = &sync.Mutex{}

const ATRIBUTO_COR = "Cor"
const ATRIBUTO_TAMANHO = "Tamanho"

type AppConfig struct {
	BaseUrl        string
	ImagePath      string
	ConsumerKey    string
	ConsumerSecret string
	WPKey          string
	WPSecret       string
	ManageStock    bool
	BackOrders     string
	PriceTable     string
	CodLoja        string
	FtpServer      string
	FtpUser        string
	FtpPassword    string
	FtpRemotePath  string
	FtpUrlBase     string
}

var singleAppConfig *AppConfig

func GetInstance() *AppConfig {
	if singleAppConfig == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleAppConfig == nil {
			singleAppConfig = &AppConfig{
				BaseUrl:        os.Getenv("BASE_URL"),
				ImagePath:      os.Getenv("IMAGE_PATH"),
				ConsumerKey:    os.Getenv("CONSUMER_KEY"),
				ConsumerSecret: os.Getenv("CONSUMER_SECRET"),
				WPKey:          os.Getenv("WP_KEY"),
				WPSecret:       os.Getenv("WP_SECRET"),
				ManageStock:    os.Getenv("WOO_MANAGE_STOCK") == "yes",
				BackOrders:     os.Getenv("WOO_BACKORDERS"),
				PriceTable:     os.Getenv("STX_PRICETABLE"),
				CodLoja:        os.Getenv("STX_CODLOJA"),
				FtpServer:      os.Getenv("FTP_SERVER"),
				FtpUser:        os.Getenv("FTP_USER"),
				FtpPassword:    os.Getenv("FTP_PASSWORD"),
				FtpRemotePath:  os.Getenv("FTP_REMOTEPATH"),
				FtpUrlBase:     os.Getenv("FTP_URLBASE"),
			}
		}
	}
	return singleAppConfig
}
