package apollo

import (
	"fmt"

	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
	plugin "github.com/jsmzr/boot-plugin"
	"github.com/spf13/viper"
)

const configPrefix = "boot.apollo."

var defaultConfig map[string]interface{} = map[string]interface{}{
	"enabled":           true,
	"order":             -21,
	"namespace":         "application",
	"cluster":           "default",
	"isBackupConfig":    false,
	"mustStart":         true,
	"syncServerTimeout": 5,
}

type ApolloPlugin struct{}

func (a *ApolloPlugin) Load() error {
	appId := viper.GetString(configPrefix + "appId")
	address := viper.GetString(configPrefix + "address")
	if appId == "" || address == "" {
		return fmt.Errorf("apollo appId and address should be set, appId:[%s], address:[%s]", appId, address)
	}
	namespace := viper.GetString(configPrefix + "namespace")
	secretKey := viper.GetString(configPrefix + "secretKey")

	client, err := createClient(address, appId, namespace, secretKey)
	if err != nil {
		return err
	}
	initProvider(client)
	if err := viper.AddRemoteProvider("apollo", address, namespace); err != nil {
		return err
	}
	return viper.ReadRemoteConfig()
}

func (a *ApolloPlugin) Order() int {
	return viper.GetInt(configPrefix + "order")
}

func (a *ApolloPlugin) Enabled() bool {
	return viper.GetBool(configPrefix + "enabled")
}

func createClient(address, appId, namespace, secretKey string) (agollo.Client, error) {
	c := &config.AppConfig{
		AppID:             appId,
		Cluster:           viper.GetString(configPrefix + "cluster"),
		IP:                address,
		NamespaceName:     namespace,
		Secret:            secretKey,
		IsBackupConfig:    viper.GetBool(configPrefix + "isBackupConfig"),
		BackupConfigPath:  viper.GetString(configPrefix + "backupConfigPath"),
		MustStart:         viper.GetBool(configPrefix + "mustStart"),
		Label:             viper.GetString(configPrefix + "label"),
		SyncServerTimeout: viper.GetInt(configPrefix + "syncServerTimeout"),
	}
	// viper read agollo content is prop
	viper.SetConfigType("prop")
	return agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})
}

func init() {
	plugin.InitDefaultConfig(configPrefix, defaultConfig)
	plugin.Register("apollo", &ApolloPlugin{})
}
