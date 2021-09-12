package main

import (
	"fmt"
	"io/ioutil"
	"sync"
	"time"
)

// urlMapping priority:
// OverrideMap > Database(if used) > Map
// Note: Content in Map will not be recovered to original state by goroutine adminitered reloading

var (
	urlMappingMutex sync.RWMutex = sync.RWMutex{}
	urlRedirectMap               = map[string]string{
		// Online Identity
		"github":   "https://github.com/Gaukas",
		"email":    "mailto:i@gaukas.wang",
		"keybase":  "https://keybase.io/gaukas",
		"telegram": "https://t.me/GaukasWang",

		// Career
		"cv":       "https://gaukas.wang/latest/cv.pdf",
		"resume":   "https://gaukas.wang/latest/resume.pdf",
		"linkedin": "https://www.linkedin.com/in/gaukaswang/",

		// Keys
		"gpg.sign":    `https://gauk.as/gpg`,
		"gpg.cert":    `https://gauk.as/gpg`,
		"gpg.auth":    `https://gauk.as/gpg`,
		"gpg.encrypt": `https://gauk.as/gpg`,
		"key":         "https://gauk.as/keys",
		"pgp":         "https://gauk.as/gpg", // by default, all keys
		"ssh":         "https://gauk.as/keys",
		"ssh-ed25519": "https://gauk.as/keys",
		"ssh-key":     "https://gauk.as/keys",
		"ssh-keys":    "https://gauk.as/keys",
	}
	urlRedirectOverrideMap = map[string]string{}

	// All paths are relative to res folder. e.g. public-keys/ssh.key matches <repository>/res/public-keys/ssh.key
	urlFilepathMap = map[string]string{
		"gpg":  `public-keys/i_at_gaukasdotwang.gpg.key`, // all keys
		"keys": `public-keys/gaukas.ssh-ed25519.key`,
	}
	urlFilepathOverrideMap = map[string]string{}
)

func initUrlMapping() error {
	if useMysql { // if not using MySQL, no way to dynamically update the mapping.
		go routineReloadUrlMapping()
	}

	return reloadUrlMapping()
}

func routineReloadUrlMapping() {
	for {
		time.Sleep(RELOAD_INTERVAL)
		err := reloadUrlMapping()
		if err != nil {
			fmt.Printf("Error in routinely reloading URL mapping: %s", err)
		}
	}
}

func reloadUrlMapping() error {
	urlMappingMutex.Lock()
	defer urlMappingMutex.Unlock()

	// return loadUrlMappingFromDatabase()
	if useMysql {
		if err := loadUrlMappingFromDatabase(); err != nil {
			return err
		}
	}
	for key, val := range urlRedirectOverrideMap {
		urlRedirectMap[key] = val
	}
	for key, val := range urlFilepathOverrideMap {
		urlRedirectMap[key] = val
	}

	return nil
}

func loadUrlMappingFromDatabase() error {
	// TODO: Load from database
	return nil
}

func urlRedirect(name string) (string, bool) {
	urlMappingMutex.RLock()
	defer urlMappingMutex.RUnlock()
	if redirectUrl, ok := urlRedirectMap[name]; ok {
		return redirectUrl, true
	}
	return "", false
}

func urlLoadFile(name string) (string, bool) {
	urlMappingMutex.RLock()
	defer urlMappingMutex.RUnlock()
	if responseTxt, ok := urlFilepathMap[name]; ok {
		filepath := fmt.Sprintf("res/%s", responseTxt)
		content, err := ioutil.ReadFile(filepath)
		if err != nil {
			fmt.Printf("Err: %s\n", err)
			return "", false
		}
		return string(content), true
		// return filepath, true
	}
	return "", false
}
