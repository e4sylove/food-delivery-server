package common

import (
	"fmt"

	"github.com/btcsuite/btcutil/base58"
)

type UID struct {
	localID uint32
	objectType int
	shardID int32
}

func NewUID(localID uint32, objType int, shardID int32) UID {
	return UID{
		localID: localID,
		objectType: objType,
		shardID: shardID,
	}
}

func (uid UID) String() string {
	val := uint64(uid.localID) << 28 | uint64(uid.objectType) << 18 | uint64(uid.shardID) << 0
	return base58.Encode([]byte(fmt.Sprintf("%v", val)))
}

func (uid UID) GetLocalID() uint32 {
	return uid.localID
}

func (uid UID) GetShardID() int32 {
	return uid.shardID
}

func (uid UID) GetObjectType() int {
	return uid.objectType
}