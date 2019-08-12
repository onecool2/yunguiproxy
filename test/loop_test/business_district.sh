curl -XPOST "http://www.yysj.xyz:3000/yungui/business_district" -H 'Content-Type: application/json' -d'
{
	"type":"business_district",
		"id": '${1}',
		"name": "万泉庄'${1}'",
		"business_area_id": '${2}',
		"status": 1,
		"create_time":1563876884,
		"update_time":1563876884
}'
