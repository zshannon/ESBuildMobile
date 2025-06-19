# GitHub Actions Setup

This directory contains two different GitHub Action workflows for building and releasing the ESBuildMobile framework:

## Workflows

### 1. `release.yml` - Simple Date-based Releases
- **Trigger**: Runs on every push to `main` branch
- **Versioning**: Uses date + commit hash format (e.g., `v2024.01.15-abc1234`)
- **Features**:
  - Builds the framework using `make build`
  - Runs tests with `make test`
  - Creates a GitHub release with the built `.xcframework`
  - Simple and straightforward

### 2. `semantic-release.yml` - Semantic Versioning (Advanced)
- **Trigger**: Runs on every push to `main` branch
- **Versioning**: Uses semantic versioning based on commit messages (e.g., `v1.2.3`)
- **Features**:
  - Analyzes commit messages to determine version bump
  - Generates changelogs automatically
  - Only creates releases when there are actual changes
  - Follows conventional commit standards

## Which One Should I Use?

### Use `release.yml` if:
- You want simple, predictable releases
- You don't want to worry about commit message formatting
- You prefer date-based versioning
- You want every push to main to create a release

### Use `semantic-release.yml` if:
- You want proper semantic versioning
- You want automatic changelog generation
- You follow conventional commit messages (feat:, fix:, etc.)
- You only want releases when there are meaningful changes

## Setup Instructions

1. **Choose one workflow** and delete the other (or rename it with `.disabled` extension)

2. **For semantic-release.yml**: Make sure your commit messages follow conventional commits:
   - `feat: add new feature` (minor version bump)
   - `fix: resolve bug` (patch version bump)
   - `feat!: breaking change` or `BREAKING CHANGE:` in body (major version bump)

3. **Permissions**: Make sure your repository has the following permissions:
   - Go to Settings → Actions → General
   - Under "Workflow permissions", select "Read and write permissions"
   - Check "Allow GitHub Actions to create and approve pull requests"

## What Gets Released

Both workflows will:
- Build the `ESBuildMobile.xcframework` using your Makefile
- Create a zip file of the framework
- Upload it as a GitHub release asset
- Include release notes

## Manual Release

You can also trigger releases manually by pushing a tag:

```bash
git tag v1.0.0
git push origin v1.0.0
```

This will trigger the release workflow regardless of which one you choose.