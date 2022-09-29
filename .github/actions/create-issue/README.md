# Create GitHub Issue - GitHub Action

This action creates a new issue against a GitHub repository.

## Input Parameters

| Parameter   | Description                                         | Required |
| ----------- | --------------------------------------------------- | -------- |
| `token`     | GitHub authentication token                         | `yes`    |
| `title`     | Issue title                                         | `yes`    |
| `body`      | Issue body                                          | `yes`    |
| `assignees` | List of GitHub assignees separated by newlines `\n` | `yes`    |

## Returned Values

| Value   | Description                       |
| ------- | --------------------------------- |
| `issue` | The issue object as a json string |

## Example Usage

```yaml
name: Example workflow

on: [push]

jobs:
  example:
    name: Example issue creation
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v3
      - name: Create GitHub issue
        id: issue
        uses: ./.github/actions/create-issue
        with:
          token: ${{ secrets.GITHUB_TOKEN }}
          title: Issue Title
          body: Issue Body
          assignees: |
            user11
      - run: echo 'Issue ${{ steps.issue.outputs.issue }}'
```

## Building Javascript Actions

[Documentation](https://docs.github.com/en/actions/creating-actions/creating-a-javascript-action)

Checking in your node_modules directory can cause problems. As an alternative, you can use a tool called @vercel/ncc to compile your code and modules into one file used for distribution.

1. Install vercel/ncc by running this command in your terminal. `npm i -g @vercel/ncc`
2. Compile the `index.js` file. `ncc build index.js --license licenses.txt`
3. A new `dist/index.js` file is created with the code and the compiled modules. The accompanying `dist/licenses.txt` file contains the licenses of the node_modules used.
