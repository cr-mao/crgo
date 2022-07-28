package appsetting

import (
	"context"
	"crgo/infra/bizerror"
	"crgo/infra/db"
	"crgo/models"
)

func GetAppSetting(ctx context.Context, bundleID string) *models.AppSettings {
	var setting models.AppSettings
	err := db.GetDb("default").Where("bundle_id = ?", bundleID).First(&setting).Error
	if err != nil {
		panic(bizerror.Wrap(1011001, "bundle_id不存在", err))
	}
	return &setting
}
