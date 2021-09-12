package main

import (
	"crypto/md5"
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	errBadHashAlgorithm = errors.New("BAD_HASH_ALG")
)

type HashResult struct {
	status string
	hash   string
	err    error
}

func (hr HashResult) String() string {
	if len(hr.hash) > 0 {
		return hr.hash
	} else {
		return fmt.Sprint(hr.err)
	}
}

func (hr HashResult) GinJSON() gin.H {
	if hr.err != nil {
		return gin.H{
			"status": hr.status,
			"hash":   hr.hash,
			"err":    fmt.Sprint(hr.err),
		}
	}
	return gin.H{
		"status": hr.status,
		"hash":   hr.hash,
	}

}

func hashSHA256(msg string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(msg)))
}

func hashSHA224(msg string) string {
	return fmt.Sprintf("%x", sha256.Sum224([]byte(msg)))
}

func hashMD5(msg string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(msg)))
}

func computeHash(msg, alg string) HashResult {
	switch alg {
	case "SHA-256":
		return HashResult{
			status: "success",
			hash:   hashSHA256(msg),
		}
	case "SHA-224":
		return HashResult{
			status: "success",
			hash:   hashSHA224(msg),
		}
	case "MD5":
		return HashResult{
			status: "success",
			hash:   hashMD5(msg),
		}
	default:
		return HashResult{
			status: "failed",
			err:    errBadHashAlgorithm,
		}
	}
}
