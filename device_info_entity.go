package gotercore

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeviceInfoEntity struct {
	UserID       primitive.ObjectID `json:"user_id" bson:"user_id"`
	DeviceCode   string             `json:"device_code" bson:"device_code"`
	DeviceModel  string             `kson:"device_model" bson:"device_model"`
	OSName       string             `json:"os_name" bson:"os_name"`
	OSVersion    string             `json:"os_version" bson:"os_version"`
	AppVersion   string             `json:"app_version" bson:"app_version"`
	FCMToken     string             `json:"fcm_token" bson:"fcm_token"`
	LanguageCode string             `json:"language_code" bson:"language_code"`
	Latitude     float32            `json:"latitude" bson:"latitude"`
	Longitude    float32            `json:"longitude" bson:"longitude"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
}

func NewEmptyDeviceInfo(userId primitive.ObjectID) *DeviceInfoEntity {
	u := &DeviceInfoEntity{
		UserID:    userId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return u
}

// NewDeviceInfo create a device information
func NewDeviceInfo(
	userId primitive.ObjectID,
	deviceCode, deviceModel, osName, osVersion, appVersion, fcmToken, languageCode string,
	latitude, longitude float32) (*DeviceInfoEntity, error) {

	u := &DeviceInfoEntity{
		UserID:       userId,
		DeviceCode:   deviceCode,
		DeviceModel:  deviceModel,
		OSName:       osName,
		OSVersion:    osVersion,
		AppVersion:   appVersion,
		FCMToken:     fcmToken,
		LanguageCode: languageCode,
		Latitude:     latitude,
		Longitude:    longitude,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	err := u.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return u, nil
}

// Validate  data
func (info *DeviceInfoEntity) Validate() error {
	if info.AppVersion == "" || info.FCMToken == "" {
		return ErrInvalidEntity
	}
	return nil
}
