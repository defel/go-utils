package utils

import "crypto/rand"

func GetRandomBytes(len int8) []byte {
	key := make([]byte, len)

	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}

	return key
}
