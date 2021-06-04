package handle

import (
	"backend_test/db_models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"sync"
)

var lockMap = sync.Map{}
var lockMapLock = sync.Mutex{}

type UpdateRelationshipsParam struct {
	State string `json:"state"`
}

func UpdateRelationships(c *gin.Context) {
	fromUid, _ := strconv.ParseInt(c.Param("user_id"), 10, 64)
	toUid, _ := strconv.ParseInt(c.Param("other_user_id"), 10, 64)
	param := UpdateRelationshipsParam{}
	c.BindJSON(&param)
	if fromUid == 0 || toUid == 0 || (param.State != "liked" && param.State != "disliked") {
		c.JSON(http.StatusBadRequest, gin.H{"message": "wrong param"})
		return
	}
	lock := getRelationshipLock(fromUid, toUid)
	lock.Lock()
	defer lock.Unlock()
	db_models.UpdateRelation(fromUid, toUid, param.State)
	c.JSON(200, gin.H{
		"message": "success",
	})
}

type RelationshipResp struct {
	UserId int64
	State  string
	Type   string
}

func GetRelationships(c *gin.Context) {
	fromUid, _ := strconv.ParseInt(c.Param("user_id"), 10, 64)
	relationships := db_models.GetAllRelationships(fromUid)
	resps := make([]RelationshipResp, len(relationships))
	for i := 0; i < len(relationships); i++ {
		resps[i] = RelationshipResp{UserId: relationships[i].ToUid, State: relationships[i].State, Type: "relationship"}
	}
	c.JSON(200, gin.H{
		"message": "success",
		"data":    resps,
	})
}

func getRelationshipLock(fromUid, toUid int64) sync.Mutex {
	firstUid := fromUid
	secondUid := toUid
	if fromUid > toUid {
		firstUid = toUid
		secondUid = fromUid
	}
	key := fmt.Sprintf("%d_%d", firstUid, secondUid)
	if lock, ok := lockMap.Load(key); ok {
		return lock.(sync.Mutex)
	}
	lockMapLock.Lock()
	defer lockMapLock.Unlock()
	if lock, ok := lockMap.Load(key); ok {
		return lock.(sync.Mutex)
	}
	lock := sync.Mutex{}
	lockMap.Store(key, lock)
	return lock
}
