package domain

import (
	"github.com/tingxin/bingo/service/resource/cmd"
	"github.com/tingxin/bingo/service/resource/common"
	"github.com/tingxin/bingo/service/resource/dao"
	"github.com/tingxin/bingo/service/resource/model"
)

// AddResource used to add new resource
func AddResource(cmd cmd.NewResourceCmd) error {
	id := common.NewIntUID()

	doneSaveResource := make(chan error)
	doneSaveFields := make(chan error)

	cmd.ID = id
	go saveResource(cmd, doneSaveResource)
	go saveFields(id, cmd.Fields, doneSaveFields)

	err := <-doneSaveFields

	if err != nil {
		return err
	}

	err = <-doneSaveResource
	if err != nil {
		return err
	}
	return nil
}

func saveResource(cmd cmd.NewResourceCmd, done chan<- error) {
	err := dao.SaveResource(&cmd.ResourceM)
	if err != nil {
		done <- err
	}
	close(done)
}

func saveFields(resourceID int64, fields []*model.FieldM, done chan<- error) {
	err := dao.SaveResourceFields(resourceID, fields)
	if err != nil {
		done <- err
	}
	close(done)
}
