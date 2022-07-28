package models

type AppSettings struct {
	BundleID      string `gorm:"primary_key;column:bundle_id"`
	ReviewMode    int64  `gorm:"column:review_mode"`
	ReviewVersion string `gorm:"column:review_version"`
}

func (AppSettings) TableName() string {
	return "ios_app_setting"
}
