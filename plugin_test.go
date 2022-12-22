package apollo

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
)

func Test(t *testing.T) {
	apolloPlugin := ApolloPlugin{}
	if !apolloPlugin.Enabled() {
		t.Fatal("apollo default enabled")
	}
	fmt.Println(viper.AllSettings())
	if apolloPlugin.Order() != defaultConfig["order"] {
		t.Fatalf("apollo order should be %v", defaultConfig["order"])
	}

}
