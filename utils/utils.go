package utils

import (
	"crypto/md5"
	"encoding/json"
	"objects"
)

func UserMapHash(m map[int](*objects.User)) [16]byte {
	arrBytes := []byte{}
	for _, item := range m {
		jsonBytes, _ := json.Marshal(item)
		arrBytes = append(arrBytes, jsonBytes...)
	}
	return md5.Sum(arrBytes)
}
