package config

import (
  "net/http"
  "io/ioutil"

  "github.com/spf13/viper"
  "github.ibm.com/AdvancedAnalyticsCanada/service-common-lib/common/util"
)

var (
	g_apiVer string
)

//@params
//
//@return
//
func init() {

  g_apiVer = util.GetENV("APIVER") ; if g_apiVer == "" { g_apiVer = "/api/v1/" }

}
//@params
//
//@return
//
func ReadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()
  v.SetConfigType("json")
	v.SetConfigFile(filename)
	v.AddConfigPath("/tmp")
	v.AutomaticEnv()
	err := v.ReadInConfig()
  // TODO: need a decision - do we make all keys case-insensetive or leave it case-sensetive?
	return v, err
}
//@params
//
//@return
//
func GetConfigKV(key string) string{
	if key == "" { return "" }
	resp, err := http.Get("http://service-config-data.default:8000" + g_apiVer + key)
	if err != nil { return "" }
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}
