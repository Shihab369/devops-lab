# Git Workflow

The repository uses short-lived feature branches.

## Start

```bash
git checkout main
git pull --ff-only
git checkout -b feat/<feature-name>
```

## Work

Make focused changes and test them regularly.

## Validate

```bash
gofmt -w .
go test ./...
go vet ./...
git diff --check
```

## Commit

```bash
git add <files>
git commit -m "feat: <description>"
```

## Push

```bash
git push -u origin feat/<feature-name>
```

## Pull Request

Open a PR from the feature branch to `main`.

## After Merge

```bash
git checkout main
git pull --ff-only
```

Delete the local feature branch:

```bash
git branch -d feat/<feature-name>
```

Delete the remote feature branch:

```bash
git push origin --delete feat/<feature-name>
```

Verify:

```bash
git branch -vv
git status
```

Expected final state:

```text
main
up to date with origin/main
nothing to commit, working tree clean
```
