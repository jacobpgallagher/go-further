package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/SocialCodeInc/go-common/api"
	mw "github.com/SocialCodeInc/go-common/echo"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

// START STRUCT_1 OMIT
type thinger struct {
	ID   uint64    `json:"id"`
	Name string    `json:"name"`
	Time time.Time `json:"time"`
}

var thingers map[uint64]thinger

func init() { thingers = make(map[uint64]thinger, 0) }

func (t *thinger) New() api.Model {
	return &thinger{}
}

func (t *thinger) Get(id string) (api.Model, error) {
	asInt, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}
	fetched, exists := thingers[asInt]
	if !exists {
		return nil, fmt.Errorf("404")
	}
	return &fetched, nil
}

// END STRUCT_1 OMIT
// START STRUCT_2 OMIT
func (t *thinger) Delete() (api.Model, error) {
	delete(thingers, t.ID)
	return nil, nil

}

func (t *thinger) Save() (api.Model, error) {
	if t.ID > 0 {
		thingers[t.ID] = *t
		return t, nil
	}
	max := uint64(0)
	for k := range thingers {
		if k > max {
			max = k
		}
	}
	t.ID = max + 1
	thingers[max+1] = *t
	return t, nil

}

// END STRUCT_2 OMIT
// START STRUCT_3 OMIT
func (t *thinger) GetList(filters map[string][]interface{}, offset uint64, limit uint64) ([]api.Model, error) {
	foundThingers := []api.Model{}
	for tidx, oneT := range thingers {
		allThere := true
		for k, values := range filters {
			for _, v := range values {
				switch k {
				case "name":
					if oneT.Name != v {
						allThere = false
					}
				case "id":
					if oneT.ID != v {
						allThere = false
					}
				default:
					return nil, fmt.Errorf("Invalid filters")
				}
			}
		}
		if allThere {
			at := thingers[tidx]
			foundThingers = append(foundThingers, &at)
		}
	}
	return foundThingers, nil
}

// END STRUCT_3 OMIT
// START STRUCT_4 OMIT
func (t *thinger) SaveMany(models []api.Model) ([]api.Model, []error) {
	saved := make([]api.Model, len(models))
	errors := make([]error, len(models))
	for i, m := range models {
		sm, err := m.Save()
		errors[i] = err
		saved[i] = sm
	}

	return saved, errors
}

func (t *thinger) GetFilterOptions() map[string]string {
	return map[string]string{
		"name": "string",
		"id":   "uint64",
	}
}

func (t *thinger) Validate() (api.Model, error) {
	if t.Name == "Leeroy Jenkins" {
		return nil, fmt.Errorf("No. Just no.")
	}
	return t, nil
}

// END STRUCT_4 OMIT
// START MAIN OMIT
func main() {

	e := echo.New()

	e.Use(mw.Logger())

	v1 := api.NewAPI("/api/v1", e)

	if err := v1.AddResource(api.NewResource(&thinger{},
		"thinger",
		[]int{api.GET, api.DELETE, api.PUT, api.POST},
		[]int{api.GET, api.DELETE, api.PUT, api.POST},
		"ID")); err != nil {
		panic(err)
	}
	v1.AddDocs("Thingers", "v1")
	v1.AddSwaggerUI("swagger-ui")

	e.Run(standard.New(":1234"))
}

// END MAIN OMIT
