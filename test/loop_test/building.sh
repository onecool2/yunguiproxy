curl -XPOST "http://www.yysj.xyz:3000/yungui/building" -H 'Content-Type: application/json' -d'
{
	"type":"building",
	"id": '${1}',
	"name": "德为科技园A座'${1}'",
	"business_district_id": '${2}',
	"warehouse_id": 29,
	"status": 1,
	"create_time":1525690444,
	"update_time":1525690444
}'

