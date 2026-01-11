# Quantumpay Supported Platforms

This document defines the supported platforms for running Quantumpay nodes.

The goal is to ensure predictable behavior, operational stability, and
reproducible deployments across the network.

## Supported Operating Systems

Quantumpay nodes are officially supported on the following operating systems:

### Linux (Primary Support)

- Ubuntu LTS (20.04, 22.04 or later)
- Debian (stable releases)
- Other modern Linux distributions may work but are not officially supported

Linux is the **reference platform** for Quantumpay node operation.

## Unsupported Operating Systems

The following platforms are **not supported** for production node operation:

- Windows (native)
- macOS
- Mobile operating systems
- Embedded or IoT operating systems

These platforms may be used for development or testing only.

## CPU Architecture

Supported CPU architectures:

- `amd64` (x86_64)

Other architectures (ARM, RISC-V, etc.) are not officially supported at this time.

## Containerized Environments

Quantumpay nodes may be deployed in:

- Bare metal servers
- Virtual machines
- Containerized environments (Docker, Kubernetes)

Containerization must not alter:
- Network behavior
- Timing assumptions
- Storage durability

Operators are responsible for ensuring container stability.

## Cloud Providers

Quantumpay is cloud-agnostic.

Nodes may run on:
- Public cloud providers
- Private clouds
- On-premise infrastructure

There are no provider-specific dependencies.

## Development Environments

For development and testing purposes only:

- Local Linux environments
- WSL (Windows Subsystem for Linux)
- CI systems

Development environments must not be used for production mainnet nodes.

## Hardware Requirements (Minimum)

Recommended minimum hardware for a production node:

- 2 CPU cores
- 4 GB RAM
- SSD storage
- Reliable network connection

Higher specifications improve stability but are not mandatory.


## Filesystem Requirements

- POSIX-compliant filesystem
- Stable disk performance
- No network-mounted storage for production nodes

## Change Policy

- Platform support changes are conservative
- New platforms may be added over time
- Platform deprecation will be announced in advance
- No platform changes without governance approval

## Summary

Quantumpay platform support prioritizes predictability over breadth.

Limiting supported platforms reduces operational risk and improves network reliability.
