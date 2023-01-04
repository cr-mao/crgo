package goods

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"crgo/infra/db"
	"crgo/models"
)

//轮播图
func (s *GoodsService) BannerList(ctx context.Context, req *emptypb.Empty) (*BannerListResponse, error) {
	bannerListResponse := BannerListResponse{}

	var banners []models.Banner
	result := db.GetDb("goods").Find(&banners)
	bannerListResponse.Total = int32(result.RowsAffected)

	var bannerReponses []*BannerResponse
	for _, banner := range banners {
		bannerReponses = append(bannerReponses, &BannerResponse{
			Id:    banner.ID,
			Image: banner.Image,
			Index: banner.Index,
			Url:   banner.Url,
		})
	}

	bannerListResponse.Data = bannerReponses

	return &bannerListResponse, nil
}

func (s *GoodsService) CreateBanner(ctx context.Context, req *BannerRequest) (*BannerResponse, error) {
	banner := models.Banner{}
	banner.Image = req.Image
	banner.Index = req.Index
	banner.Url = req.Url
	db.GetDb("goods").Save(&banner)
	return &BannerResponse{Id: banner.ID}, nil
}

func (s *GoodsService) DeleteBanner(ctx context.Context, req *BannerRequest) (*emptypb.Empty, error) {
	if result := db.GetDb("goods").Delete(&models.Banner{}, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "轮播图不存在")
	}
	return &emptypb.Empty{}, nil
}

func (s *GoodsService) UpdateBanner(ctx context.Context, req *BannerRequest) (*emptypb.Empty, error) {
	var banner models.Banner
	if result := db.GetDb("goods").First(&banner, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "轮播图不存在")
	}
	if req.Url != "" {
		banner.Url = req.Url
	}
	if req.Image != "" {
		banner.Image = req.Image
	}
	if req.Index != 0 {
		banner.Index = req.Index
	}
	db.GetDb("goods").Save(&banner)
	return &emptypb.Empty{}, nil
}
