# Security Policy

This is the default security policy for every headlesslab repository that has no `SECURITY.md` of its own.

## Supported versions

For each headlesslab module, the latest Milestone release and, while one is out, its release candidate receive security fixes. A Milestone release is the latest minor release of the module (for wand, the one cut for the current Chrome stable milestone). Older releases receive nothing.

Fixes ship as patch releases of the supported version.

## Reporting a vulnerability

Report vulnerabilities through GitHub private vulnerability reporting on the affected repository: open the repository's **Security** tab and choose **Report a vulnerability**. Do not open a public issue and do not send exploit details by e-mail.

You will receive an acknowledgement within 7 days of the report.

## Disclosure

Once a fix has shipped, the vulnerability is published as a GitHub Security Advisory (GHSA) on the affected repository, with a CVE requested through GitHub. For Go modules the advisory reaches the Go vulnerability database, so `govulncheck` users learn about it automatically.

## Bounty

headlesslab runs no bug bounty programme and pays no rewards for reports.
