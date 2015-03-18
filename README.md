Http APIs   
1. 地理位置报告接口，返回所有玩家信息和豆子信息

URL: /bean/reportLocation  
Params: id=123&longitude=117.01&latitude=34.1  

Return: 
{
    "errno": 0,
     "errmsg": "",
     "players": [
     {
         "id": 123,
         "longitude": 117.01,
         "latitude": 23.1
     },
     {
         "id": 456,
         "longitude": 117.01,
         "latitude": 23.1
     }
     ],
     "beans": [
     {
         "id": 1,
         "state": 1,
         "longitude": 117.01,
         "latitude": 23.1
     },
     {
         "id": 2,
         "state": 1,
         "longitude": 117.01,
         "latitude": 23.1
     }
     ]
}
