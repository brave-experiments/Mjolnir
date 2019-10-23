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

This logic should be sustained to clarify where code should be executed

### IDE Goland
To forward go dependencies from container to your host write `go mod vendor` within container

### Terraform Code:
Never use "`" (&#96;) sign in terraform code.
We pass it to build as static asset, so this sign will be removed from whole string

All variables that are strictly used like ```"${var.something}"``` must be declared as variable in TF recipe.
Otherwise validator will fail

Do it like:
```
variable "something" {
    description = "variable description"
    default = "default value"
}

```