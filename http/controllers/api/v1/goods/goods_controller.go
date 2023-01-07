package goods

import (
	"context"
	"crgo/http/util"
	"crgo/infra/errcode"
	"crgo/infra/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	goodsPb "crgo/biz/goods"
	"crgo/http/global"
	"crgo/infra/log"
)

type GoodsController struct {
}

//商品列表
func (cc *GoodsController) List(ctx *gin.Context) {
	//商品的列表 pmin=abc, spring cloud, go-micro
	request := &goodsPb.GoodsFilterRequest{}

	priceMin := ctx.DefaultQuery("pmin", "0")
	priceMinInt, _ := strconv.Atoi(priceMin)
	request.PriceMin = int32(priceMinInt)

	priceMax := ctx.DefaultQuery("pmax", "0")
	priceMaxInt, _ := strconv.Atoi(priceMax)
	request.PriceMax = int32(priceMaxInt)

	isHot := ctx.DefaultQuery("ih", "0")
	if isHot == "1" {
		request.IsHot = true
	}
	isNew := ctx.DefaultQuery("in", "0")
	if isNew == "1" {
		request.IsNew = true
	}

	isTab := ctx.DefaultQuery("it", "0")
	if isTab == "1" {
		request.IsTab = true
	}
	categoryId := ctx.DefaultQuery("c", "0")
	categoryIdInt, _ := strconv.Atoi(categoryId)
	request.TopCategory = int32(categoryIdInt)
	pages := ctx.DefaultQuery("p", "0")
	pagesInt, _ := strconv.Atoi(pages)
	request.Pages = int32(pagesInt)
	perNums := ctx.DefaultQuery("pnum", "0")
	perNumsInt, _ := strconv.Atoi(perNums)
	request.PagePerNums = int32(perNumsInt)
	keywords := ctx.DefaultQuery("q", "")
	request.KeyWords = keywords
	brandId := ctx.DefaultQuery("b", "0")
	brandIdInt, _ := strconv.Atoi(brandId)
	request.Brand = int32(brandIdInt)
	client := goodsPb.NewGoodsClient(global.GrpcConnect)
	r, err := client.GoodsList(context.WithValue(context.Background(), "ginContext", ctx), request)
	if err != nil {
		log.Errorf("get  goodslist grpc error:%v", err)
		util.HandleGrpcErrorToHttp(err, ctx)
		return
	}

	reMap := map[string]interface{}{
		"total": r.Total,
	}
	goodsList := make([]interface{}, 0)
	for _, value := range r.Data {
		goodsList = append(goodsList, map[string]interface{}{
			"id":          value.Id,
			"name":        value.Name,
			"goods_brief": value.GoodsBrief,
			"desc":        value.GoodsDesc,
			"ship_free":   value.ShipFree,
			"images":      value.Images,
			"desc_images": value.DescImages,
			"front_image": value.GoodsFrontImage,
			"shop_price":  value.ShopPrice,
			"category": map[string]interface{}{
				"id":   value.Category.Id,
				"name": value.Category.Name,
			},
			"brand": map[string]interface{}{
				"id":   value.Brand.Id,
				"name": value.Brand.Name,
				"logo": value.Brand.Logo,
			},
			"is_hot":  value.IsHot,
			"is_new":  value.IsNew,
			"on_sale": value.OnSale,
		})
	}
	reMap["data"] = goodsList

	response.Success(ctx, errcode.ErrCodes.ErrNo, reMap)
	return
}

func (cc *GoodsController) New(ctx *gin.Context) {

}

func (cc *GoodsController) Delete(ctx *gin.Context) {

}

//商品详情api 。
func (cc *GoodsController) Detail(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	client := goodsPb.NewGoodsClient(global.GrpcConnect)
	r, err := client.GetGoodsDetail(context.WithValue(context.Background(), "ginContext", ctx), &goodsPb.GoodInfoRequest{
		Id: int32(i),
	})
	if err != nil {
		util.HandleGrpcErrorToHttp(err, ctx)
		return
	}

	rsp := map[string]interface{}{
		"id":          r.Id,
		"name":        r.Name,
		"goods_brief": r.GoodsBrief,
		"desc":        r.GoodsDesc,
		"ship_free":   r.ShipFree,
		"images":      r.Images,
		"desc_images": r.DescImages,
		"front_image": r.GoodsFrontImage,
		"shop_price":  r.ShopPrice,
		"ctegory": map[string]interface{}{
			"id":   r.Category.Id,
			"name": r.Category.Name,
		},
		"brand": map[string]interface{}{
			"id":   r.Brand.Id,
			"name": r.Brand.Name,
			"logo": r.Brand.Logo,
		},
		"is_hot":  r.IsHot,
		"is_new":  r.IsNew,
		"on_sale": r.OnSale,
	}
	ctx.JSON(http.StatusOK, rsp)
}
