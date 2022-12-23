package apollo

import (
	"bytes"
	"io"

	"github.com/apolloconfig/agollo/v4"
	"github.com/spf13/viper"
)

type apolloRemoteConfig struct {
	client agollo.Client
}

func (a *apolloRemoteConfig) Get(rp viper.RemoteProvider) (io.Reader, error) {
	res := a.client.GetConfig(rp.Path()).GetContent()
	return bytes.NewReader([]byte(res)), nil
}

func (a *apolloRemoteConfig) Watch(rp viper.RemoteProvider) (io.Reader, error) {
	return a.Get(rp)
}

func (a *apolloRemoteConfig) WatchChannel(rp viper.RemoteProvider) (<-chan *viper.RemoteResponse, chan bool) {
	// TODO
	return nil, nil
}

func initProvider(c agollo.Client) {
	viper.SupportedRemoteProviders = append(viper.SupportedRemoteProviders, "apollo")
	viper.RemoteConfig = &apolloRemoteConfig{c}
}
