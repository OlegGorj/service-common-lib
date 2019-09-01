# Service: service-common-lib

[![GitHub release](https://img.shields.io/github/release/qubyte/rubidium.svg)](https://github.com/OlegGorJ/service-common-lib/releases)
[![GitHub commits](https://img.shields.io/github/commits-since/SubtitleEdit/subtitleedit/3.4.7.svg)](https://github.com/OlegGorJ/service-common-lib/commits)
[![GitHub issues](https://img.shields.io/github/issues/OlegGorJ/service-common-lib.svg)](https://github.com/OlegGorJ/service-common-lib/issues)

Common repository to host common packages and code across multiple services and projects.


## Package `service-common-lib/common/logging`

### Use of the package

Import package:

```
import (
  log "service-common-lib/common/logging"
)
```

Set application name in `Init` function

```
func init() {
  ...
	log.SetAppName("myApp")
  ...
}
```

Start using it:

```
func main() {
  ...
	log.Info("Starting service...")
  ...

  log.Error("Some error..")
}
```


## Package `service-common-lib/common/config`

Import package:

```
import (
  log "service-common-lib/common/config"
)
```

Start using it:

```
func main() {
  //  get config from Git branch
	configfile, err := config.GetGitRepoConfigFile("git-username", "api-tocken", "my-repo", "master", "/config.json")

  // load config into viper
  v, err := config.ReadConfig(configfile)

  // access values using keys
  value=v.Get("my-key")
  ...
}

```


## Package `service-common-lib/service`

### Use of the package

Import package:

```
import (
  log "service-common-lib/common/logging"
)
```

Start using it:

```
func ApiHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write( []byte("v1") )
}

func main() {

  service.RegisterHandler("/api", "GET", ApiHandler)
	service.StartServer(g_port)

  ...
}

```
