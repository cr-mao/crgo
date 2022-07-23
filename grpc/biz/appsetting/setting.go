package appsetting

import (
	"context"
	"crgo/infra/bizerror"
	"crgo/infra/db"
)

type Settings struct {
	BundleID      string
	ReviewMode    int64
	ReviewVersion string
}

func GetAppSetting(ctx context.Context, bundleID string) *Settings {
	var setting *Settings
	err := db.GetDb("default").Where("bundle_id = ?").First(&setting).Error
	if err != nil {
		panic(bizerror.Wrap(1011001, "bundle_id不存在", err))
	}
	return setting
}
