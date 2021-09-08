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
		"key":         "https://gauk.as/keys",
		"pgp":         "https://gauk.as/gpg", // by default, all keys
		"ssh":         "https://gauk.as/keys",
		"ssh-ed25519": "https://gauk.as/keys",
		"ssh-key":     "https://gauk.as/keys",
		"ssh-keys":    "https://gauk.as/keys",
	}

	// All paths are relative to res folder. e.g. public-keys/ssh.key matches <repository>/res/public-keys/ssh.key
	urlFilepathMap = map[string]string{
		"gpg":         `public-keys/i_at_gaukasdotwang.gpg.key`, // all keys
		"gpg.sign":    `public-keys/i_at_gaukasdotwang_sign_cert_auth.gpg.key`,
		"gpg.cert":    `public-keys/i_at_gaukasdotwang_sign_cert_auth.gpg.key`,
		"gpg.auth":    `public-keys/i_at_gaukasdotwang_sign_cert_auth.gpg.key`,
		"gpg.encrypt": `public-keys/i_at_gaukasdotwang_encrypt.gpg.key`,
		"keys":        `public-keys/gaukas.ssh-ed25519.key`,
	}
)
