package admitad

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/4kayDev/admitad-integration/internal/pkg/clients/admitad/models"
	"github.com/4kayDev/admitad-integration/internal/utils/config"
	"github.com/4kayDev/admitad-integration/internal/utils/jsoner"
	"github.com/dr3dnought/exerror"
	requestbuidler "github.com/dr3dnought/request_builder"
)

type Client struct {
	cfg     *config.AdmitadConfig
	builder *requestbuidler.RequestBuilder

	httpClient *http.Client

	refreshAccessToken string
}

func New(cfg *config.AdmitadConfig) *Client {
	return &Client{
		cfg:     cfg,
		builder: requestbuidler.New(cfg.URL),

		httpClient: &http.Client{},
	}
}

type GetPublisherRecordsInput struct {
	Offset int
	Limit  int
}

func (c *Client) GetPublisherRecords(input *GetAffiliatesInput) ([]models.PublisherRecord, *exerror.ExtendedError) {
	token, exerr := c.syncToken()
	if exerr != nil {
		return nil, exerr
	}

	type responseType struct {
		Items []models.PublisherRecord `json:"results"`
	}

	response, err := c.builder.SetMethod("GET").SetPath(fmt.Sprintf("statistics/campaigns/?offset=%d&limit=%d&language=ru", input.Offset, input.Limit)).SetHeaders(map[string]string{
		"Authorization": "Bearer " + token,
	}).Build().Execute(c.httpClient)

	rawBody, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return nil, exerror.New(ErrInvalidEntity, exerror.Important(), exerror.Message(fmt.Sprintf("API Response can not be read")))
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return nil, parseApiError(response.StatusCode, nil)
	}

	records := new(responseType)
	err = json.Unmarshal(rawBody, records)
	if err != nil {
		return nil, exerror.New(ErrInvalidEntity, exerror.Important(), exerror.Message(fmt.Sprintf("%s can not be casted to Authorization type", string(rawBody))))
	}

	return records.Items, nil
}

type GetAffiliatesInput struct {
	Limit  int
	Offset int
}

func (c *Client) GetAffiliates(input *GetAffiliatesInput) ([]models.Affiliate, *exerror.ExtendedError) {
	token, exerr := c.syncToken()
	if exerr != nil {
		return nil, exerr
	}

	type responseType struct {
		Items []models.Affiliate `json:"results"`
	}

	response, err := c.builder.SetMethod("GET").SetPath(fmt.Sprintf("advcampaigns/?website=%s&offset=%d&limit=%d&language=ru", c.cfg.WebmasterId, input.Offset, input.Limit)).SetHeaders(map[string]string{
		"Authorization": "Bearer " + token,
	}).Build().Execute(c.httpClient)

	rawBody, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return nil, exerror.New(ErrInvalidEntity, exerror.Important(), exerror.Message(fmt.Sprintf("API Response can not be read")))
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return nil, parseApiError(response.StatusCode, nil)
	}

	affilates := new(responseType)
	err = json.Unmarshal(rawBody, affilates)
	if err != nil {
		fmt.Println(err)
		return nil, exerror.New(ErrInvalidEntity, exerror.Important(), exerror.Message(fmt.Sprintf("API Response can not be casted to Affiliates type")))
	}

	return affilates.Items, nil
}

type GetAffiliateByIdInput struct {
	AdmiatdId int
}

func (c *Client) GetAffiliateById(input *GetAffiliateByIdInput) (*models.Affiliate, *exerror.ExtendedError) {
	token, exerr := c.syncToken()
	if exerr != nil {
		return nil, exerr
	}

	response, err := c.builder.SetMethod("GET").SetPath(fmt.Sprintf("advcampaigns/%d/?language=ru", input.AdmiatdId)).SetHeaders(map[string]string{
		"Authorization": "Bearer " + token,
	}).Build().Execute(c.httpClient)
	if err != nil {
		return nil, exerror.New(ErrRequest, exerror.Important(), exerror.Message(fmt.Sprintf("API Response can not be read")))
	}
	rawBody, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return nil, parseApiError(response.StatusCode, rawBody)
	}

	affiliate := new(models.Affiliate)
	err = json.Unmarshal(rawBody, affiliate)
	if err != nil {
		fmt.Println(err)
		return nil, exerror.New(ErrInvalidEntity, exerror.Important(), exerror.Message(fmt.Sprintf("API Response can not be casted to Affiliates type")))
	}

	return affiliate, nil
}

type BuildDeeplinkInput struct {
	AdmitadId int
	SubId     string
	SiteURL   string
}

func (c *Client) BuildDeeplink(input *BuildDeeplinkInput) (*string, *exerror.ExtendedError) {
	token, exerr := c.syncToken()
	if exerr != nil {
		return nil, exerr
	}

	response, err := c.builder.SetMethod("GET").SetPath(fmt.Sprintf("deeplink/%s/advcampaign/%d/?subid=%s&ulp=%s", c.cfg.WebmasterId, input.AdmitadId, input.SubId, input.SiteURL)).SetHeaders(map[string]string{
		"Authorization": "Bearer " + token,
	}).Build().Execute(c.httpClient)
	if err != nil {
		return nil, exerror.New(ErrRequest, exerror.Important(), exerror.Message(fmt.Sprintf("API Response can not be read")))
	}
	rawBody, err := io.ReadAll(response.Body)
	fmt.Println(string(rawBody))
	defer response.Body.Close()

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return nil, parseApiError(response.StatusCode, nil)
	}

	deeplink := make([]string, 0)
	err = json.Unmarshal(rawBody, &deeplink)
	if err != nil {
		fmt.Println(err)
		return nil, exerror.New(ErrInvalidEntity, exerror.Important(), exerror.Message(fmt.Sprintf("API Response can not be casted to Deeplink type")))
	}

	fmt.Println(jsoner.Jsonify(deeplink))

	if len(deeplink) != 0 {
		return &deeplink[0], nil
	}

	return nil, exerror.New(ErrInternal, exerror.Important(), exerror.Message("Deeplink was not found in response body"))
}

func (c *Client) syncToken() (string, *exerror.ExtendedError) {
	auth, exerr := c.refreshToken()
	if exerr != nil {
		return "", exerr
	}

	c.refreshAccessToken = auth.RefreshToken
	return auth.AccessToken, nil
}

func (c *Client) refreshToken() (*models.Authorization, *exerror.ExtendedError) {
	response, err := c.builder.SetMethod("POST").SetPath("token/").SetHeaders(map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": "Basic " + c.cfg.ClientB64,
	}).SetBody([]byte(c.buildRefreshTokenBody())).Build().Execute(c.httpClient)
	if err != nil {
		fmt.Println(err)
		return nil, exErrRequest
	}

	rawBody, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return nil, exerror.New(ErrInvalidEntity, exerror.Important(), exerror.Message(fmt.Sprintf("API Response can not be read")))
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return nil, parseApiError(response.StatusCode, rawBody)
	}

	authorization := new(models.Authorization)
	err = json.Unmarshal(rawBody, authorization)
	if err != nil {
		return nil, exerror.New(ErrInvalidEntity, exerror.Important(), exerror.Message(fmt.Sprintf("%s can not be casted to Authorization type", string(rawBody))))
	}

	return authorization, nil
}

func (c *Client) buildRefreshTokenBody() string {
	if c.refreshAccessToken != "" {
		return fmt.Sprintf("grant_type=refresh_token&client_id=%s&client_secret=%s&refresh_token=%s&scope=advcampaigns statistics deeplink_generator", c.cfg.ClientId, c.cfg.ClientSecret, c.refreshAccessToken)
	}
	return fmt.Sprintf("grant_type=client_credentials&client_id=%s&scope=advcampaigns statistics deeplink_generator", c.cfg.ClientId)
}

func parseApiError(statusCode int, data []byte) *exerror.ExtendedError {
	apiError := new(models.ApiError)
	err := json.Unmarshal(data, apiError)
	if err != nil {
		return exerror.New(ErrInvalidEntity, exerror.Important(), exerror.Message(fmt.Sprintf("%s can not be casted to ApiError type", string(data))))
	}

	switch statusCode {
	case http.StatusBadRequest:
		return exerror.New(ErrBadRequest, exerror.Important(), exerror.Message(apiError.Description))
	case http.StatusNotFound:
		return exerror.New(ErrNotFound, exerror.Message(apiError.Error))
	case http.StatusUnauthorized:
		return exerror.New(ErrUnauthorized, exerror.Temporary(), exerror.Message(apiError.Description))
	case http.StatusForbidden:
		return exerror.New(ErrNotEnoughRights, exerror.Message(apiError.Description))
	case http.StatusInternalServerError:
		return exerror.New(ErrInternal, exerror.Important(), exerror.Message(apiError.Description))
	default:
		return exerror.New(ErrUnknown, exerror.Important(), exerror.Message(apiError.Description))
	}
}
