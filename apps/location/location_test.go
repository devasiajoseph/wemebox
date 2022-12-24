package location

import (
	"fmt"
	"testing"

	"github.com/devasiajoseph/wemebox/core"
	"github.com/devasiajoseph/wemebox/db/postgres"
)

func TestGetCityDistrict(t *testing.T) {
	core.Start()
	postgres.InitDB()

	id, err := GetCityDistrictID(1)
	if err != nil {
		t.Errorf("get city district id failed")
		fmt.Println(err)
	}

	fmt.Println(id)
}
