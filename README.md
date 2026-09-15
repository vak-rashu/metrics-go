## Introduction

METRICS is a lightweight, terminal-based system monitoring tool written in Go, designed to provide clear and accessible insights into system performance without the complexity of a full monitoring stack.

When working with environments such as WSL, lightweight servers, or development machines, traditional monitoring solutions can sometimes feel unnecessarily heavy. Full observability stacks such as Prometheus and Grafana are powerful, but they also introduce additional setup and infrastructure.

METRICS takes a simpler approach by reading system statistics directly from Linux procfs and presenting them through a lightweight terminal UI.

> **Understand what is happening on your system without unnecessary complexity.**

---

## Monitoring System Resources

METRICS currently monitors four major areas of system activity:

### CPU

CPU utilization is calculated from the cumulative CPU-time counters exposed through `/proc/stat`. METRICS periodically samples these counters and uses the difference between consecutive snapshots to determine CPU utilization over the sampling interval.

<!-- Add CPU image here -->
<img width="390" height="163" alt="Screenshot 2026-09-15 193951" src="https://github.com/user-attachments/assets/fbcbbe6c-b608-4835-b446-c59c23a3e67d" />

<br>

### Disk I/O

Disk activity is collected from `/proc/diskstats`. Cumulative read and write operation counters are converted into read and write IOPS using periodic snapshots.

<!-- Add Disk image here -->
<img width="360" height="337" alt="Screenshot 2026-09-15 193930" src="https://github.com/user-attachments/assets/9157acc9-ffb0-444b-9c71-81feee5d391d" />
<br>

### Network

Network statistics are collected from `/proc/net/dev`. METRICS currently tracks packets received and transmitted by the network interface and converts the cumulative counters into packets-per-second rates.

<!-- Add Network image here -->
<img width="360" height="337" alt="Screenshot 2026-09-15 193918" src="https://github.com/user-attachments/assets/17b7767e-78a4-4510-b716-7c03e9d1d7c0" />
<br>

### Memory

Memory statistics are read directly from `/proc/meminfo`, providing the current total, free, and available memory on the system.

<!-- Add Memory image here -->
<img width="635" height="157" alt="Screenshot 2026-09-13 224118" src="https://github.com/user-attachments/assets/924dc9a7-7765-4e84-8d07-67646d71b253" />
---

## How It Works

Linux exposes system information through several procfs interfaces:

```text
/proc/stat
    └── CPU statistics

/proc/diskstats
    └── Disk I/O statistics

/proc/net/dev
    └── Network statistics

/proc/meminfo
    └── Memory statistics
