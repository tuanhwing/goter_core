package gotercore

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
)

func RegisSet(c redis.Client, key string, value interface{}) error {
	p, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.Set(context.Background(), key, p, 0).Err()
}

func RegisGet(c redis.Client, key string, dest interface{}) error {
	p, err := c.Get(context.Background(), key).Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(p, dest)
}

func RegisDel(c *redis.Client, key string) error {
	return c.Del(context.Background(), key).Err()
}

func ToDoc(v interface{}) (doc bson.D, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}

func ParseAppVersion(s string) [4]int {
	var results [4]int

	parts := strings.Split(s, ".")
	if len(parts) < 3 {
		return results
	}

	nameAndBuild := strings.Split(parts[2], "+")
	if len(nameAndBuild) < 2 {
		return results
	}

	//Version number Part 1
	versionNumber1, _ := strconv.Atoi(parts[0])
	results[0] = versionNumber1

	//Version number Part 2
	versionNumber2, _ := strconv.Atoi(parts[1])
	results[1] = versionNumber2

	//Version number Part 3
	versionNumber3, _ := strconv.Atoi(nameAndBuild[0])
	results[2] = versionNumber3

	//Build number
	buildNumber, _ := strconv.Atoi(nameAndBuild[1])
	results[3] = buildNumber

	return results
}

func ParseDeviceInfo(s string) (DeviceInfoHeaderRequest, error) {
	var data DeviceInfoHeaderRequest
	err := json.Unmarshal([]byte(s), &data)
	return data, err
}

func GetAuthenticationKey(userId string) string {
	return "authentication_" + userId
}

func IsRefreshToken(path string) bool {
	return strings.Contains(path, "/refreshToken")
}
