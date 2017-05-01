package models

import (
	"taskmgr/common"

	"gopkg.in/mgo.v2/bson"
)

type Task struct {
	ID   bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name string        `bson:"name" json:"name"`
	Desc string        `bson:"desc" json:"desc"`
}

var Tasks = new(tasks)

type tasks struct{}

func (tasks) FindAll() ([]*Task, error) {
	var ts []*Task
	return ts, common.DB.Tasks.Find(nil).All(&ts)
}

func (tasks) FindOne(id string) (*Task, error) {
	var t *Task
	return t, common.DB.Tasks.FindId(bson.ObjectIdHex(id)).One(&t)
}

func (tasks) Create(name, desc string) (*Task, error) {
	t := &Task{
		ID:   bson.NewObjectId(),
		Name: name,
		Desc: desc,
	}
	if err := common.DB.Tasks.Insert(t); err != nil {
		return nil, err
	}
	return t, nil

}

func (tasks) Update(id, name, desc string) error {
	if err := common.DB.Tasks.UpdateId(
		bson.ObjectIdHex(id),
		bson.M{"$set": bson.M{
			"name": name,
			"desc": desc,
		}}); err != nil {
		return err
	}
	return nil
}

func (tasks) Delete(id string) error {
	return common.DB.Tasks.RemoveId(bson.ObjectIdHex(id))
}
