package hashcode

import "hash/crc32"

func Get(data string) int {
	return int(crc32.ChecksumIEEE([]byte(data)))
}

func Equa(data string) int {
	return int(crc32.ChecksumIEEE([]byte(data))) % 6
}
