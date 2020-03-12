#!bin/bash

echo "[Jean]: Enter the first_name: ggxxxx"
read first_name
echo 



echo "[Jean]: Enter the family_name:"
read family_name

echo "[Jean]: Enter the email name"
read email

echo
echo "-------"
echo $first_name
echo $email
echo $family_name
echo "-------"
echo 

json_body="{\"first_name\":\"$first_name\",\"family_name\":\"$family_name\",\"email\":\"$email\"}"

curl -v --header "Content-Type: application/json" --request POST --data $json_body  http://localhost:12345/register
