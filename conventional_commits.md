# Conventional Commits

## test

indicates any type of creation or modification of test code

Example: Creation of unit tests

Message: `test: add unit tests for user authentication`

## feat

indicates the development of a new feature for the project

Example: Addition of a service, functionality, endpoint, etc

Message: `feat: add user registration endpoint`

## refactor

used when there is code refactoring that has no impact on the system's business logic/rules

Example: Code changes after a code review

Message: `refactor: reorganize authentication middleware`

## style

used when there are formatting and code style changes that do not alter the system in any way

Examples: Changing the style-guide, changing lint convention, fixing indentation, removing whitespace, removing comments

Message: `style: format code according to new style guide`

## fix

used when fixing errors that are generating bugs in the system

Example: Applying treatment for a function that is not having the expected behavior and returning error

Message: `fix: resolve user session timeout issue`

## chore

indicates project changes that do not affect the system or test files. These are development changes

Example: Change eslint rules, add prettier, add more file extensions to .gitignore

Message: `chore: update eslint configuration`

## docs

used when there are changes in project documentation

Example: adding information to API documentation, changing README, etc

Message: `docs: update API authentication documentation`

## build

used to indicate changes that affect the project's build process or external dependencies

Example: Gulp, adding/removing npm dependencies, etc

Message: `build: upgrade dependencies to latest versions`

## perf

indicates a change that improved system performance

Example: changing ForEach to while, improving database query, etc

Message: `perf: optimize database queries in user search`

## ci

used for changes in CI configuration files

Example: Github Actions, etc

Message: `ci: add automated deployment workflow`

## revert

indicates the reversal of a previous commit

Message: `revert: feat: add user registration endpoint`
