## Introduction

METRICS is a lightweight, terminal-based system monitoring tool written in Go, designed to provide clear and accessible insights into system performance without the complexity of a full monitoring stack.

When working with environments such as WSL, lightweight servers, or development machines, traditional monitoring solutions can sometimes feel unnecessarily heavy. Full observability stacks such as Prometheus and Grafana are powerful, but they also introduce additional setup and infrastructure.

METRICS takes a simpler approach by reading system statistics directly from Linux procfs and presenting them through a lightweight terminal UI.

> **Understand what is happening on your system without unnecessary complexity.**

<img width="1280" height="680" alt="rashuRAJESHWARI__metrics2026-09-2619-49-22online-video-cutter com-ezgif com-video-to-gif-converter" src="https://github.com/user-attachments/assets/aa2d6cf0-1d6e-462c-8cee-3caa195a2612" />

---

## Monitoring System Resources

METRICS currently monitors four major areas of system activity:

### CPU

CPU utilization is calculated from the cumulative CPU-time counters exposed through `/proc/stat`. METRICS periodically samples these counters and uses the difference between consecutive snapshots to determine CPU utilization over the sampling interval.

<br>

### Disk I/O

Disk activity is collected from `/proc/diskstats`. Cumulative read and write operation counters are converted into read and write IOPS using periodic snapshots.

<br>

### Network

Network statistics are collected from `/proc/net/dev`. METRICS currently tracks packets received and transmitted by the network interface and converts the cumulative counters into packets-per-second rates.

<br>

### Memory

Memory statistics are read directly from `/proc/meminfo`, providing the current total, free, and available memory on the system.

---
