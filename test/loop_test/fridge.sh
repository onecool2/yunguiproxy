curl -XPOST "http://www.yysj.xyz:3000/yungui/fridge" -H 'Content-Type: application/json' -d'
{
	"type":"fridge",
	"id": '${1}',
	"partner_id":9,
	"terminal_no":"155077eacdd8482da90b6658e3ed98cf",
	"name":"大健康展示3-4柜",
	"fridge_no":"XXYTNWGU",
	"business_area_id":'${4}',
	"business_district_id":'${3}',
	"warehouse_id":52,
	"building_id":'${2}',
	"goods_table_id":0,
	"build_goods_table_status":2,
	"build_goods_table_time":0,
	"max_weight":30000.00,
	"address":"address",
	"status":3,
	"create_time":123123213,
	"update_time":123132131,
	"generate_order_interval":1,
	"dot_type":1,
	"grade":"A",
	"user_count":4,
	"has_sell_data":1,
	"dec_weight_error":25.00,
	"inc_weight_error":50.00
}'
