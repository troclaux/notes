
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
  - set of steps
  - each job runs on a separate runner
- step: single task that can run commands
  - e.g. installing dependencies, running tests
  - execute sequentially within the same job
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

[example of ci workflow](./code/workflows/ci.yml)
[example of cd workflow](./code/workflows/cd.yml)

## keywords

- `name`: name of the workflow
- `on`: contains the events that trigger the workflow
  - `<event_name>`: name of the event that triggers the workflow (e.g. `pull_request`, `push`)
    - `branches: [<branch_name>]`: defines branches where the workflow will run on
- `jobs`: contains jobs that will run in parallel
  - `<job_name>`: job's name
    - `name`: label for a step within a job (used in GitHub Actions' UI)
    - `id`: unique identifier for a step within a job
    - `runs-on`: defines the job's runner (e.g. `ubuntu-latest`)
    - `env`: set environment variables for runner (used for shell commands with `run:`)
      - `DATABASE_URL: ${{ secrets.DATABASE_URL }}`: set environment variable with value from github repository's secrets
    - `steps`: contains all steps in a job
      - `- name`: name of single step
        - `uses`: defines action that will be used, along with the version (e.g. `actions/checkout@v4`)
        - `with`: adds configuration options to an action (input to an action)
          - `<config>`: can be `env`, `go-version`, `working-directory`, etc
        - `run`: runs a shell command


- examples of official actions:
  - `uses: actions/checkout@v4`: copies your github repo into the runner
  - `uses: actions/setup-go@v5`: installs golang into the runner

- a step succeeds when exiting with status code of `0` and fails if it exits with a status code that isn't `0`

- Dynamic badge: show the current status of your CI workflow. They automatically update to reflect:
  - ✅ Green: Passing - all checks succeeded
  - ❌ Red: Failing - one or more checks failed
  - 🟡 Yellow: In progress - workflow is currently running

To add a badge to your README.md:

1. Replace `<OWNER>` with your GitHub username
2. Replace `<REPOSITORY>` with your repository name
3. Replace `<WORKFLOW_FILE>` with your workflow filename (e.g. ci.yml)

- examples:
  - `![CI](https://github.com/johndoe/my-repo/actions/workflows/ci.yml/badge.svg)`
  - `![dynamic badge](https://github.com/<OWNER>/<REPOSITORY>/actions/workflows/<WORKFLOW_FILE>/badge.svg)`
