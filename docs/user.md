

### create user

    curl -X POST "http://127.0.0.1:8089/smart-edu-server/user" -i -d '{"username":"kenmy","email":"hello@qq.com", "mobile":"1234324"}'
        HTTP/1.1 200 OK
        Content-Type: application/json; charset=utf-8
        Date: Sun, 28 Oct 2018 05:10:52 GMT
        Content-Length: 195

    {
        "code": 200,
        "result": "ok",
        "user": {
            "id": "otohsjux7bnc3bne5nqsettbko",
            "username": "kenmy",
            "role": "",
            "location": "",
            "email": "hello@qq.com",
            "mobile": "1234324",
            "union_id": "",
            "created_time": 1540703872915,
            "deleted_time": 0,
            "updated_time": 1540703872915
        }
    }

### get user

    curl -X GET "http://127.0.0.1:8089/smart-edu-server/user/otohsjux7bnc3bne5nqsettbko" -i -d '{"username":"kenmy","email":"hello@qq.com", "mobile":"1234324"}'
        HTTP/1.1 200 OK
        Content-Type: application/json; charset=utf-8
        Date: Sun, 28 Oct 2018 05:19:03 GMT
        Content-Length: 243

    {
        "code": 200,
        "result": "ok",
        "user": {
            "id": "otohsjux7bnc3bne5nqsettbko",
            "username": "kenmy",
            "role": "",
            "location": "",
            "email": "hello@qq.com",
            "mobile": "1234324",
            "union_id": "",
            "created_time": 1540703872915,
            "deleted_time": 0,
            "updated_time": 1540703872915
        }
    }



