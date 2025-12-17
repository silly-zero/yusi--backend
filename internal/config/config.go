package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

	Mysql struct {
		DataSource string
	}

	Redis struct {
		Host string
		Type string
		Pass string
	}

	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

	AI struct {
		QwenApiKey  string
		MilvusUri   string
		MilvusToken string
	}

	Encryption struct {
		Key string
	}
}
