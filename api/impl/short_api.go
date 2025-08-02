package impl

import (
	"golang.org/x/net/context"
	"goshort/m/v2/api"
	"gorm.io/gorm"
	"github.com/go-redis/redis/v8"
)	

type ShortApiImpl struct{
	db  *gorm.DB
	redis  *redis.Client
}

func (ShortApiImpl) GetListCampaign(ctx context.Context, request api.GetListCampaignRequestObject) (api.
GetListCampaignResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (ShortApiImpl) GetListShort(ctx context.Context, request api.GetListShortRequestObject) (api.GetListShortResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (ShortApiImpl) PostShort(ctx context.Context, request api.PostShortRequestObject) (api.PostShortResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (ShortApiImpl) GetShortCode(ctx context.Context, request api.GetShortCodeRequestObject) (api.GetShortCodeResponseObject, error) {

	return api.GetShortCode302Response{
		Headers: api.GetShortCode302ResponseHeaders{
			Location: "https://google.com",
		},
	}, nil
}

func (ShortApiImpl) GetShortCodeInfo(ctx context.Context, request api.GetShortCodeInfoRequestObject) (api.GetShortCodeInfoResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (ShortApiImpl) PostShortCodeReport(ctx context.Context, request api.PostShortCodeReportRequestObject) (api.PostShortCodeReportResponseObject, error) {
	//TODO implement me
	panic("implement me")
}
