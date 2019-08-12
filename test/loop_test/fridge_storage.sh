curl -XPOST "http://www.yysj.xyz:3000/yungui/fridge_storage" -H 'Content-Type: application/json' -d'
{
	"type": "fridge_storage",
	"fridge_id": '${1}',
	"partner_id": 123,
	"warehouse_id": 123,
	"goods_id": '${2}',
	"goods_name": "goods_name",
	"goods_code": "goods_code",
	"quantity": -1,
	"weight": 12.2,
	"type":"fridge_storage",
	"create_time":123123213,
	"update_time":123132131
}'
