package sales

import (
	"0121_1/app/models"
	"0121_1/utils"
	"time"
)

type SaleDaily struct {
	models.ID
	Date                             time.Time `json:"date" gorm:"column:date"`                     // 统计日期
	PcGaSessions                     int       `json:"pc_ga_sessions" gorm:"column:pc_ga_sessions"` // pc端流量
	PcOrdersNum                      int       `json:"pc_orders_num" gorm:"column:pc_orders_num"`   // pc端订单数
	PcOrdersFee                      float32   `json:"pc_orders_fee" gorm:"column:pc_orders_fee"`   // pc端订单金额
	PcAdFb                           float32   `json:"pc_ad_fb" gorm:"column:pc_ad_fb"`             // pc端facebook广告费用
	PcAdGoogle                       float32   `json:"pc_ad_google" gorm:"column:pc_ad_google"`     // pc端谷歌花费
	PcStripe                         int       `json:"pc_stripe" gorm:"column:pc_stripe"`
	PcPaypalExpress                  int       `json:"pc_paypal_express" gorm:"column:pc_paypal_express"`
	PcAdyenCc                        int       `json:"pc_adyen_cc" gorm:"column:pc_adyen_cc"`
	MbGaSessions                     int       `json:"mb_ga_sessions" gorm:"column:mb_ga_sessions"` // 移动端流量
	MbOrdersNum                      int       `json:"mb_orders_num" gorm:"column:mb_orders_num"`   // 移动端订单数
	MbOrdersFee                      float32   `json:"mb_orders_fee" gorm:"column:mb_orders_fee"`   // 移动端订单金额
	MbAdFb                           float32   `json:"mb_ad_fb" gorm:"column:mb_ad_fb"`             // 移动端facebook广告费用
	MbAdGoogle                       float32   `json:"mb_ad_google" gorm:"column:mb_ad_google"`     // 移动端谷歌广告
	MbStripe                         int       `json:"mb_stripe" gorm:"column:mb_stripe"`
	MbPaypalExpress                  int       `json:"mb_paypal_express" gorm:"column:mb_paypal_express"`
	MbAdyenCc                        int       `json:"mb_adyen_cc" gorm:"column:mb_adyen_cc"`
	SalesNum                         int       `json:"sales_num" gorm:"column:sales_num"`             // 销售件数（除去礼品/换货)
	UserRegistered                   int       `json:"user_registered" gorm:"column:user_registered"` // 用户注册数
	GaDirectSessions                 int       `json:"ga_direct_sessions" gorm:"column:ga_direct_sessions"`
	GaDirectTransactions             int       `json:"ga_direct_transactions" gorm:"column:ga_direct_transactions"`
	GaDirectRevenue                  float32   `json:"ga_direct_revenue" gorm:"column:ga_direct_revenue"`
	GaDirectConversionRate           float32   `json:"ga_direct_conversion_rate" gorm:"column:ga_direct_conversion_rate"`
	GaDirectCompletions              int       `json:"ga_direct_completions" gorm:"column:ga_direct_completions"`
	GaNewVisitor                     int       `json:"ga_new_visitor" gorm:"column:ga_new_visitor"`
	GaReturningVisitor               int       `json:"ga_returning_visitor" gorm:"column:ga_returning_visitor"`
	GaNewTransactions                int       `json:"ga_new_transactions" gorm:"column:ga_new_transactions"`               // 新客订单数(ga)
	GaTransactions                   int       `json:"ga_transactions" gorm:"column:ga_transactions"`                       // 总订单数(ga)
	GaFbNewTransactions              int       `json:"ga_fb_new_transactions" gorm:"column:ga_fb_new_transactions"`         // fb新客订单数
	GaFbTransactions                 int       `json:"ga_fb_transactions" gorm:"column:ga_fb_transactions"`                 // fb总订单数
	GaGoogleNewTransactions          int       `json:"ga_google_new_transactions" gorm:"column:ga_google_new_transactions"` // google新客订单数
	GaGoogleTransactions             int       `json:"ga_google_transactions" gorm:"column:ga_google_transactions"`         // google总订单数
	AdPpc                            float32   `json:"ad_ppc" gorm:"column:ad_ppc"`                                         // PPC广告费用（暂时废弃）
	AdCriteo                         float32   `json:"ad_criteo" gorm:"column:ad_criteo"`                                   // Criteo广告费用
	AdAlliance                       float32   `json:"ad_alliance" gorm:"column:ad_alliance"`                               // 联盟广告费用
	AllianceOrderNum                 int       `json:"alliance_order_num" gorm:"column:alliance_order_num"`                 // 联盟广告订单数
	AdBing                           float32   `json:"ad_bing" gorm:"column:ad_bing"`                                       // Bing广告费用
	AdBingRevenue                    float64   `json:"ad_bing_revenue" gorm:"column:ad_bing_revenue"`                       // Bing广告销售额
	AdBingRoas                       float64   `json:"ad_bing_roas" gorm:"column:ad_bing_roas"`
	BingOrderNum                     int       `json:"bing_order_num" gorm:"column:bing_order_num"` // Bing广告订单数
	AdPinterest                      float32   `json:"ad_pinterest" gorm:"column:ad_pinterest"`     // printerest广告费用
	AdOutbrain                       float32   `json:"ad_outbrain" gorm:"column:ad_outbrain"`       // outbrain广告花费
	AdTwitter                        float32   `json:"ad_twitter" gorm:"column:ad_twitter"`         // ad_twitter广告花费
	AppID                            int       `json:"app_id" gorm:"column:app_id"`
	Remark                           string    `json:"remark" gorm:"column:remark"` // 标注
	AdTiktok                         float64   `json:"ad_tiktok" gorm:"column:ad_tiktok"`
	TiktokOrdersNum                  int       `json:"tiktok_orders_num" gorm:"column:tiktok_orders_num"`
	CostMicros                       float64   `json:"cost_micros" gorm:"column:cost_micros"`         // 谷歌账号金额
	SendOrderNum                     int       `json:"send_order_num" gorm:"column:send_order_num"`   // 发货订单数量
	SendOrdersFee                    float64   `json:"send_orders_fee" gorm:"column:send_orders_fee"` // 发货订单金额
	AdLine                           float64   `json:"ad_line" gorm:"column:ad_line"`
	GaLineTransactions               int       `json:"ga_line_transactions" gorm:"column:ga_line_transactions"`                                 // line订单
	GaLineTransRevenue               float64   `json:"ga_line_trans_revenue" gorm:"column:ga_line_trans_revenue"`                               // line获取的收益
	ConversionsValueByConversionDate float64   `json:"conversions_value_by_conversion_date" gorm:"column:conversions_value_by_conversion_date"` // google的转化价值（按转化时间）
	ConversionsValue                 float64   `json:"conversions_value" gorm:"column:conversions_value"`                                       // google的转化价值（按点击时间）
	FbIncome                         float64   `json:"fb_income" gorm:"column:fb_income"`
	ApplovinOrdersNum                int       `json:"applovin_orders_num" gorm:"column:applovin_orders_num"`
	ApplovinOrdersFee                float64   `json:"applovin_orders_fee" gorm:"column:applovin_orders_fee"`
	models.Timestamps
	models.SoftDeletes
}

func (s *SaleDaily) TableName() string {
	return "sales_daily"
}

type GetTodayListForm struct {
	AppId    string `json:"app_id" form:"app_id" binding:"required"`
	Date     string `json:"date" form:"date" binding:"required"`
	DateType string `json:"date_type" form:"date_type"`
	IsBrand  int64  `json:"is_brand" form:"is_brand"`
}

type GetDateInfo struct {
	PcGaSessions                     int         `json:"pc_ga_sessions"` // pc端流量
	PcOrdersNum                      int         `json:"pc_orders_num"`  // pc端订单数
	PcOrdersFee                      float32     `json:"pc_orders_fee"`  // pc端订单金额
	PcAdFb                           float32     `json:"pc_ad_fb"`       // pc端facebook广告费用
	PcAdGoogle                       float32     `json:"pc_ad_google"`   // pc端谷歌花费
	PcStripe                         int         `json:"pc_stripe"`
	PcPaypalExpress                  int         `json:"pc_paypal_express"`
	PcAdyenCc                        int         `json:"pc_adyen_cc"`
	MbGaSessions                     int         `json:"mb_ga_sessions"` // 移动端流量
	MbOrdersNum                      int         `json:"mb_orders_num"`  // 移动端订单数
	MbOrdersFee                      float32     `json:"mb_orders_fee"`  // 移动端订单金额
	MbAdFb                           float32     `json:"mb_ad_fb"`       // 移动端facebook广告费用
	MbAdGoogle                       float32     `json:"mb_ad_google"`   // 移动端谷歌广告
	MbStripe                         int         `json:"mb_stripe"`
	MbPaypalExpress                  int         `json:"mb_paypal_express"`
	MbAdyenCc                        int         `json:"mb_adyen_cc"`
	SalesNum                         int         `json:"sales_num"`       // 销售件数（除去礼品/换货)
	UserRegistered                   int         `json:"user_registered"` // 用户注册数
	GaDirectSessions                 int         `json:"ga_direct_sessions"`
	GaDirectTransactions             int         `json:"ga_direct_transactions"`
	GaDirectRevenue                  float32     `json:"ga_direct_revenue"`
	GaDirectConversionRate           float32     `json:"ga_direct_conversion_rate"`
	GaDirectCompletions              int         `json:"ga_direct_completions"`
	GaNewVisitor                     int         `json:"ga_new_visitor"`
	GaReturningVisitor               int         `json:"ga_returning_visitor"`
	GaNewTransactions                int         `json:"ga_new_transactions"`        // 新客订单数(ga)
	GaTransactions                   int         `json:"ga_transactions"`            // 总订单数(ga)
	GaFbNewTransactions              int         `json:"ga_fb_new_transactions"`     // fb新客订单数
	GaFbTransactions                 int         `json:"ga_fb_transactions"`         // fb总订单数
	GaGoogleNewTransactions          int         `json:"ga_google_new_transactions"` // google新客订单数
	GaGoogleTransactions             int         `json:"ga_google_transactions"`     // google总订单数
	AdPpc                            float32     `json:"ad_ppc"`                     // PPC广告费用（暂时废弃）
	AdCriteo                         float32     `json:"ad_criteo"`                  // Criteo广告费用
	AdAlliance                       float32     `json:"ad_alliance"`                // 联盟广告费用
	AllianceOrderNum                 int         `json:"alliance_order_num"`         // 联盟广告订单数
	AdBing                           float32     `json:"ad_bing"`                    // Bing广告费用
	AdBingRevenue                    float64     `json:"ad_bing_revenue"`            // Bing广告销售额
	AdBingRoas                       float64     `json:"ad_bing_roas"`
	BingOrderNum                     int         `json:"bing_order_num"` // Bing广告订单数
	AdPinterest                      float32     `json:"ad_pinterest"`   // printerest广告费用
	AdOutbrain                       float32     `json:"ad_outbrain"`    // outbrain广告花费
	AdTwitter                        float32     `json:"ad_twitter"`     // ad_twitter广告花费
	AppID                            int         `json:"app_id"`
	Remark                           string      `json:"remark"` // 标注
	AdTiktok                         float64     `json:"ad_tiktok"`
	TiktokOrdersNum                  int         `json:"tiktok_orders_num"`
	CostMicros                       float64     `json:"cost_micros"`     // 谷歌账号金额
	SendOrderNum                     int         `json:"send_order_num"`  // 发货订单数量
	SendOrdersFee                    float64     `json:"send_orders_fee"` // 发货订单金额
	AdLine                           float64     `json:"ad_line"`
	GaLineTransactions               int         `json:"ga_line_transactions"`                 // line订单
	GaLineTransRevenue               float64     `json:"ga_line_trans_revenue"`                // line获取的收益
	ConversionsValueByConversionDate float64     `json:"conversions_value_by_conversion_date"` // google的转化价值（按转化时间）
	ConversionsValue                 float64     `json:"conversions_value"`                    // google的转化价值（按点击时间）
	FbIncome                         float64     `json:"fb_income"`
	ApplovinOrdersNum                int         `json:"applovin_orders_num"`
	ApplovinOrdersFee                float64     `json:"applovin_orders_fee"`
	Dimensions                       string      `json:"dimensions"`
	PcKeDanjia                       interface{} `json:"pc_ke_danjia"`
}

func ProcessUserData(sales *[]GetDateInfo) []GetDateInfo {

	var userData []GetDateInfo
	for _, d := range *sales {
		d.PcKeDanjia, _ = utils.Divide(d.PcOrdersFee, d.PcOrdersNum, false, 2)

		userData = append(userData, d)
	}
	return userData
}
