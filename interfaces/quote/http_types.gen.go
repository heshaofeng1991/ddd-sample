// Package interfaces provides primitives to interact with the openapi HTTP API.
//
// Code generated by  version  DO NOT EDIT.
package interfaces

const (
	Jwt_header_AuthorizationScopes = "jwt_header_Authorization.Scopes"
)

// Error defines model for Error.
type Error struct {
	// code (错误码).
	Code int `json:"code"`

	// message (错误信息).
	Message string `json:"message"`
}

// Quote defines model for Quote.
type Quote struct {
	// 实际测量重量(unit g)
	ActualWeight int64 `json:"actual_weight"`

	// 渠道批次ID
	ChannelCostId int64 `json:"channel_cost_id"`

	// 渠道ID
	ChannelId int64 `json:"channel_id"`

	// 渠道名称
	ChannelName string `json:"channel_name"`

	// 渠道物流类型（1 快递 2 专线 3邮政）
	ChannelType int8 `json:"channel_type"`

	// 计费重量(unit g)
	ChargeWeight int64 `json:"charge_weight"`

	// 货币名称
	Currency string `json:"currency"`

	// 渠道其它类型（目的国关税（ddu ddp））
	DeliverDuty string `json:"deliver_duty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 燃油费 (decimal(15,2))
	FuelFee float64 `json:"fuel_fee"`

	// 是否推荐渠道
	IsRecommended bool `json:"is_recommended"`

	// 渠道物流类型（1 Economic 经济 2 Fast 快速 3 Recommend 推荐）
	LogisticsType int8 `json:"logistics_type"`

	// 最大可派达天数
	MaxNormalDays int `json:"max_normal_days"`

	// 最小可派达天数
	MinNormalDays int `json:"min_normal_days"`

	// 渠道杂费 (decimal(15,2))
	MiscFee float64 `json:"misc_fee"`

	// 渠道操作费 (decimal(15,2))
	ProcessingFee float64 `json:"processing_fee"`

	// 渠道挂号费 (decimal(15,2))
	RegistrationFee float64 `json:"registration_fee"`

	// 国外段运费 (decimal(15,2))
	ShippingFee float64 `json:"shipping_fee"`

	// 标签
	Tag *string `json:"tag,omitempty"`

	// 总费用 (decimal(15,2))
	TotalFee float64 `json:"total_fee"`

	// 体积重量(unit g)
	VolumeWeight int64 `json:"volume_weight"`

	// 仓库ID
	WarehouseId int64 `json:"warehouse_id"`
}

// QuoteInfo defines model for QuoteInfo.
type QuoteInfo struct {
	// fees
	List []Quote `json:"list"`
}

// QuoteRsp defines model for QuoteRsp.
type QuoteRsp struct {
	// code
	Code int64      `json:"code"`
	Data *QuoteInfo `json:"data,omitempty"`

	// message
	Message string `json:"message"`
}

// GetParams defines parameters for Get.
type GetParams struct {
	// 出发地国家
	OriginCountry string `json:"origin_country"`

	// 目的地国家
	DestCountry string `json:"dest_country"`

	// 目的地州｜省
	DestState string `json:"dest_state"`

	// 目的地邮政编码
	DestZipCode string `json:"dest_zip_code"`

	// 重量, unit(g)
	Weight int `json:"weight"`

	// 长度, unit(mm)
	Length int `json:"length"`

	// 宽度, unit(mm)
	Width int `json:"width"`

	// 高度, unit(mm)
	Height int `json:"height"`

	// 产品属性
	ProductAttributes *[]string `json:"product_attributes,omitempty"`

	// 预设渠道
	PresetChannelIds *[]int64 `json:"preset_channel_ids,omitempty"`

	// 测试渠道
	TestChannelIds *[]int64 `json:"test_channel_ids,omitempty"`

	// 工厂地址
	Factory *string `json:"factory,omitempty"`

	// 日期
	Date *string `json:"date,omitempty"`

	// 体积
	Volume *int64 `json:"volume,omitempty"`

	// 申报价值
	DeclaredValue *float64 `json:"declared_value,omitempty"`

	// 预扣关税
	PrepayTariff *bool `json:"prepay_tariff,omitempty"`

	// 仓库ID
	WarehouseId *int64 `json:"warehouse_id,omitempty"`

	// 排除渠道
	ExcludeChannelIds *[]int64 `json:"exclude_channel_ids,omitempty"`

	// 结算币种
	SettlementCurrency *GetParamsSettlementCurrency `json:"settlement_currency,omitempty"`
}

// GetParamsSettlementCurrency defines parameters for Get.
type GetParamsSettlementCurrency string
