# Pull Requests

Pull requests are used to review and merge changes into `main`.

## Before Opening a PR

Run:

```bash
gofmt -w .
go test ./...
go vet ./...
git diff --check
```

Verify:

```bash
git status
```

Only intended changes should remain.

## PR Title

Use a clear title.

Example:

```text
feat: add network diagnostics
```

## PR Description

A small feature PR should explain:

### What changed?

Describe the implementation.

### Why?

Explain the purpose.

### Validation

List the checks that passed:

```text
go test ./...
go vet ./...
git diff --check
```

## PR Scope

Keep one PR focused on one change.

Avoid unrelated refactoring, formatting changes, or documentation changes unless required by the feature.

## Merge

Merge after:

- CI passes.
- Review is complete.
- Acceptance criteria are satisfied.

After merging, update local `main` and remove the feature branch.
