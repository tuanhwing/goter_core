package gotercore

import (
	"regexp"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RoleType int
type GenderType string

const (
	AdminRole   RoleType = 20
	ManagerRole RoleType = 10
	PremiumRole RoleType = 5
	GuestRole   RoleType = 0

	MaleGenderType   GenderType = "male"
	FemaleGenderType GenderType = "female"
	OtherGenderType  GenderType = "other"
)

type AuthorizationEntity struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type FileEntity struct {
	Key          string `bson:"key" json:"key"`
	Type         string `bson:"type" json:"type"` //value = image/video
	OriginUrl    string `bson:"origin_url" json:"origin_url"`
	ThumbnailUrl string `bson:"thumbnail_url,omitempty" json:"thumbnail_url"`
}

type DeviceInfoHeaderRequest struct {
	DeviceCode  string `json:"device_code,omitempty" bson:"device_code,omitempty"`
	DeviceModel string `json:"device_model,omitempty" bson:"device_model,omitempty"`
	OSName      string `json:"os_name,omitempty" bson:"os_name,omitempty"`
	OSVersion   string `json:"os_version,omitempty" bson:"os_version,omitempty"`
	AppVersion  string `json:"app_version,omitempty" bson:"app_version,omitempty"`
}

type UserEntity struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	Name       string             `json:"name"`
	Address    string             `json:"address"`
	Gender     GenderType         `json:"gender"`
	Avatar     string             `json:"avatar"`
	RoleType   RoleType           `bson:"role_type" json:"role_type"`
	DayOfBirth string             `bson:"day_of_birth" json:"day_of_birth"`
	Phone      PhoneNumber        `json:"phone"`
	CreatedAt  time.Time          `bson:"created_at" json:"-"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"-"`
}

func NewAuthorization(accessToken, refreshToken string) *AuthorizationEntity {
	return &AuthorizationEntity{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

// NewUser create a new user
func NewUser(dialCode, phone string) (*UserEntity, error) {
	p, err := NewPhoneNumber(dialCode, phone)
	defaultName := "Goter"

	if err != nil {
		return nil, err
	}

	u := &UserEntity{
		ID:        primitive.NewObjectID(),
		Phone:     *p,
		Name:      defaultName,
		RoleType:  GuestRole,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = u.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return u, nil
}

// Validate  data
func (u *UserEntity) Validate() error {
	return nil
}

// Get name with role type
func (role RoleType) String() string {
	// list of role names
	names := [...]string{
		"admin",
		"manager",
		"guest"}

	// `role`: là một trong các giá trị của hằng số RoleType.
	// Nếu hằng số là Admin, thì role có giá trị là 0.
	// Bắt lỗi trong trường hợp `role` nằm ngoài khoảng của RoleType
	if role < AdminRole || role > GuestRole {
		return "Unknown"
	}
	// trả về tên của 1 hằng số Weekday từ mảng names bên trên
	return names[role]

}

// Return brief user information
func (user *UserEntity) Brief() interface{} {
	briefInfo := struct {
		ID       primitive.ObjectID `json:"id"`
		Name     string             `json:"name"`
		Avatar   string             `json:"avatar"`
		RoleType RoleType           `bson:"role_type" json:"role_type"`
	}{
		ID:       user.ID,
		Name:     user.Name,
		Avatar:   user.Avatar,
		RoleType: user.RoleType,
	}

	return briefInfo
}

type PhoneNumber struct {
	DialCode        string `json:"dial_code" bson:"dial_code,omitempty" form:"dial_code" binding:"required,min=2"`
	PhoneNumber     string `json:"phone_number" form:"phone_number" bson:"phone_number,omitempty" binding:"required,min=6,numeric"`
	FullPhoneNumber string `json:"full_phone_number" bson:"full_phone_number"`
}

// NewPhoneNumber create a new user
func NewPhoneNumber(dialCode, phoneNumber string) (*PhoneNumber, error) {
	p := &PhoneNumber{
		DialCode:        dialCode,
		PhoneNumber:     phoneNumber,
		FullPhoneNumber: dialCode + phoneNumber,
	}
	err := p.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return p, nil
}

// Validate  data
func (p *PhoneNumber) Validate() error {
	re := regexp.MustCompile(`\+[1-9][0-9]{6,12}`)
	phoneNumber := p.DialCode + p.PhoneNumber
	if re.MatchString(phoneNumber) {
		return nil
	}
	return ErrInvalidEntity
}
