# Github runners remove

## Parameters

| Name | Description | Required | Default value |
| ---- | ---- | ---- | ---- |
| access_token | Access token with privileges | Yes | ---- |
| repo_owner | Owner's name of repository | Yes | ---- |
| repo_name | Name of the repository | Yes | ---- |
| days_limit | Execution history retention days | Yes | ---- |

- name: Github runners remove
  uses: matheusrosmaninho/github-runners-remove
  with:
    access_token: "${{ secrets.GITHUB_TOKEN }}"
    repo_owner: "terminalbaka"
    repo_name: "hello-actions"
    days_limit: 30