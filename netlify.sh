#!/bin/bash


output=$(curl --location --request POST "https://api.netlify.com/api/v1/sites" \
--header "Authorization: Bearer ${NETLIFY_ACCESS_TOKEN}" \
--header "Content-Type: application/json" \
--data-raw '{
    "repo": {
        "provider": "github",
        "owner_type": "User",
        "installation_id": 32021058,
        "repo": "ishanrai13/voter",
        "private": true,
        "branch": "main",
        "frameworkName": "Create React App",
        "dir": "build",
        "base": "voter",
        "cmd": "npm run build",
        "framework": "create-react-app"
    }
}')



echo $output

validateFile=$(echo $output | jq . > /dev/null)
validateFileOutput=$?
echo $validateFileOutput
if [ $validateFileOutput -eq 4 ]; then
    exit;
fi

site_id=$(echo $output | jq -r '.site_id')

echo $site_id


account_output=$(curl --location --request GET "https://api.netlify.com/api/v1/accounts" \
--header "Authorization: Bearer ${NETLIFY_ACCESS_TOKEN}" \
--header "Content-Type: application/json" \
--data-raw '')


echo $account_output

validateFile=$(echo $account_output | jq . > /dev/null)
validateFileOutput=$?
echo $validateFileOutput
if [ $validateFileOutput -eq 4 ]; then
    exit;
fi

account_id=$(echo $account_output | jq -r '.[0].id')

echo $account_id

EncodedBallotEndpoint=$(echo $ROOST_CLUSTER_IP:30040 | base64 )
EncodedEcserverEndpoint=$(echo $ROOST_CLUSTER_IP:30042 | base64 )

setenv_output=$(curl --location --request POST "https://api.netlify.com/api/v1/accounts/${account_id}/env?site_id=${site_id}" \
--header "Authorization: Bearer ${NETLIFY_ACCESS_TOKEN}" \
--header 'Content-Type: application/json' \
--data-raw "[
  {
    \"key\": \"REACT_APP_BALLOT_ENDPOINT\",
    \"values\": [
      {
        \"value\": \"https://zbio.roost.io/proxy/${EncodedBallotEndpoint}\"
      }
    ]
  },
  {
    \"key\": \"REACT_APP_ECSERVER_ENDPOINT\",
    \"values\": [
      {
        \"value\": \"https://zbio.roost.io/proxy/${EncodedEcserverEndpoint}\"
      }
    ]
  }]"
)

echo ${setenv_output}

create_deploy_output=$(curl --location --request POST "https://api.netlify.com/api/v1/sites/${site_id}/builds" \
--header "Authorization: Bearer ${NETLIFY_ACCESS_TOKEN}" \
)



echo $create_deploy_output

