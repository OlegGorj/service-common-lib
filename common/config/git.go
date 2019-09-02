package config

import (
  "fmt"
  "os"
  "io"
  git "gopkg.in/src-d/go-git.v4"
  "gopkg.in/src-d/go-git.v4/storage/memory"
  "gopkg.in/src-d/go-git.v4/plumbing"
  "gopkg.in/src-d/go-billy.v4/memfs"

  log "github.ibm.com/AdvancedAnalyticsCanada/service-common-lib/common/logging"
)

// getGitRepoConfigFile function
//@params
//
//@return
//
func GetGitRepoConfigFile (gitAccount, apiToken, repoName, branch, configFile string) (string, error) {

	url := fmt.Sprintf("https://%s:%s@%s", gitAccount, apiToken, repoName)
  log.Info("Cloning ", repoName ," ..")

  fs := memfs.New()
	storer := memory.NewStorage()
  _, err := git.Clone(storer, fs, &git.CloneOptions{
      URL: url,
      ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", branch)),
      SingleBranch:  true,
			Progress: os.Stdout,
  })
  if err != nil {
		log.Fatal(err)
		return "", err
	}
  thefile, err := fs.Open(configFile)
	if err != nil {
		log.Error(err, ": ", configFile)
		return "", err
	}

	tmp_configfile := fmt.Sprintf("/tmp/%s", configFile)
	log.Info("Creating temp file: ", tmp_configfile)
	file, err := os.Create(tmp_configfile)
	if err != nil {
		log.Error(err) ; return "", err
	}
	if _, err := io.Copy(file, thefile); err != nil {
		log.Error(err) // and exit
	}
	log.Info("INFO: temp file created: ", tmp_configfile)
	file.Close()

	// TODO: check if the file is readable or any errors

	return tmp_configfile, nil
}
