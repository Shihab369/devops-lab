# Feature Development

Use this process for every new feature or collector.

## 1. Define the Problem

Before writing code, answer:

- What problem are we solving?
- What information do we need?
- Where does Linux expose that information?
- Which existing collector is most similar?
- What should be included?
- What is explicitly out of scope?

## 2. Create an Issue

The issue should contain:

- Objective
- Scope
- Requirements
- Acceptance Criteria
- Out of Scope

Keep the initial feature small.

## 3. Create a Branch

Start from the latest main:

```bash
git checkout main
git pull --ff-only
git checkout -b feat/<feature-name>
```

## 4. Implement

Follow the existing architecture.

Keep changes focused on the feature.

## 5. Add Tests

Add tests beside the implementation:

```text
collector.go
collector_test.go
```

Test parsing, collection logic, and important edge cases.

## 6. Register the Feature

If the collector is a new collector, register it in `cmd/devops-lab/main.go`.

## 7. Run Quality Checks

```bash
gofmt -w .
go test ./...
go vet ./...
git diff --check
```

## 8. Commit

Use a focused commit:

```bash
git add <files>
git commit -m "feat: <description>"
```

## 9. Push

```bash
git push -u origin feat/<feature-name>
```

## 10. Pull Request

Open a PR against `main`.

Wait for CI and review before merging.

## 11. Cleanup

After merge:

```bash
git checkout main
git pull --ff-only
git branch -d feat/<feature-name>
git push origin --delete feat/<feature-name>
```

Verify:

```bash
git status
git branch -vv
```
