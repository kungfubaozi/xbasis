package gs_commons_encrypt

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"zskparker.com/foundation/pkg/tool/number"
)

func str(str string) string {
	var b string
	for _, c := range str {
		b = fmt.Sprint(b, c)
	}
	return b
}

func SHA1(value string) string {
	h := sha1.New()
	h.Write([]byte(str(value)))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func SHA256(value string) string {
	h := sha256.New()
	h.Write([]byte(str(value)))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func SHA512(value string) string {
	h := sha512.New()
	h.Write([]byte(str(value)))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func SHA256_BASE64(value string) string {
	return base64.StdEncoding.EncodeToString([]byte(SHA256(value)))
}

func SHA1_BASE64(value string) string {
	return base64.StdEncoding.EncodeToString([]byte(SHA1(value)))
}

func MD5_SHA256(value string) string {
	return SHA256(Md5(value))
}

func MD5_SHA1_256(value string) string {
	return SHA256(SHA1(Md5(value)))
}

func MD5_SHA1_256_512(value string) string {
	return SHA512(SHA256(SHA1(Md5(value))))
}

func SHA1_256_512(value string) string {
	return SHA512(SHA256(SHA1(value)))
}

func SHA256_1_512(value string) string {
	return SHA512(SHA1(SHA256(value)))
}

func Rnd_SHA1_256_512(size int) string {
	return SHA1_256_512(fs_tools_number.GetRndNumber(size))
}
