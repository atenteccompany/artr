# Contributing to Aten Remote Task Runner

First of all, thank you for taking the time to contribute! We welcome contributions from the community — whether it’s a bug fix, feature proposal, documentation improvement, or just an issue report.

This guide will help you get started and align with the project’s goals and standards.

---

## How to Contribute

There are many ways to help:

- Report a Bug: Please open an issue with a clear title, steps to reproduce, and relevant logs or stack traces if applicable.
- Request a Feature: Use the issue tracker to describe the feature, its use case, and possible implementation ideas.
- Submit Code: Fork the repo, make your changes on a new branch, and submit a Pull Request (PR).
- Improve Documentation: Found something unclear or missing? PRs for README or usage examples are very welcome.

## Project Standards

We expect all contributions to follow these basic principles:

- Keep it simple. Simplicity and maintainability are valued over clever hacks.
- Follow the existing coding style and conventions.
- Write clear, descriptive commit messages.
- All code must be secure by default. Avoid exposing credentials, sensitive paths, or unsafe defaults.
- Include tests or sample usage where relevant.

## Local Development Setup

Here’s a general setup

```bash
# Clone your fork
git clone https://github.com/atenteccompany/artr
cd artr

# Create a new branch for your changes
git checkout -b fix-something

# Make your changes, then commit
git commit -m "Fix: clarify error output from backup script"

# Push and open a Pull Request
git push origin fix-something
```

## Pull Request Checklist

Before submitting a PR, make sure:

- The code builds and passes existing tests.
- The feature or fix is clearly documented.
- No hardcoded credentials, paths, or temporary hacks remain.
- You’ve squashed or cleaned up WIP commits if possible.
- Your PR includes a short description of why the change is useful.

## License

By contributing, you agree that your contributions will be licensed under the MIT License, the same as the project.

## Code of Conduct

All contributors are expected to follow our Code of Conduct. Please read it before participating.
