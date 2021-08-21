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
		"gpg":         "https://gauk.as/pgp",
		"keys":        "https://gauk.as/ssh",
		"ssh-ed25519": "https://gauk.as/ssh",
		"ssh-key":     "https://gauk.as/ssh",
		"ssh-keys":    "https://gauk.as/ssh",
	}

	// All paths are relative to res folder. e.g. public-keys/ssh.key matches <repository>/res/public-keys/ssh.key
	urlFilepathMap = map[string]string{
		"pgp": `public-keys/gaukas-pgp-i_at_gaukasdotwang.key`,
		"ssh": `public-keys/gaukas-ssh-ed25519.key`,
	}
)
