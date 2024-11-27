
# GitHub Actions

to add CI to a github repository:

1. create `.github` directory in the root of your repository
1. create `workflows` directory inside `.github`
1. create `<workflow_name>.yml` inside of `.github/workflows`
  - examples of workflow names: `ci.yml`, `build.yml`, `deploy.yml`

- good CI pipeline typically includes:
  - unit tests
  - integration tests
  - styling checks
  - linting checks
  - security checks

## jargon

- workflows: starts when an event occurs in github repository
  - made of multiple jobs
- event: activity that triggers a workflow run
  - e.g. pull requests, push
- job: set of steps that run on the same runner
  - in the example below, the workflow contains a single job called "Tests"
  - composed by one or more steps
- step: single task that can run commands
  - e.g. installing dependencies, running tests
  - in the example below, we have 3 steps:
    - set up Golang
    - check out the code
    - force failure of the CI job
- runner: server that runs the workflows

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

## keywords

- `name`: name of the workflow
- `on`: contains the events that trigger the workflow
  - `<event_name>`: name of the event that triggers the workflow (e.g. `pull_request`, `push`)
    - `branches: [<branch_name>]`: defines branches where the workflow will run on
- `jobs`: contains jobs that will run
  - `<job_name>`: job's name
    - `name`: job's description (used in GitHub Actions' UI)
    - `runs-on`: defines the job's runner (e.g. `ubuntu-latest`)
    - `steps`: contains all steps in a job
      - `- name`: name of single step
        - `uses`: defines action that will be used, along with the version (e.g. `actions/checkout@v4`)
        - `with`: adds configuration options to an action
          - `<config>`: can be `env`, `go-version`, `working-directory`, etc
        - `run`: runs a shell command

- a step succeeds when exiting with status code of `0` and fails if it exits with a status code that isn't `0`
