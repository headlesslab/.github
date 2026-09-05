# Security Policy

This is the default security policy for every headlesslab repository that has no `SECURITY.md` of its own.

## Supported versions

Security fixes go to the latest release of each headlesslab module. For wand that is the latest Milestone release (the minor release cut for the current Chrome stable milestone) and, while one is out, its release candidate; for the other modules it is their latest tagged release. Older releases receive nothing.

Fixes ship as patch releases of the supported version.

## Reporting a vulnerability

Report vulnerabilities through GitHub private vulnerability reporting on the affected repository: open the repository's **Security** tab and choose **Report a vulnerability**. Do not open a public issue and do not send exploit details by e-mail.

You will receive an acknowledgement within 7 days of the report.

## Disclosure

Once a fix has shipped, the vulnerability is published as a GitHub Security Advisory (GHSA) on the affected repository, with a CVE requested through GitHub. For Go modules the advisory reaches the Go vulnerability database, so `govulncheck` users learn about it automatically.

## Bounty

headlesslab runs no bug bounty programme and pays no rewards for reports.
