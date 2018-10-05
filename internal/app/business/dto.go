package business

type RegionDTO struct {
	Id        string `jsonapi:"primary,regions"`
	Name      string `jsonapi:"attr,name"`
	BannerUrl string `jsonapi:"attr,banner_url"`
}

type VerticalDTO struct {
	Id        string `jsonapi:"primary,regions"`
	Name      string `jsonapi:"attr,name"`
	BannerUrl string `jsonapi:"attr,banner_url"`
}
