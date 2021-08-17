package main

var (
	urlRedirectMap = map[string]string{
		"cv":          "https://gaukas.wang/latest/cv.pdf",
		"resume":      "https://gaukas.wang/latest/resume.pdf",
		"linkedin":    "https://www.linkedin.com/in/gaukaswang/",
		"github":      "https://github.com/Gaukas",
		"email":       "mailto:i@gaukas.wang",
		"keybase":     "https://keybase.io/gaukas",
		"gpg":         "https://gauk.as/pgp",
		"keys":        "https://gauk.as/ssh",
		"ssh-ed25519": "https://gauk.as/ssh",
		"ssh-key":     "https://gauk.as/ssh",
		"ssh-keys":    "https://gauk.as/ssh",
	}

	urlTextMap = map[string]string{
		"pgp": `-----BEGIN PGP PUBLIC KEY BLOCK-----
mDMEYPnkJxYJKwYBBAHaRw8BAQdASvBFKl2sfl9wuMWKlRFNwZpAEhpbL70JB9N1
8YTANI20G0dhdWthcyBXYW5nIDxpQGdhdWthcy53YW5nPoiWBBMWCAA+FiEE0Rxe
+jJU5+M/ilnUktinbec9oA8FAmD55CcCGyMFCQPCNPkFCwkIBwIGFQoJCAsCBBYC
AwECHgECF4AACgkQktinbec9oA+cHAD+Lxf3yguW/TVgYeXVa0mqhTVNPhvmZJWw
7++SHf56YkMBAP5U9a7eXvqbxOJLLjfa5heS74EmN8TMVgmHxm4YTecEuDgEYPnk
JxIKKwYBBAGXVQEFAQEHQME5/4UBKXmOWVVfN7/hXaSa7LyZkPHi+kfQSOK33VYc
AwEIB4h+BBgWCAAmFiEE0Rxe+jJU5+M/ilnUktinbec9oA8FAmD55CcCGwwFCQPC
NPkACgkQktinbec9oA+G7wD/bZZi/P8eONJjc3VNUVgOeWuf9JmDWuHHW99G6eAd
ZfMBAPQRD8jLVz2DF2JyqSSxZl911HpTX2rqCB6ZoFl+oWsJ
=Yvvs
-----END PGP PUBLIC KEY BLOCK-----
		`,
		"ssh": `ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIECqLe7H6DOvYqB7AiINtID7OFmrGIi6sNzILrEcqLtX`,
	}
)
