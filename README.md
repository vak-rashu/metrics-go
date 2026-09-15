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
<img width="1252" height="542" alt="Screenshot 2026-09-13 224128" src="https://github.com/user-attachments/assets/5a830995-dbe1-45d3-bb3e-629d0f2194fa" /><img width="635" height="157" alt="Screenshot 2026-09-13 224118" src="https://github.com/user-attachments/assets/a903ff4d-5b26-4bf1-b63a-7f23fb90384c" />



![Uploading Screenshot 2026-09-13 224118.png…]()




<img width="651" height="362" alt="Screenshot 2026-09-13 221123" src="https://github.com/user-attachments/assets/3df7f0cf-ccf1-4df5-aef8-be6aa6080390" />


<img width="651" height="362" alt="Screenshot 2026-09-13 221123" src="https://github.com/user-attachments/assets/1c523f34-4ee5-4394-b68a-11555927cf7d" />
<img width="625" height="547" alt="Screenshot 2026-09-13 221101" src="https://github.com/user-attachments/assets/e8ce93af-e97c-4953-ab13-f5e3931e8790" />

![Uploading Screenshot 2026-09-13 221123.png…]()



<br>

### Disk I/O

Disk activity is collected from `/proc/diskstats`. Cumulative read and write operation counters are converted into read and write IOPS using periodic snapshots.

<!-- Add Disk image here -->
<!-- ![Disk I/O Monitoring](assets/disk.png) -->

<br>

### Network

Network statistics are collected from `/proc/net/dev`. METRICS currently tracks packets received and transmitted by the network interface and converts the cumulative counters into packets-per-second rates.

<!-- Add Network image here -->
<!-- ![Network Monitoring](assets/network.png) -->

<br>

### Memory

Memory statistics are read directly from `/proc/meminfo`, providing the current total, free, and available memory on the system.

<!-- Add Memory image here -->
<!-- ![Memory Monitoring](assets/memory.png) -->

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
