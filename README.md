# cli-test

## Installation

```
brew tap lksmsr/cli-test
brew install --cask lksmsr/cli-test/homebrew-cli-test
```

run with
```
homebrew-cli-test
```


## Authentication

The cli expects two environment variables for authentication:

- `SLIPLANE_API_KEY`: Your API key for authenticating requests.
- `SLIPLANE_ORG_ID`: The ID of your Sliplane organization.

You can retrieve both from your Sliplane team settings:
 
https://sliplane.io/app/team/api


## Compile

Run 
```bash
go build -o sliplane-cli
```


## Release

Add Github PAT

```bash
export GITHUB_TOKEN=your_github_pat
```

Tag Release

```bash
git tag v0.1.0
git push origin v0.1.0
```

Run goreleaser

```bash
goreleaser --clean
```