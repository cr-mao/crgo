package goods

import (
	"context"
	"crgo/infra/db"
	"crgo/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

//品牌
func (s *GoodsService) BrandList(ctx context.Context, req *BrandFilterRequest) (*BrandListResponse, error) {
	brandListResponse := BrandListResponse{}
	var brands []models.Brands
	result := db.GetDb("goods").Scopes(Paginate(int(req.Pages), int(req.PagePerNums))).Find(&brands)
	if result.Error != nil {
		return nil, result.Error
	}
	var total int64
	db.GetDb("goods").Model(models.Brands{}).Count(&total)
	brandListResponse.Total = int32(total)
	var brandResponses []*BrandInfoResponse
	for _, brand := range brands {
		brandResponses = append(brandResponses, &BrandInfoResponse{
			Id:   brand.ID,
			Name: brand.Name,
			Logo: brand.Logo,
		})
	}
	brandListResponse.Data = brandResponses
	return &brandListResponse, nil
}

func (s *GoodsService) CreateBrand(ctx context.Context, req *BrandRequest) (*BrandInfoResponse, error) {
	//新建品牌
	if result := db.GetDb("goods").Where("name=?", req.Name).First(&models.Brands{}); result.RowsAffected == 1 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌已存在")
	}
	brand := &models.Brands{
		Name: req.Name,
		Logo: req.Logo,
	}
	db.GetDb("goods").Save(brand)
	return &BrandInfoResponse{Id: brand.ID}, nil
}

func (s *GoodsService) DeleteBrand(ctx context.Context, req *BrandRequest) (*emptypb.Empty, error) {
	if result := db.GetDb("goods").Delete(&models.Brands{}, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "品牌不存在")
	}
	return &emptypb.Empty{}, nil
}

func (s *GoodsService) UpdateBrand(ctx context.Context, req *BrandRequest) (*emptypb.Empty, error) {
	brands := models.Brands{}
	if result := db.GetDb("goods").First(&brands); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
	}
	if req.Name != "" {
		brands.Name = req.Name
	}
	if req.Logo != "" {
		brands.Logo = req.Logo
	}
	db.GetDb("goods").Save(&brands)
	return &emptypb.Empty{}, nil
}
