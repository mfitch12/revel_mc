package controllers

import (
	"github.com/revel/revel"
	"mailCongress/app/models"
)

type RepresentativeController struct {
	*revel.Controller
}

func (c RepresentativeController) GetByZip(id string) revel.Result {
	reps := models.Representatives{
		models.Representative{"1", "Matt"},
	}
	response := JsonResponse{}
	response.Data = reps
	return c.RenderJSON(response)
}

// func getRepresentatives() []*models.Representative {
// 	return
// }