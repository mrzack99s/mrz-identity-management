package config

var (
	SystemConfigVar SystemConfig
)

type SystemConfig struct {
	Prod bool `env:"PROD,default=false"`
	Port int  `env:"PORT,default=8000"`
	DB   struct {
		Hostname string `env:"DB_HOSTNAME"`
		Username string `env:"DB_USRNAME"`
		Password string `env:"DB_PASSWORD"`
		DBName   string `env:"DB_NAME"`
	}
	LDAP struct {
		Hostname        string `env:"LDAP_HOSTNAME"`
		Username        string `env:"LDAP_USERNAME"`
		Password        string `env:"LDAP_PASSWORD"`
		BaseDN          string `env:"LDAP_BASE_DN"`
		DomainName      string `env:"LDAP_DOMAIN_NAME"`
		EmailDomainName string `env:"LDAP_EMAIL_DOMAIN_NAME"`
	}
	Security struct {
		Salt      string `env:"SECURE_SALT"`
		A2ASecret string `env:"SECURE_A2A_SECRET"`
	}
}
