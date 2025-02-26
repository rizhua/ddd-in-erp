package entity

import (
	"time"

	"rizhua.com/application/command"
	"rizhua.com/application/query"
	"rizhua.com/domain/repository"
	"rizhua.com/infrastructure/persistence/po"
)

var SpuRepo repository.Spu

type mediaVO struct {
	Url      string
	Size     int32
	Width    int
	Height   int
	Name     string
	MimeType string
}

// 库存量单位
type sku struct {
	ID        int64
	SpuID     int64
	Price     int16
	Stock     int16
	Barcode   string
	Discount  int
	StoreID   int32
	Attribute []string
	Media     *mediaVO
}

// 代销商品
type SpuProxyEntity struct {
	ID           int64
	SpuID        int32
	ApplyText    string
	ApplyTime    time.Time
	VerifyText   string
	VerifyTime   time.Time
	VerifyStatus string
	OrgID        int32
}

// 商品评价
type spuComment struct {
	ID         int64
	SpuID      int32
	Content    string
	Media      string
	UserID     int32
	UserAvatar string
	Nickname   string
	LoveCount  int32
	StarCount  int8
	ReadCount  int32
	OrgID      int32
	Status     int8
}

// 根实体：标准化产品单元
type Spu struct {
	ID        int64
	Code      string
	Name      string
	LowPrice  int32
	Brand     *Brand
	SaleCount int32
	RateCount int32
	OrgID     int64
	Barcode   string
	Media     []map[string]any
	Detail    string
	Status    int8
	SkuList   []*sku
	Category  *Category
	Comment   []*spuComment
	UpdateAt  time.Time
	CreateAt  time.Time
}

func (t *Spu) Create(cmd command.CreateSpu) error {
	info := po.Spu{
		Code:       cmd.Code,
		Name:       cmd.Title,
		LowPrice:   cmd.LowPrice,
		Barcode:    cmd.Barcode,
		Detail:     cmd.Detail,
		CategoryID: cmd.CategoryID,
		BrandID:    cmd.BrandID,
		// Media:      cmd.Media,
		Status: cmd.Status,
	}
	return SpuRepo.Create(info)
}

func (t *Spu) Update(cmd command.UpdateSpu) error {
	info := po.Spu{
		ID:       cmd.ID,
		Code:     cmd.Code,
		Name:     cmd.Title,
		LowPrice: cmd.LowPrice,
		Barcode:  cmd.Barcode,
		Detail:   cmd.Detail,
	}
	return SpuRepo.Update(info)
}

func (t *Spu) Delete(id []int64) error {
	return SpuRepo.Delete(id)
}

func (t *Spu) toEntity(in po.Spu) Spu {
	return Spu{
		ID:        in.ID,
		Code:      in.Code,
		Name:      in.Name,
		Brand:     &Brand{ID: in.BrandID},
		LowPrice:  in.LowPrice,
		SaleCount: in.SaleCount,
		RateCount: in.RateCount,
		Barcode:   in.Barcode,
		Category:  &Category{ID: in.CategoryID},
		// Media:     in.Media,
		Status:   in.Status,
		UpdateAt: in.UpdateAt,
		CreateAt: in.CreateAt,
	}
}

func (t *Spu) Get() (info Spu, err error) {
	ret, err := SpuRepo.Get(t.ID)
	if err != nil {
		return
	}
	info = t.toEntity(ret)
	return
}

func (t *Spu) Find(req query.Request) (list []Spu, total int64, err error) {
	ret, total, err := SpuRepo.Find(req)
	for _, item := range ret {
		list = append(list, t.toEntity(item))
	}
	return
}
