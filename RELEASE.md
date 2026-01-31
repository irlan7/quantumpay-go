# Release Process

This document describes the release and versioning process for the Quantumpay protocol.

Quantumpay prioritizes network stability, determinism, and safety over release frequency.

## Release Philosophy

- Releases are **intentional and conservative**
- Stability is preferred over rapid iteration
- Backward compatibility is preserved whenever possible
- Operators are never forced to upgrade

A release does not imply urgency unless explicitly stated.

## Versioning

Quantumpay follows a semantic versioning model:

- **MAJOR**: Protocol-breaking changes (rare)
- **MINOR**: Backward-compatible protocol improvements
- **PATCH**: Bug fixes and non-behavioral changes

Pre-release versions may be labeled as:
- `-alpha`
- `-beta`
- `-rc`

## Release Types

### Patch Release
- Bug fixes
- Performance optimizations
- No protocol rule changes

### Minor Release
- Backward-compatible protocol improvements
- Optional features
- No forced network upgrades

### Major Release
- Protocol rule changes
- Requires explicit operator coordination
- Announced well in advance


## Release Process

1. Changes are merged into the main branch
2. Changes are tested and reviewed
3. A release tag is created
4. Release notes are published
5. Operators may upgrade at their discretion

There are no automatic updates.


## Operator Responsibility

- Node operators are responsible for deciding when to upgrade
- Operators should review release notes before upgrading
- Running older versions is allowed unless explicitly deprecated


## Emergency Releases

In rare cases involving critical security issues:
- A patch release may be issued with limited disclosure
- Details may be published after a fix is available
- Operators are encouraged to upgrade promptly

## Release Notes

Each release includes:
- Summary of changes
- Compatibility notes
- Upgrade considerations (if any)

## Summary

Quantumpay releases are designed to be predictable, transparent, and safe.

The goal is to maintain long-term network reliability rather than rapid feature delivery.
