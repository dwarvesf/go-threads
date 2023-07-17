package threads

import (
	"errors"
	"fmt"

	"github.com/dwarvesf/go-threads/private"
	"github.com/dwarvesf/go-threads/util"
)

// Config config to init thread client
type Config struct {
	Username           string
	Password           string
	UserID             int
	APIToken           string
	TimezoneOffset     int
	DeviceID           string
	DeviceModel        string
	DeviceManufacturer string
	DeviceOsVersion    int
	DeviceOsRelease    string
}

func (c Config) ReadyCheck() error {
	if c.Username == "" || c.Password == "" {
		return errors.New("credential is required")
	}

	if c.UserID == 0 {
		return errors.New("user id is required")
	}

	if c.APIToken == "" {
		return errors.New("api token is empty")
	}
	return nil
}

// ConfigFn define the function to update config data
type ConfigFn func(Config) (Config, error)

// InitConfig init configs
func InitConfig(configs ...ConfigFn) (*Config, error) {
	configs = append([]ConfigFn{WithDefaultValue()}, configs...)

	c := Config{}
	for idx := range configs {
		fn := configs[idx]

		tmp, err := fn(c)
		if err != nil {
			return nil, err
		}
		c = tmp
	}

	return &c, nil
}

// WithDefaultValue initial default value
func WithDefaultValue() ConfigFn {
	return func(c Config) (Config, error) {
		c.TimezoneOffset = -14400
		c.DeviceID = util.GenerateAndroidDeviceID()
		c.DeviceManufacturer = "OnePlus"
		c.DeviceModel = "ONEPLUS+A3010"
		c.DeviceOsVersion = 25
		c.DeviceOsRelease = "7.1.1"

		return c, nil
	}
}

// WithCridential update the user cridential
func WithCridential(username string, password string) ConfigFn {
	return func(c Config) (Config, error) {
		if c.UserID == 0 {
			tmp, err := WithUserIDFetching(username)(c)
			if err != nil {
				return c, err
			}
			c = tmp
		}
		c.Username = username
		c.Password = password
		return c, nil
	}
}

// WithUserID update the user ID
func WithUserID(userID int) ConfigFn {
	return func(c Config) (Config, error) {
		c.UserID = userID
		return c, nil
	}
}

// WithDeviceID update the device ID
func WithDeviceID(deviceID string) ConfigFn {
	return func(c Config) (Config, error) {
		c.DeviceID = deviceID
		return c, nil
	}
}

// WithUserIDFetching fetch user id from instagram
func WithUserIDFetching(username string) ConfigFn {
	return func(c Config) (Config, error) {

		userID, err := GetUserByUsername(username)
		if err != nil {
			return c, fmt.Errorf("unable fetch user id: %v", err)
		}

		c.UserID = userID
		c.Username = username

		return c, nil
	}
}

// WithAPIToken
func WithAPIToken(deviceID string, apitoken string) ConfigFn {
	return func(c Config) (Config, error) {
		c.DeviceID = deviceID
		c.APIToken = apitoken
		return c, nil
	}
}

// WithDoLogin fetch user id from instagram
func WithDoLogin(username string, password string) ConfigFn {
	return func(c Config) (Config, error) {

		if c.DeviceID == "" {
			cTemp, err := WithDeviceID(util.GenerateAndroidDeviceID())(c)
			if err != nil {
				return c, err
			}
			c = cTemp
		}
		if c.UserID == 0 {
			cTemp, err := WithUserIDFetching(username)(c)
			if err != nil {
				return c, err
			}
			c = cTemp
		}
		if password == "" {
			return c, errors.New("passowrd is empty")
		}

		if c.APIToken != "" {
			ar, err := private.Auth(username, password, c.DeviceID)
			if err != nil {
				return c, err
			}

			c.APIToken = ar.Token
		}
		c.Username = username
		c.Password = password

		return c, nil
	}
}
