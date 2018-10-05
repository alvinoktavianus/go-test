package business

import (
	"time"
)

type Region struct {
	Id        string `gorm:"primary_key"`
	Name      string
	BannerUrl string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Vertical struct {
	Id        string `gorm:"primary_key"`
	Name      string
	BannerUrl string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ProductLine struct {
	Id        string `gorm:"primary_key"`
	Name      string
	BannerUrl string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service struct {
	RegionId      string `gorm:"primary_key"`
	VerticalId    string `gorm:"primary_key"`
	ProductLineId string `gorm:"primary_key"`
	IsActive      bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
	OrderPosition int
	Region        Region      `gorm:"foreignkey:RegionId"`
	Vertical      Vertical    `gorm:"foreignkey:VerticalId"`
	ProductLine   ProductLine `gorm:"foreignkey:ProductLineId"`
}
