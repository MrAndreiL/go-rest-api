package utils

import (
	"hash/fnv"
	"strconv"
	"time"
)

func GenerateEntityTag() string {
	date := time.Now().String()

	h := fnv.New32()
	h.Write([]byte(date))
	hash_date := int(h.Sum32())

	return strconv.Itoa(hash_date)
}
