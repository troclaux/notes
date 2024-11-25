
# GitHub Actions

to add CI to a github repository:

1. create `.github` directory in the root of your repository
1. create `workflows` directory inside .github
1. create `ci.yml` inside of `.github/workflows`

example of workflow:

```yml
name: CI

on:
  pull_request:
    branches: [main]

jobs:
  tests:
    name: Tests
    runs-on: ubuntu-latest

    steps:
      - name: Check out code
        uses: actions/checkout@v4

      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.23.0'

      - name: Force Failure
        run: (exit 1)
```

- workflows: starts when an event occurs in github repository
  - made of multiple jobs
- event: activity that triggers a workflow run
  - e.g. pull requests, push
- job: set of steps that run on the same runner
  - in the example above, the workflow contains a single job called "Tests"
  - composed by one or more steps
  - steps: single task that can run commands
    - e.g. installing dependencies, running tests
    - in the example above, we have 3 steps:
      - set up Golang
      - check out the code
      - force failure of the CI job
- runner: server that runs the workflows

- a step succeeds when exiting with status code of `0` and fails if it exits with a status code that isn't `0`
