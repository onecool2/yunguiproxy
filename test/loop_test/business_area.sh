curl -XPOST "http://www.yysj.xyz:3000/yungui/business_area" -H 'Content-Type: application/json' -d'
{
	"type":"business_area",
	"id": '${1}',
	"name": "海淀区'${1}'",
	"province_code" : "110000",
	"city_code": "'${2}'",
	"partner_id": 9,
	"status": 1,
	"create_time":123123213,
	"update_time":123132131
}'
