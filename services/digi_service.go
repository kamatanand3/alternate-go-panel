package services

import (
	"altpanel/repositories"
	"context"

	// "net/http"
	// "time"
	"altpanel/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type DigiScoreRequest struct {
	UserRefNumber     string `json:"user_reference_number" validate:"required"`
	Delay             int    `json:"delay"`
	EmploymentType    string `json:"employment_type" validate:"required"`
	CustomerFullName  string `json:"customer_full_name"`
	CustomerPanNumber string `json:"customer_pan_number"`
	Imei              string `json:"imei"`
	AndroidID         string `json:"android_id"`
	AdvertisingID     string `json:"advertising_id"`
	GlobalDeviceID    string `json:"global_device_id"`
	RequestSource     string `json:"request_source"`
	TransactionRefNo  string `json:"transaction_reference_number"`
}

func GetDigiScore(c *gin.Context, req DigiScoreRequest) (bson.M, error) {
	// ab JSON body se value aayegi
	utils.AppLog(c, "INFO", "GetDigiScore", map[string]interface{}{"user_ref_num": req.UserRefNumber}, "DigiService", "GetDigiScore")

	ctx := context.TODO()
	filter := bson.M{"user_reference_number": req.UserRefNumber}
	projection := bson.M{
		"user_reference_number": 1,
		"first_name":            1,
		"last_name":             1,
		"email":                 1,
	}
	customerRepo := repositories.NewCustomerRepository()
	return customerRepo.FindOne(ctx, filter, projection)
}
