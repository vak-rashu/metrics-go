# METRICS: STREAMLINE YOUR SYSTEM INSIGHTS

METRICS is a lightweight, terminal-based system monitoring tool written in Go, designed to provide clear and accessible insights into system performance without the complexity of a full monitoring stack.

## Introduction

When working with environments such as **WSL, lightweight servers, or development machines**, traditional monitoring solutions can sometimes feel unnecessarily heavy for the problem at hand. Full observability stacks such as **Prometheus + Grafana** are powerful, but they also introduce additional setup, configuration, and infrastructure.

METRICS takes a simpler approach.

Instead of relying on a large monitoring stack, METRICS reads system statistics directly from **Linux procfs** and presents them through a lightweight terminal UI.

The goal is simple:

> **Understand what is happening on your system without unnecessary complexity.**

### Key Highlights

- **Lightweight:** Runs directly in the terminal with minimal setup.
- **Linux-native:** Reads system statistics directly from `/proc`.
- **Real-time monitoring:** Periodically samples system statistics and updates the TUI.
- **Simple visualization:** Uses terminal-based sparklines to visualize changing metrics.
- **No external monitoring stack:** No Prometheus server, Grafana instance, or exporters required.

---

## Features

### CPU

- Overall CPU utilization
- Uses cumulative CPU counters from `/proc/stat`
- Calculates utilization using counter deltas between samples

### Disk I/O

- Read IOPS
- Write IOPS
- Reads cumulative disk statistics from `/proc/diskstats`

### Network

- Packets received per second
- Packets transmitted per second
- Reads interface statistics from `/proc/net/dev`

### Memory

- Total memory
- Free memory
- Available memory
- Reads current memory statistics from `/proc/meminfo`

---

## How It Works

Linux exposes a large amount of system information through `procfs`.

METRICS uses these interfaces as its data source:

```text
/proc/stat
      │
      └── CPU statistics

/proc/diskstats
      │
      └── Disk I/O statistics

/proc/net/dev
      │
      └── Network statistics

/proc/meminfo
      │
      └── Memory statistics
