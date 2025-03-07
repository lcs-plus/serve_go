CREATE TABLE `apps`
(
    `id`                int(11) NOT NULL AUTO_INCREMENT,
    `name`              varchar(50) NOT NULL,
    `config`            text        NOT NULL,
    `users`             text,
    `brand`             char(50)    NOT NULL,
    `status`            tinyint(4) NOT NULL DEFAULT '1',
    `is_plan`           tinyint(1) DEFAULT '1',
    `uid`               int(11) NOT NULL DEFAULT '0',
    `created_at`        datetime    NOT NULL,
    `updated_at`        datetime    NOT NULL,
    `sort`              int(11) DEFAULT NULL,
    `ua_config`         text,
    `is_upload_cookies` tinyint(1) NOT NULL DEFAULT '0',
    `base_currency`     char(5) DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `name` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=134 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='应用';



CREATE TABLE `reports_sales_daily`
(
    `id`                int(11) NOT NULL AUTO_INCREMENT,
    `date`                                 date NOT NULL COMMENT '统计日期',
    `pc_ga_sessions`                       int(10) unsigned DEFAULT NULL COMMENT 'pc端流量',
    `pc_orders_num`                        int(10) unsigned DEFAULT NULL COMMENT 'pc端订单数',
    `pc_orders_fee`                        decimal(10, 2) unsigned DEFAULT NULL COMMENT 'pc端订单金额',
    `pc_ad_fb`                             decimal(10, 2) unsigned DEFAULT NULL COMMENT 'pc端facebook广告费用',
    `pc_ad_google`                         decimal(10, 2) unsigned DEFAULT NULL COMMENT 'pc端谷歌花费',
    `pc_stripe`                            int(10) unsigned DEFAULT NULL,
    `pc_paypal_express`                    int(10) unsigned DEFAULT NULL,
    `pc_adyen_cc`                          int(11) DEFAULT NULL,
    `mb_ga_sessions`                       int(10) unsigned DEFAULT NULL COMMENT '移动端流量',
    `mb_orders_num`                        int(10) unsigned DEFAULT NULL COMMENT '移动端订单数',
    `mb_orders_fee`                        decimal(10, 2) unsigned DEFAULT NULL COMMENT '移动端订单金额',
    `mb_ad_fb`                             decimal(10, 2) unsigned DEFAULT NULL COMMENT '移动端facebook广告费用',
    `mb_ad_google`                         decimal(10, 2) unsigned DEFAULT NULL COMMENT '移动端谷歌广告',
    `mb_stripe`                            int(10) unsigned DEFAULT NULL,
    `mb_paypal_express`                    int(10) unsigned DEFAULT NULL,
    `mb_adyen_cc`                          int(11) DEFAULT '0',
    `sales_num`                            int(10) unsigned DEFAULT NULL COMMENT '销售件数（除去礼品/换货)',
    `user_registered`                      int(10) unsigned DEFAULT NULL COMMENT '用户注册数',
    `ga_direct_sessions`                   int(10) unsigned DEFAULT NULL,
    `ga_direct_transactions`               int(10) unsigned DEFAULT NULL,
    `ga_direct_revenue`                    decimal(10, 2) unsigned DEFAULT NULL,
    `ga_direct_conversion_rate`            decimal(10, 4) unsigned DEFAULT NULL,
    `ga_direct_completions`                int(10) unsigned DEFAULT NULL,
    `ga_new_visitor`                       int(10) unsigned DEFAULT NULL,
    `ga_returning_visitor`                 int(10) unsigned DEFAULT NULL,
    `ga_new_transactions`                  int(10) unsigned DEFAULT NULL COMMENT '新客订单数(ga)',
    `ga_transactions`                      int(10) unsigned DEFAULT NULL COMMENT '总订单数(ga)',
    `ga_fb_new_transactions`               int(10) unsigned DEFAULT NULL COMMENT 'fb新客订单数',
    `ga_fb_transactions`                   int(10) unsigned DEFAULT NULL COMMENT 'fb总订单数',
    `ga_google_new_transactions`           int(10) unsigned DEFAULT NULL COMMENT 'google新客订单数',
    `ga_google_transactions`               int(10) unsigned DEFAULT NULL COMMENT 'google总订单数',
    `ad_ppc`                               decimal(10, 2) unsigned DEFAULT NULL COMMENT 'PPC广告费用（暂时废弃）',
    `ad_criteo`                            decimal(10, 2) unsigned DEFAULT NULL COMMENT 'Criteo广告费用',
    `ad_alliance`                          decimal(10, 2) unsigned DEFAULT NULL COMMENT '联盟广告费用',
    `alliance_order_num`                   int(11) DEFAULT NULL COMMENT '联盟广告订单数',
    `ad_bing`                              decimal(10, 2) unsigned DEFAULT NULL COMMENT 'Bing广告费用',
    `ad_bing_revenue`                      decimal(10, 2) DEFAULT NULL COMMENT 'Bing广告销售额',
    `ad_bing_roas`                         decimal(10, 2) DEFAULT NULL,
    `bing_order_num`                       int(11) DEFAULT NULL COMMENT 'Bing广告订单数',
    `ad_pinterest`                         decimal(10, 2) unsigned DEFAULT NULL COMMENT 'printerest广告费用',
    `ad_outbrain`                          decimal(10, 2) unsigned DEFAULT NULL COMMENT 'outbrain广告花费',
    `ad_twitter`                           decimal(10, 2) unsigned DEFAULT NULL COMMENT 'ad_twitter广告花费',
    `app_id`                               int(10) unsigned NOT NULL,
    `remark`                               text COMMENT '标注',
    `created_at`                           datetime       DEFAULT NULL,
    `updated_at`                           datetime       DEFAULT NULL,
    `ad_tiktok`                            decimal(10, 2) DEFAULT NULL,
    `tiktok_orders_num`                    int(11) DEFAULT NULL,
    `cost_micros`                          decimal(10, 2) DEFAULT NULL COMMENT '谷歌账号金额',
    `send_order_num`                       int(11) DEFAULT '0' COMMENT '发货订单数量',
    `send_orders_fee`                      decimal(10, 2) DEFAULT '0.00' COMMENT '发货订单金额',
    `ad_line`                              decimal(10, 2) DEFAULT '0.00',
    `ga_line_transactions`                 int(11) DEFAULT NULL COMMENT 'line订单',
    `ga_line_trans_revenue`                decimal(10, 2) DEFAULT NULL COMMENT 'line获取的收益',
    `conversions_value_by_conversion_date` decimal(10, 2) DEFAULT '0.00' COMMENT 'google的转化价值（按转化时间）',
    `conversions_value`                    decimal(10, 2) DEFAULT '0.00' COMMENT 'google的转化价值（按点击时间）',
    `fb_income`                            decimal(10, 2) DEFAULT '0.00',
    `applovin_orders_num`                  int(11) DEFAULT NULL,
    `applovin_orders_fee`                  decimal(10, 2) DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='销售日报';