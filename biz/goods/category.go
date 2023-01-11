package goods

import (
	"context"
	"crgo/biz/common"
	"crgo/infra/db"
	"crgo/models"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/protobuf/types/known/emptypb"
)

////商品分类
func (s *GoodsService) GetAllCategorysList(context.Context, *emptypb.Empty) (*CategoryListResponse, error) {
	/*
		[
			{
				"id":xxx,
				"name":"",
				"level":1,
				"is_tab":false,
				"parent":13xxx,
				"sub_category":[
					"id":xxx,
					"name":"",
					"level":1,
					"is_tab":false,
					"sub_category":[]
				]
			}
		]
	*/
	var categorys []models.Category
	db.GetDb("goods").Where(&models.Category{Level: 1}).Preload("SubCategory.SubCategory").Find(&categorys)
	b, _ := json.Marshal(&categorys)
	return &CategoryListResponse{JsonData: string(b)}, nil
}

////获取子分类
func (s *GoodsService) GetSubCategory(ctx context.Context, req *CategoryListRequest) (*SubCategoryListResponse, error) {
	categoryListResponse := SubCategoryListResponse{}

	var category models.Category
	if result := db.GetDb("goods").First(&category, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品分类不存在")
	}

	categoryListResponse.Info = &CategoryInfoResponse{
		Id:             category.ID,
		Name:           category.Name,
		Level:          category.Level,
		IsTab:          category.IsTab,
		ParentCategory: category.ParentCategoryID,
	}

	var subCategorys []models.Category
	var subCategoryResponse []*CategoryInfoResponse
	//preloads := "SubCategory"
	//if category.Level == 1 {
	//	preloads = "SubCategory.SubCategory"
	//}
	db.GetDb("goods").Where(&models.Category{ParentCategoryID: req.Id}).Find(&subCategorys)
	for _, subCategory := range subCategorys {
		subCategoryResponse = append(subCategoryResponse, &CategoryInfoResponse{
			Id:             subCategory.ID,
			Name:           subCategory.Name,
			Level:          subCategory.Level,
			IsTab:          subCategory.IsTab,
			ParentCategory: subCategory.ParentCategoryID,
		})
	}
	categoryListResponse.SubCategorys = subCategoryResponse
	return &categoryListResponse, nil
}
func (s *GoodsService) CreateCategory(ctx context.Context, req *CategoryInfoRequest) (*CategoryInfoResponse, error) {
	category := models.Category{}
	cMap := map[string]interface{}{}
	cMap["name"] = req.Name
	cMap["level"] = req.Level
	cMap["is_tab"] = req.IsTab
	if req.Level != 1 {
		//去查询父类目是否存在
		cMap["parent_category_id"] = req.ParentCategory
	}
	tx := db.GetDb("goods").Model(&models.Category{}).Create(cMap)
	fmt.Println(tx)
	return &CategoryInfoResponse{Id: int32(category.ID)}, nil
}

func (s *GoodsService) DeleteCategory(ctx context.Context, req *DeleteCategoryRequest) (*emptypb.Empty, error) {
	if result := db.GetDb("goods").Delete(&models.Category{}, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品分类不存在")
	}
	return &emptypb.Empty{}, nil
}

func (s *GoodsService) UpdateCategory(ctx context.Context, req *CategoryInfoRequest) (*common.Empty2, error) {
	var category models.Category

	if result := db.GetDb("goods").First(&category, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品分类不存在")
	}

	if req.Name != "" {
		category.Name = req.Name
	}
	if req.ParentCategory != 0 {
		category.ParentCategoryID = req.ParentCategory
	}
	if req.Level != 0 {
		category.Level = req.Level
	}
	if req.IsTab {
		category.IsTab = req.IsTab
	}

	db.GetDb("goods").Save(&category)

	return &common.Empty2{}, nil
}
