#!/usr/bin/env bash
set -eo pipefail

owner="minchao"
repo="go-ptx"
isSpecChanged=false

function setup() {
    if [[ "${TRAVIS}" == "true" ]]; then
        IFS='/' read -r -a substrings <<< "${TRAVIS_REPO_SLUG}"
        owner="${substrings[0]}"
        repo="${substrings[1]}"

        git config --global user.email "travis@example.org"
        git config --global user.name "Travis CI"
        git remote rm origin
        git remote add origin "https://${GITHUB_TOKEN}@github.com/${owner}/${repo}.git" > /dev/null 2>&1
    fi
}

function check_oas_spec() {
    for spec in ./oas.*.*.json
    do
        IFS='.' read -r -a substrings <<< "${spec}"
        target_folders="${target_folders} ${substrings[2]}"
    done
    target_folders="$(echo -e "${target_folders}" | sed -e 's/^[[:space:]]*//' -e 's/[[:space:]]*$//')"
    export target_folders

    local output
    output=$(git diff origin/master -- "${target_folders}")
    if [[ -n "${output}" ]]; then
        isSpecChanged=true
    fi
}

function git_push() {
    local version
    version=$(printf '%(%Y-%m-%d)T\n' -1)

    new_branch="spec-changes-on-${version}"
    export new_branch
    message="OAS spec changes on the ${version}"
    export message

    git add "${target_folders}" oas.*.*.json
    git checkout -b "${new_branch}"
    git commit -m "${message}"
    git push --quiet origin "${new_branch}"
}

function create_github_pull_request() {
    curl -X POST \
         -H "authorization: Bearer ${GITHUB_TOKEN}" \
         -H "Content-Type: application/json" \
         -d "{\"title\":\"${message}\",\"head\":\"${owner}:${new_branch}\",\"base\":\"master\",\"draft\":true}" \
         "https://api.github.com/repos/${owner}/${repo}/pulls" \
         > /dev/null 2>&1
}

setup
git fetch origin master
check_oas_spec

if [[ ${isSpecChanged} == true ]]; then
    git_push
    create_github_pull_request
fi
