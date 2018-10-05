package business

import (
	"backend_go/internal/pkg/common"
	"encoding/json"
	"fmt"
	"github.com/google/jsonapi"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/http"
	"strconv"
	"time"
)

var connString = common.GetDatabaseConnection()

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	healthCheckDTO := common.HealthCheckDTO{
		AppName:     "styletheory_go_featured_v2",
		AppVersion:  "0.0.1-SNAPSHOT",
		BuildNumber: 0,
		BuildTime:   time.Now().UTC(),
	}

	json.NewEncoder(w).Encode(healthCheckDTO)
}

func GetAllRegions(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(common.POSTGRES, connString)
	defer db.Close()
	if common.IsAppEnvLocal() {
		db.LogMode(true)
	}
	if err != nil {
		panic(err)
	}

	var regions []Region
	db.Find(&regions)

	var regionsDTO []*RegionDTO
	for i := 0; i < len(regions); i++ {
		regionsDTO = append(regionsDTO, &RegionDTO{
			Id:        regions[i].Id,
			Name:      regions[i].Name,
			BannerUrl: regions[i].BannerUrl,
		})
	}

	w.WriteHeader(http.StatusOK)
	if err := jsonapi.MarshalPayload(w, regionsDTO); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetAllVerticals(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	regionId := q.Get("region")
	//order := q.Get("order")
	isActive, _ := strconv.ParseBool(q.Get("is_active"))

	db, err := gorm.Open(common.POSTGRES, connString)
	defer db.Close()
	if common.IsAppEnvLocal() {
		db.LogMode(true)
	}
	if err != nil {
		panic(err)
	}

	var verticals []Vertical
	var verticalsDTO []*VerticalDTO

	if regionId != "" {
		db.Raw("select distinct v.* "+
			"from verticals v "+
			"inner join services s2 on v.id = s2.vertical_id "+
			"where s2.region_id = ? "+
			"and s2.is_active = ? "+
			"order by v.order_position asc", regionId, isActive).Scan(&verticals)

	} else {
		db.Find(&verticals)
	}

	for i := 0; i < len(verticals); i++ {
		verticalsDTO = append(verticalsDTO, &VerticalDTO{
			Id:        verticals[i].Id,
			Name:      verticals[i].Name,
			BannerUrl: verticals[i].BannerUrl,
		})
	}

	w.WriteHeader(http.StatusOK)
	if err := jsonapi.MarshalPayload(w, verticalsDTO); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetAllProductLines(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	regionId := q.Get("region")
	order := q.Get("order")
	isActive := q.Get("is_active")
	vertical := q.Get("vertical")

	fmt.Println(regionId)
	fmt.Println(order)
	fmt.Println(isActive)
	fmt.Println(vertical)
}
