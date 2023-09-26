package config

import "time"

type (
	Env struct {
		App        App        `mapstructure:"APP"`
		Database   Database   `mapstructure:"DATABASE"`
		Repository Repository `mapstructure:"REPOSITORY"`
		Secret     Secret     `mapstructure:"SECRET"`
	}

	App struct {
		Server   Server `mapstructure:"SERVER"`
		Login    Login  `mapstructure:"LOGIN"`
		BypassDB bool   `mapstructure:"BYPASS_DB"`
	}

	Server struct {
		Address   string    `mapstructure:"ADDRESS"`
		Port      string    `mapstructure:"PORT"`
		RateLimit RateLimit `mapstructure:"RATE_LIMIT"`
	}

	RateLimit struct {
		Enabled  bool          `mapstructure:"ENABLED"`
		Max      int           `mapstructure:"MAX"`
		Duration time.Duration `mapstructure:"DURATION"`
	}

	Login struct {
		Timeout time.Duration `mapstructure:"TIMEOUT"`
	}

	Database struct {
		Mongo    Mongo    `mapstructure:"MONGO"`
		Postgres Postgres `mapstructure:"POSTGRES"`
		Redis    Redis    `mapstructure:"REDIS"`
	}

	Mongo struct {
		Host     string `mapstructure:"HOST"`
		Port     int    `mapstructure:"PORT"`
		User     string `mapstructure:"USER"`
		Password string `mapstructure:"PASSWORD"`
		Name     string `mapstructure:"NAME"`
	}

	Postgres struct {
		Host     string `mapstructure:"HOST"`
		Port     int    `mapstructure:"PORT"`
		User     string `mapstructure:"USER"`
		Password string `mapstructure:"PASSWORD"`
		Name     string `mapstructure:"NAME"`
	}

	Redis struct {
		Host     string `mapstructure:"HOST"`
		Port     int    `mapstructure:"PORT"`
		Password string `mapstructure:"PASSWORD"`
	}

	Repository struct {
		User RepositoryStore `mapstructure:"USER"`
	}

	RepositoryStore struct {
		Store string `mapstructure:"STORE"`
	}

	Secret struct {
		JwtSecret JwtSecret `mapstructure:"JWT"`
		OAuth2    OAuth2    `mapstructure:"OAUTH2"`
	}

	JwtSecret struct {
		Access  JwtAccessTokenSecret  `mapstructure:"ACCESS"`
		Refresh JwtRefreshTokenSecret `mapstructure:"REFRESH"`
	}

	JwtAccessTokenSecret struct {
		Key        string        `mapstructure:"KEY"`
		Expiration time.Duration `mapstructure:"EXPIRATION"`
	}

	JwtRefreshTokenSecret struct {
		Key        string        `mapstructure:"KEY"`
		Expiration time.Duration `mapstructure:"EXPIRATION"`
	}

	OAuth2 struct {
		Provider OAuth2Provider `mapstructure:"PROVIDER"`
	}

	OAuth2Provider struct {
		Google OAuth2Google `mapstructure:"GOOGLE"`
	}

	OAuth2Google struct {
		ClientID     string `mapstructure:"CLIENT_ID"`
		ClientSecret string `mapstructure:"CLIENT_SECRET"`
	}
)

var defaultConfig = Env{
	Database: Database{
		Mongo: Mongo{
			Host:     "default",
			Port:     2,
			User:     "default",
			Password: "default",
			Name:     "default",
		},
		Postgres: Postgres{
			Host:     "default",
			Port:     2,
			User:     "default",
			Password: "default",
			Name:     "default",
		},
	},
}
