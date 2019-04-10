package hashcode

import "hash/crc32"

func Get(data string) int {
	return int(crc32.ChecksumIEEE([]byte(data)) % 5)
}
