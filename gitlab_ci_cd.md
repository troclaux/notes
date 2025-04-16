
# GitLab CI/CD

> built-in tool in GitLab that automates tasks

- configured in `.gitlab-ci.yml` in the repository's root
- use cases
  - running tests
  - building code
  - deploying to servers

example:

```yaml
stages:
  - install
  - test
  - build

variables:
  NODE_ENV: development

cache:
  key: ${CI_COMMIT_REF_SLUG}
  paths:
    - node_modules/

install_job:
  stage: install
  image: node:18
  script:
    - npm ci

test_job:
  stage: test
  image: node:18
  script:
    - npm test

build_job:
  stage: build
  image: node:18
  script:
    - npm run build
  artifacts:
    paths:
      - dist/
    expire_in: 1 hour
```

- `stages`: defines the order of stages
- `variables`: define environment variables
- `install_job`/`test_job`/`build-job`: a job that runs in the build stage (multiple jobs can run in a single stage)
  - `stage`: define which stage the job belongs to
  - `script`: shell commands that run during the job
  - `only`/`except`: control when jobs run (e.g. only on `main`)
  - `timeout`: max time before job is killed
- `artifacts`: save files between stages (e.g. builds)
- `image`: choose docker image to run the job

- gitlab CI/CD comes with built-in environment variables that can be used when prefixed with `$`:
  - `CI_COMMIT_REF_NAME`: branch or tag name for which project is built (e.g. `main`, `feature-x`)
  - `CI_COMMIT_REF_SLUG`: cleaned-up version of the branch name (e.g. `feature/x` => `feature-x`)
  - `CI_COMMIT_SHA`: full git commit SHA
  - `CI_COMMIT_SHORT_SHA`: short version of the commit SHA
  - `CI_COMMIT_TAG`
  - `CI_PIPELINE_ID`: unique id of the current pipeline
  - `CI_PIPELINE_SOURCE`: what triggered the pipeline (push, merge_request_event, etc)

## cache vs artifacts

- `cache`: caches the `node_modules/` directory
  - uses `CI_COMMIT_REF_SLUG` (i.e. branch name) as unique key, so each branch has its own cache
  - speeds up subsequent pipeline runs because `npm install` will be skipped
  - is shared across jobs and pipelines if `key` matches
  - is not used to pass files between jobs in the same pipeline, that's what `artifacts` are for


```yaml
cache:
  key: $CI_COMMIT_REF_SLUG
  paths:
    - node_modules/
```

- `artifacts`: saves the `dist/` folder after the job finished
  - lets you download these files from gitlab UI
  - allows later jobs to access these files
  - `expire_in`: defines how long gitlab keeps the artifacts
    - after that, the files are deleted to save space

```yaml
artifacts:
  paths:
    - dist/
  expire_in: 1 hour
```

## secrets

1. access gitlab project (e.g. `https://gitlab.com/your-username/your-repo`)
1. navigate to: `Settings` => `CI/CD` => `Variables`
1. click `Add Variable`
1. fill in the fields
1. click `Add variable`
