package apiv1

import (
	"prts-backend/src/model"
	"prts-backend/src/service"
)

func Stage(id string) model.Stage {
	var result model.Stage
	service.DB.
		Model(&model.Stage{}).
		Where(&model.Stage{Id: id}).
		Preload("Drops").
		Preload("Drops.Item").
		Preload("Hard").
		Preload("EnemyInstances.EnemyInstance").
		Preload("EnemyInstances.EnemyInstance.Enemy").
		Limit(1).
		Find(&result)
	return result
}
