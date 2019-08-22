### GIT FLOW

Approach is simple, create branch with issue you are working with, 4eg:

`feature/01/CLI-initialize-project`

If it is `fix` for current code that already on master:

`fix/01/CLI-initialize-project-output`

Do pull request to branch `master`

Master is the developer branch

Releases will lay on certain locked branches, it will occur here after we will be ready with CI/CD pattern.

### Build, deploy, run

- All commands that must be run to fill CI/CD process are described in `Makefile`
- All commands/scripts that are run from host to set up environment are run via `bin/run`

This logic should be sustained to clearify where code should be executed

### Relase
to trigger release type: `source ./.env && git tag -a ${CLI_VERSION} -f`
and then simply `git push origin --tags`