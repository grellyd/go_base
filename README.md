# Grellyd Go Base

This is a base package for any go project which can be used with go-get.

## Priorities:
 - Abstraction of CLI interface and unimportable main.go
 - Subpackages are grouped by dependency.
 - Testing Mock package.

The use of the mocks depends on writing objects with dependencies to use dependency injection, not hardcoded values.

## Sources:

CLI/RUN structure from: https://npf.io/2016/10/reusable-commands/

Package Structure from: https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1
