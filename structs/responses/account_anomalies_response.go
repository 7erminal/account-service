package responses

import "account_service/models"

type AccountAnomalyResponse struct {
	StatusCode    string                    `json:"statusCode"`
	StatusMessage string                    `json:"statusMessage"`
	Result        *models.Account_anomalies `json:"result,omitempty"`
}

type CustomerAccountAnomalyResponse struct {
	StatusCode    string                             `json:"statusCode"`
	StatusMessage string                             `json:"statusMessage"`
	Result        *models.Customer_account_anomalies `json:"result,omitempty"`
}

type CustomerAccountAnomaliesResponse struct {
	StatusCode    string                               `json:"statusCode"`
	StatusMessage string                               `json:"statusMessage"`
	Result        *[]models.Customer_account_anomalies `json:"result,omitempty"`
}
