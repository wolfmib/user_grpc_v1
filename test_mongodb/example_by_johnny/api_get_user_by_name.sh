echo "[Johnny]: Please input the firstname you want to query !"
echo " all url example:"
echo "     http://localhost:12345/user/name/henry"
echo "     http://localhost:12345/user/name/jason"
echo 
read query_firstname

dynamic_url="http://localhost:12345/user/name/${query_firstname}"
echo "-----------------------------"
echo $dynamic_url
echo "-----------------------------"
echo 
curl -v --header "Content-Type: application/json" --request GET $dynamic_url
