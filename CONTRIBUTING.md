<!--
SPDX-FileCopyrightText: 2025 James Pond <james@cipher.host>

SPDX-License-Identifier: CC0-1.0
-->

# Contribution guidelines

Thank you for your interest in improving `taskemon`! The project was put
together quickly in a couple of hours and needs a lot of improvements.
If you want to help, follow these guidelines to contribute effectively
and get your patch or PR accepted.

## Communication channels

I'm using both [Sourcehut](https://sr.ht/~jamesponddotco/taskemon) and
[GitHub](https://github.com/jamesponddotco/taskemon) to manage the
project, so choose your preferred communication channel.

- **Mailing list**: For submitting patches, use [the mailing
  list](https://lists.sr.ht/~jamesponddotco/taskemon-devel).
- **PRs:** For submitting pull requests, use
  [GitHub](https://github.com/jamesponddotco/taskemon/pulls).

## Submission guidelines

Adhere to the following rules when submitting your patch or PR:

- **Keep it small**: Avoid changing too many things at once.
- **Individual patches and PRs**: One patch or PR per issue, please.
- **Commit messages**: Take a moment to write meaningful commit
  messages. [Drew DeVault's post on
  this](https://drewdevault.com/2019/02/25/Using-git-with-discipline.html#the-basics-writing-good-commit-messages)
  is a good read.
- **Quality assurance**: [Review your spelling and
  grammar](https://languagetool.org/).
- **Documentation**: Update or write documentation as needed.
- **Coding style**: Maintain the style of the codebase in your
  contributions.
- **Dependencies**: Avoid adding new dependencies to the project. If you
  still need to add one and have a good reason to, don't forget to
  update [LICENSE-3rdparty.csv](LICENSE-3rdparty.csv).
- **Changelog**: Update the `Unreleased` section of
  [CHANGELOG.md](CHANGELOG.md).

Remember to sign-off on your commits by running `git commit --signoff`
before pushing. To understand what this means, read [the Linux Kernel
Developer's Certificate of
Origin](https://www.kernel.org/doc/html/latest/process/submitting-patches.html#sign-your-work-the-developer-s-certificate-of-origin).
