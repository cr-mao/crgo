package goods

import (
	"context"

	"crgo/infra/db"
	"crgo/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *GoodsService) CategoryBrandList(ctx context.Context, req *CategoryBrandFilterRequest) (*CategoryBrandListResponse, error) {
	var categoryBrands []models.GoodsCategoryBrand
	categoryBrandListResponse := CategoryBrandListResponse{}

	var total int64
	db.GetDb("goods").Model(&models.GoodsCategoryBrand{}).Count(&total)
	categoryBrandListResponse.Total = int32(total)

	db.GetDb("goods").Preload("Category").Preload("Brands").Scopes(Paginate(int(req.Pages), int(req.PagePerNums))).Find(&categoryBrands)

	var categoryResponses []*CategoryBrandResponse
	for _, categoryBrand := range categoryBrands {
		categoryResponses = append(categoryResponses, &CategoryBrandResponse{
			Category: &CategoryInfoResponse{
				Id:             categoryBrand.Category.ID,
				Name:           categoryBrand.Category.Name,
				Level:          categoryBrand.Category.Level,
				IsTab:          categoryBrand.Category.IsTab,
				ParentCategory: categoryBrand.Category.ParentCategoryID,
			},
			Brand: &BrandInfoResponse{
				Id:   categoryBrand.Brands.ID,
				Name: categoryBrand.Brands.Name,
				Logo: categoryBrand.Brands.Logo,
			},
		})
	}

	categoryBrandListResponse.Data = categoryResponses
	return &categoryBrandListResponse, nil
}

func (s *GoodsService) GetCategoryBrandList(ctx context.Context, req *CategoryInfoRequest) (*BrandListResponse, error) {
	brandListResponse := BrandListResponse{}

	var category models.Category
	if result := db.GetDb("goods").Find(&category, req.Id).First(&category); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "商品分类不存在")
	}

	var categoryBrands []models.GoodsCategoryBrand
	if result := db.GetDb("goods").Preload("Brands").Where(&models.GoodsCategoryBrand{CategoryID: req.Id}).Find(&categoryBrands); result.RowsAffected > 0 {
		brandListResponse.Total = int32(result.RowsAffected)
	}

	var brandInfoResponses []*BrandInfoResponse
	for _, categoryBrand := range categoryBrands {
		brandInfoResponses = append(brandInfoResponses, &BrandInfoResponse{
			Id:   categoryBrand.Brands.ID,
			Name: categoryBrand.Brands.Name,
			Logo: categoryBrand.Brands.Logo,
		})
	}

	brandListResponse.Data = brandInfoResponses

	return &brandListResponse, nil
}

func (s *GoodsService) CreateCategoryBrand(ctx context.Context, req *CategoryBrandRequest) (*CategoryBrandResponse, error) {
	var category models.Category
	if result := db.GetDb("goods").First(&category, req.CategoryId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "商品分类不存在")
	}

	var brand models.Brands
	if result := db.GetDb("goods").First(&brand, req.BrandId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
	}

	categoryBrand := models.GoodsCategoryBrand{
		CategoryID: req.CategoryId,
		BrandsID:   req.BrandId,
	}

	db.GetDb("goods").Save(&categoryBrand)
	return &CategoryBrandResponse{Id: categoryBrand.ID}, nil
}

func (s *GoodsService) DeleteCategoryBrand(ctx context.Context, req *CategoryBrandRequest) (*emptypb.Empty, error) {
	if result := db.GetDb("goods").Delete(&models.GoodsCategoryBrand{}, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "品牌分类不存在")
	}
	return &emptypb.Empty{}, nil
}

func (s *GoodsService) UpdateCategoryBrand(ctx context.Context, req *CategoryBrandRequest) (*emptypb.Empty, error) {
	var categoryBrand models.GoodsCategoryBrand

	if result := db.GetDb("goods").First(&categoryBrand, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌分类不存在")
	}

	var category models.Category
	if result := db.GetDb("goods").First(&category, req.CategoryId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "商品分类不存在")
	}

	var brand models.Brands
	if result := db.GetDb("goods").First(&brand, req.BrandId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
	}

	categoryBrand.CategoryID = req.CategoryId
	categoryBrand.BrandsID = req.BrandId

	db.GetDb("goods").Save(&categoryBrand)

	return &emptypb.Empty{}, nil
}
