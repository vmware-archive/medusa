# medusa
This project is a WIP!

TODO:
list of V1 features
set up CLI commands
figure out how to handle API key
implement main app code, with tests
add linting, finish up builds
bug fixing and documentation
v2 feature list

The goal is to wrap, combine and enhance calls the to the GitHub org REST API:

curl -s -H "Authorization: token API_KEY" https://api.github.com/repos/carbonblack/automation-groovy/collaborators

curl -s -H "Authorization: token API_KEY" https://api.github.com/repos/carbonblack/automation-groovy

for repo in $(curl -s -H "Authorization: token API_KEY" 'https://api.github.com/orgs/carbonblack/repos?type=private&per_page=100' | jq -r '.[].name');do curl -s -H "Authorization: token API_KEY" https://api.github.com/repos/carbonblack/${repo}/collaborators &> ${repo}.txt;done

for repo_file in *.txt;do repo=${repo_file%.txt};echo ${repo};cat ${repo_file} | jq -r '.[] | select(.permissions.admin == true) | .login';echo "";done > repo_admins.out

for repo_file in *.txt;do repo=${repo_file%.txt};echo ${repo};cat ${repo_file} | jq -r '.[] | select(.permissions.admin == true) | .login';echo "";done > repo_admins.out        

curl -s -H "Authorization: token API_KEY" 'https://api.github.com/orgs/carbonblack/members?type=private&per_page=100&page=1&role=admin' | jq '.[]'

curl -s -H "Authorization: token API_KEY" 'https://api.github.com/orgs/carbonblack/members?type=private&per_page=100&role=admin' | jq '.[]'     

curl -s -H "Authorization: token API_KEY" https://api.github.com/orgs/carbonblack/members
