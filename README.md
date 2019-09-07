## Getting started

### Development mode
To run project locally type:

`bin/run`

### Test
To run tests without watcher:

`bin/run test`

### Test-watch
To run test watcher type:

`bin/run test-watch`

### Build
To build from source:
`bin/run ci`

After success built files will lay within ./dist/{arch}/{binaryName}

To execute apollo binary file:
try `./apollo` to see all commands that are registered
try `./apollo {cmdName} --help` to see help from command

### Quorum execution
after build
`./apollo apply quorum {values.yml}`

### Providing values
See `example/values.yml` that shows how to attach values to apply execution. 
Since any `values-local.yml` file is gitignored
you should copy `example/values.yml` to `values-local.yml` and provide values that you need.

In test mode cli runs with isolated scope with predefined variables and constants.

> Also see CONTRIBUTING.md

