# cloudtesting ![Build Status](https://github.com/corymacd/cloudtesting/actions/workflows/build.yml/badge.svg)

Minimal golang http api for testing of deployment pipelines

## Building and Testing

This project uses GitHub Actions for CI/CD and GoReleaser for building and releasing.

### Development Build
bash
go build -v
### Releases
Releases are automated using GoReleaser. To create a new release:

1. Tag the release:
git tag -a v1.0.0 -m "Release message"
git push origin v1.0.0

This will trigger the GitHub Actions workflow which:
- Runs all tests
- Creates cross-platform builds (Linux, Windows, macOS)
- Generates a GitHub release with artifacts
- Creates release notes automatically

Binaries will be available on the GitHub Releases page.