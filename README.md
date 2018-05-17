# medusa
This project is a WIP!

# Setting up your GitHub API key
1. Log into GitHub.
2. Click on the emoji dropdown in the upper right corner and choose "Settings"
3. Click on the Developer Settings link at the bottom of the left hand nav bar.
4. Click on Personal access tokens.
5. Click on Generate new token.
6. Enter a token description
7. Check repo:status, repo_deployment and public_repo under repo, read:org under admin:org, read:user and user:email under user
8. Copy the API key and store it somewhere safe, you will need to copy it to your clipboard when you run Medusa the first time.

TODO:
list of V2 features
medusa invite <user_name>
medusa remove <user_name>
what else???

The goal is to wrap, combine and enhance calls the to the GitHub org REST API:

curl -s -H "Authorization: token API_KEY" https://api.github.com/repos/carbonblack/automation-groovy/collaborators

curl -s -H "Authorization: token API_KEY" https://api.github.com/repos/carbonblack/automation-groovy

for repo in $(curl -s -H "Authorization: token API_KEY" 'https://api.github.com/orgs/carbonblack/repos?type=private&per_page=100' | jq -r '.[].name');do curl -s -H "Authorization: token API_KEY" https://api.github.com/repos/carbonblack/${repo}/collaborators &> ${repo}.txt;done

for repo_file in *.txt;do repo=${repo_file%.txt};echo ${repo};cat ${repo_file} | jq -r '.[] | select(.permissions.admin == true) | .login';echo "";done > repo_admins.out

for repo_file in *.txt;do repo=${repo_file%.txt};echo ${repo};cat ${repo_file} | jq -r '.[] | select(.permissions.admin == true) | .login';echo "";done > repo_admins.out        

curl -s -H "Authorization: token API_KEY" 'https://api.github.com/orgs/carbonblack/members?type=private&per_page=100&page=1&role=admin' | jq '.[]'

curl -s -H "Authorization: token API_KEY" 'https://api.github.com/orgs/carbonblack/members?type=private&per_page=100&role=admin' | jq '.[]'     

curl -s -H "Authorization: token API_KEY" https://api.github.com/orgs/carbonblack/members
