package db_models

import "fmt"

type Relationship struct {
	tableName struct{} `sql:"relationships"`
	Id        int64
	FromUid   int64
	ToUid     int64
	State     string
}

func UpdateRelation(fromUid, toUid int64, state string) {
	db := GetDB()
	relationShip := Relationship{}
	otherRelation := Relationship{}
	err := db.Model(&relationShip).Where("from_uid= ? and to_uid=?", fromUid, toUid).First()
	fmt.Printf("UpdateRelation with err: %v", err)
	err = db.Model(&otherRelation).Where("to_uid = ? and from_uid=?", fromUid, toUid).First()
	fmt.Printf("otherRelation with err: %v", err)
	if relationShip.State == "" {
		relationShip.State = state
		if otherRelation.State == "liked" && state == "liked" {
			relationShip.State = "matched"
			otherRelation.State = "matched"
			db.Model(&otherRelation).WherePK().Update()
		}
		relationShip.FromUid = fromUid
		relationShip.ToUid = toUid
		db.Model(&relationShip).Insert()
	} else {
		relationShip.State = state
		if otherRelation.State == "liked" && state == "liked" {
			relationShip.State = "matched"
			otherRelation.State = "matched"
			db.Model(&otherRelation).WherePK().Update()
		}
		if otherRelation.State == "matched" && state == "disliked" {
			otherRelation.State = "liked"
			db.Model(&otherRelation).WherePK().Update()
		}
		db.Model(&relationShip).WherePK().Update()
	}
}

func GetAllRelationships(uid int64) []Relationship {
	relationships := []Relationship{}
	db := GetDB()
	db.Model(&relationships).Where("from_uid=?", uid, uid).Select()
	return relationships
}
