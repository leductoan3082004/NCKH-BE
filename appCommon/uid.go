package appCommon

import (
	"database/sql/driver"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const xorKey = "ThisIsKB2A*##!@$SecretKey((!%#!HIHI!@$%ThuMinkCuti"

func xorCipher(input string, key string) string {
	keyLength := len(key)
	output := make([]byte, len(input))

	for i := 0; i < len(input); i++ {
		output[i] = input[i] ^ key[i%keyLength]
	}

	return string(output)
}

type UID struct {
	localID    uint32
	objectType int
	shardID    uint32
}

func NewUID(localID uint32, objType int, shardID uint32) UID {
	return UID{
		localID:    localID,
		objectType: objType,
		shardID:    shardID,
	}
}

func (uid UID) String() string {
	val := uint64(uid.localID)<<28 | uint64(uid.objectType)<<18 | uint64(uid.shardID)<<0
	encodedVal := fmt.Sprintf("%v", val)
	xorVal := xorCipher(encodedVal, xorKey)
	return base64.RawURLEncoding.EncodeToString([]byte(xorVal))
}

func (uid UID) GetLocalID() uint32 {
	return uid.localID
}

func (uid UID) GetShardID() uint32 {
	return uid.shardID
}

func (uid UID) GetObjectType() int {
	return uid.objectType
}

func DecomposeUID(s string) (UID, error) {
	decoded, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return UID{}, err
	}

	xorVal := xorCipher(string(decoded), xorKey)
	uidVal, err := strconv.ParseUint(xorVal, 10, 64)
	if err != nil {
		return UID{}, err
	}

	if (1 << 18) > uidVal {
		return UID{}, errors.New("wrong uid")
	}

	u := UID{
		localID:    uint32(uidVal >> 28),
		objectType: int(uidVal >> 18 & 0x3FF),
		shardID:    uint32(uidVal >> 0 & 0x3FFFF),
	}

	return u, nil
}

func FromBase58(s string) (UID, error) {
	return DecomposeUID(s)
}

func (uid UID) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", uid.String())), nil
}

func (uid *UID) UnmarshalJSON(data []byte) error {
	decodeUID, err := FromBase58(strings.Replace(string(data), "\"", "", -1))

	if err != nil {
		return err
	}

	uid.localID = decodeUID.localID
	uid.shardID = decodeUID.shardID
	uid.objectType = decodeUID.objectType

	return nil
}

func (uid *UID) Value() (driver.Value, error) {
	if uid == nil {
		return nil, nil
	}
	return int64(uid.localID), nil
}

func (uid *UID) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var i uint32

	switch t := value.(type) {
	case int:
		i = uint32(t)
	case int8:
		i = uint32(t)
	case int16:
		i = uint32(t)
	case int32:
		i = uint32(t)
	case int64:
		i = uint32(t)
	case uint8:
		i = uint32(t)
	case uint16:
		i = uint32(t)
	case uint32:
		i = t
	case uint64:
		i = uint32(t)
	case []byte:
		a, err := strconv.Atoi(string(t))
		if err != nil {
			return err
		}
		i = uint32(a)
	default:
		return errors.New("invalid Scan Source")
	}

	*uid = NewUID(i, 0, 1)

	return nil
}
