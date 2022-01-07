# Contributing Guidelines

svc project is [Apache licensed](LICENSE) and accepts contributions via
GitHub pull requests. This document outlines some of the conventions on
development workflow, commit message formatting, contact points, and other
resources to make it easier to get your contribution accepted.

## Support Channels

The official support channels, for both users and contributors, are:

- GitHub [issues](https://github.com/voi-oss/protoc-gen-event/issues)

## How to Contribute

Pull Requests (PRs) are the main and exclusive way to contribute to the
official `protoc-gen-event` project.

### Setup

[Fork][fork], then clone the repository:

```
git clone git@github.com:your_github_username/protoc-gen-event.git
cd protoc-gen-event
git remote add upstream https://github.com/voi-oss/protoc-gen-event.git
git fetch upstream
```

Install dependencies:

```
go mod vendor
```

Make sure the tests and linters pass:

```
make lint
make tests
```

### Making Changes

Start by creating a new branch for your changes:

```
git checkout master
git fetch upstream
git rebase upstream/master
git checkout -b new-feature
```

Make your changes, then ensure that `make lint` and `make test` still pass. If
you're satisfied with your changes, push them to your fork.

```
git push origin new-feature
```

Then use the GitHub UI to open a pull request.

At this point, you're waiting on us to review your changes. We *try* to respond
to issues and pull requests within a few business days, and we may suggest some
improvements or alternatives. Once your changes are approved, one of the
project maintainers will merge them.

We're much more likely to approve your changes if you:

* Add tests for new functionality.
* Write a [good commit message][commit-message].
* Maintain backward compatibility.

[fork]: https://github.com/voi-oss/protoc-gen-event/fork
[open-issue]: https://github.com/voi-oss/protoc-gen-event/issues/new
[commit-message]: http://tbaggery.com/2008/04/19/a-note-about-git-commit-messages.html
