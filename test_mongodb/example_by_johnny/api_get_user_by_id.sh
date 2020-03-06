#!bin/bash

echo "[Johnny]: 請輸入你需要查找的ID, please input the id you want to query !"
echo " all url example:"
echo "     http://localhost:12345/user/5e5fcfdb5465a6ac83dffa6b"
echo 
read query_id

dynamic_url="http://localhost:12345/user/${query_id}"
echo "-----------------------------"
echo $dynamic_url
echo "-----------------------------"
echo 
curl -v --header "Content-Type: application/json" --request GET $dynamic_url
