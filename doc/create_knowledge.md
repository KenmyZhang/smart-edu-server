### example 1

curl -X POST "http://127.0.0.1:8089/smart-edu-server/knowledge/point" -i -d '{"label":"化学"}'
HTTP/1.1 200 OK
Access-Control-Allow-Origin:
Content-Type: application/json; charset=utf-8
Date: Mon, 31 Dec 2018 03:07:01 GMT
Content-Length: 115

{"code":200,"knowledge_point":{"id":"u9rx3biuftr9nnsjebk6ekzapo","label":"化学","parent_id":"","deleted_time":0}}\

### example 2

curl -X POST "http://127.0.0.1:8089/smart-edu-server/knowledge/point" -i -d '{"label":"化学实验基础", "parent_id":"55cexjwn93bbjxuusf9ctf1c8a"}'
HTTP/1.1 200 OK
Access-Control-Allow-Origin:
Content-Type: application/json; charset=utf-8
Date: Mon, 31 Dec 2018 03:18:00 GMT
Content-Length: 153

{"code":200,"knowledge_point":{"id":"npgj4n4xcpeanetj379iqpuh9h","label":"化学实验基础","parent_id":"55cexjwn93bbjxuusf9ctf1c8a","deleted_time":0}}