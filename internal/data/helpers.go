package data

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/dmikhr/auth/pkg/user_v1"
)

// константы для фнукций конверсии
const (
	USER    = "USER"
	ADMIN   = "ADMIN"
	UNKNOWN = "UNKNOWN"
)

// RoleToStr - конвертация protobuf Role в string
func RoleToStr(role user_v1.Role) string {
	switch role {
	case 1:
		return USER
	case 2:
		return ADMIN
	default:
		return UNKNOWN
	}
}

// StrToRole - конвертация string в protobuf Role
func StrToRole(role string) user_v1.Role {
	switch role {
	case USER:
		return user_v1.Role_USER
	case ADMIN:
		return user_v1.Role_ADMIN
	default:
		return user_v1.Role_UNKNOWN
	}
}

// TimeToTimestamppb - конвертация time.Time в protobuf timestamppb.Timestamp
func TimeToTimestamppb(time time.Time) *timestamppb.Timestamp {
	return timestamppb.New(time)
}
