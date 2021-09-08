package main

var (
	urlRedirectMap = map[string]string{
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

	// All paths are relative to res folder. e.g. public-keys/ssh.key matches <repository>/res/public-keys/ssh.key
	urlFilepathMap = map[string]string{
		"gpg":  `public-keys/i_at_gaukasdotwang.gpg.key`, // all keys
		"keys": `public-keys/gaukas.ssh-ed25519.key`,
	}
)
