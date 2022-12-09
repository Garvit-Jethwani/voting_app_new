#!/bin/bash

NETLIFY_ACCESS_TOKEN=hB0t9oK07gmj4ea_yLBx6Ey9aI9GFHSfmbsA5FXstAg
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
        "env" : {
            "REACT_APP_BALLOT_ENDPOINT": "test",
            "REACT_APP_ECSERVER_ENDPOINT": "test"
        },
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

setenv_output=$(curl --location --request POST "https://api.netlify.com/api/v1/accounts/${account_id}/env?site_id=${site_id}" \
--header "Authorization: Bearer ${NETLIFY_ACCESS_TOKEN}" \
--header 'Content-Type: application/json' \
--data-raw '[
  {
    "key": "REACT_APP_BALLOT",
    "values": [
      {
        "value": "string"
      }
    ]
  }
]'
)

echo ${setenv_output}

create_deploy_output=$(curl --location --request POST "https://api.netlify.com/api/v1/sites/${site_id}/builds" \
--header "Authorization: Bearer ${NETLIFY_ACCESS_TOKEN}" \
)



echo $create_deploy_output

