package v1

import (
	"0121_1/app/models/v1/sales"
	"0121_1/global"
	"strings"
)

func GetList(listForm *sales.GetTodayListForm) []sales.GetDateInfo {
	date := strings.Split(listForm.Date, "/")
	model := global.App.DB.Model(sales.SaleDaily{}).Where("date >=? and date<=?", date[0], date[1])

	selectString := "sum(pc_ga_sessions) as pc_ga_sessions,sum(pc_orders_num) as pc_orders_num,sum(pc_orders_fee) as pc_orders_fee,sum(pc_ad_fb) as pc_ad_fb,sum(pc_ad_google) as pc_ad_google,sum(pc_stripe) as pc_stripe,sum(pc_adyen_cc) as pc_adyen_cc,sum(pc_paypal_express) as pc_paypal_express,sum(mb_ga_sessions) as mb_ga_sessions,sum(mb_orders_num) as mb_orders_num,sum(mb_orders_fee) as mb_orders_fee,sum(mb_ad_fb) as mb_ad_fb,sum(mb_ad_google) as mb_ad_google,sum(sales_num) as sales_num,sum(mb_stripe) as mb_stripe,sum(mb_adyen_cc) as mb_adyen_cc,sum(mb_paypal_express) as mb_paypal_express,sum(ad_ppc) as ad_ppc,sum(ad_criteo) as ad_criteo,sum(ad_tiktok) as ad_tiktok,sum(tiktok_orders_num) as tiktok_orders_num,sum(ad_alliance) as ad_alliance,sum(alliance_order_num) as alliance_order_num,sum(ad_bing) as ad_bing,sum(ad_bing_revenue) as ad_bing_revenue,avg(ad_bing_roas) as ad_bing_roas,sum(bing_order_num) as bing_order_num,sum(ad_pinterest) as ad_pinterest,sum(ga_new_visitor) as ga_new_visitor,sum(ga_new_transactions) as ga_new_transactions,sum(ga_transactions) as ga_transactions,sum(ga_fb_new_transactions) as ga_fb_new_transactions,sum(ga_fb_transactions) as ga_fb_transactions,sum(ga_google_new_transactions) as ga_google_new_transactions,sum(ga_google_transactions) as ga_google_transactions,sum(cost_micros) as cost_micros,sum(ad_line) as ad_line,sum(ga_line_transactions) as ga_line_transactions,sum(ga_line_trans_revenue) as ga_line_trans_revenue,sum(conversions_value_by_conversion_date) as conversions_value_by_conversion_date,sum(conversions_value) as conversions_value,sum(fb_income) as fb_income,sum(applovin_orders_num) as applovin_orders_num,sum(applovin_orders_fee) as applovin_orders_fee,"

	switch listForm.DateType {
	case "day":
		selectString += "date as dimensions"
		break
	case "week":
		selectString += "DATE_FORMAT(date,'%x|%v') as dimensions"
		break
	case "month":
		selectString += "DATE_FORMAT(date,'%Y/%m') as dimensions"
		break
	case "quarter":
		selectString += "CONCAT(DATE_FORMAT(date,'%Y'),'-',QUARTER(date)) as dimensions"
		break
	case "year":
		selectString += "year(date) as dimensions"
		break
	}

	if listForm.IsBrand != 1 {
		model.Where("app_id=?", listForm.AppId)
	} else {

	}

	model.Group("dimensions").Select(selectString)
	var list *[]sales.GetDateInfo
	if err := model.Find(&list).Error; err != nil {
		return nil
	}

	lists := sales.ProcessUserData(list)

	return lists
}
