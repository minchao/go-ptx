#!/usr/bin/env bash
set -eo pipefail

owner="minchao"
repo="go-ptx"

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

function git_diff() {
    local specs
    specs=$(find . -maxdepth 1 -type f -name "oas.*")
    for spec in ${specs}
    do
        IFS='.' read -r -a substrings <<< "${spec}"
        target_folders="${target_folders} ${substrings[2]}"
    done
    export target_folders

    output=$(git diff origin/master -- ${target_folders} README.md)
    export output
}

function git_push() {
    new_branch="spec-changes-on-$(date +"%Y%m%d")"
    message="OAS spec changes on the $(date +"%Y%m%d")"
    git checkout -b "${new_branch}"
    touch bus/a
    echo "a" > bus/a
    git add ${target_folders} README.md
    git commit -m "${message}"
    git push --quiet origin "${new_branch}"
}

function github_draft_pr() {
    curl -X POST \
         -H "authorization: Bearer ${GITHUB_TOKEN}" \
         -H "Content-Type: application/json" \
         -d "{\"title\":\"${message}\",\"head\":\"${owner}:${new_branch}\",\"base\":\"master\",\"draft\":true}" \
         "https://api.github.com/repos/${owner}/${repo}/pulls" \
         > /dev/null 2>&1
}

setup
git fetch origin master
git_diff

if [[ -n "${output}" ]]; then
    git_push
    github_draft_pr
fi
