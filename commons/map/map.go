package gs_commons_map

import "hash/crc32"

func Modulo(data string) int {
	return int(crc32.ChecksumIEEE([]byte(data)) % 10)
}
