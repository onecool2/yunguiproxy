
curl -XPOST "http://www.yysj.xyz:3000/yungui/city_info" -H 'Content-Type: application/json' -d'
{
	"type":"city_info",
	"id": '${1}',
	"city_code": "'${1}'",
	"city_name": "市辖区'${1}'",
	"status": 1,
	"create_time":123123213,
	"update_time":123132131
}'
