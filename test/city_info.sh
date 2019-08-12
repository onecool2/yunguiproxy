curl -XPOST "http://www.yysj.xyz:3000/yungui/city_info" -H 'Content-Type: application/json' -d'
{
	"type":"city_info",
	"id": 3,
	"city_code": "110100",
	"city_name": "市辖区",
	"status": 1,
	"create_time":123123213,
	"update_time":123132131
}'

curl -XPOST "http://www.yysj.xyz:3000/yungui/city_info" -H 'Content-Type: application/json' -d'
{
	"type":"city_info",
	"id": 376,
	"city_code": "710000",
	"city_name": "台湾省",
	"status": 1,
	"create_time":123123213,
	"update_time":123132131
}'
