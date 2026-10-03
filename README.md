#  METRICS

> **Unified, lightweight system observability right in your terminal. No context-switching, no heavy infrastructure.**

---

METRICS is a lightweight, terminal-based system monitoring tool written in Go, designed to provide clear and accessible insights into system performance without the complexity of a full monitoring stack.

When working with environments such as WSL, lightweight servers, or development machines, traditional monitoring solutions can sometimes feel unnecessarily heavy. Full observability stacks such as Prometheus and Grafana are powerful, but they also introduce additional setup and infrastructure.


**METRICS solves this by unifying resource monitoring, process tracking, systemd unit inspection, and live journal logs into a single, keyboard-driven terminal dashboard.**

> **Understand everything happening on your system in one screen.**

**CPU Window**
<img width="1280" height="680" alt="rashuRAJESHWARI__metrics2026-09-2619-49-22online-video-cutter com-ezgif com-video-to-gif-converter" src="https://github.com/user-attachments/assets/aa2d6cf0-1d6e-462c-8cee-3caa195a2612" />

---

**Watch all your services with their names and logs at one place**
<img width="800" height="425" alt="ezgif com-video-to-gif-converter" src="https://github.com/user-attachments/assets/9fb3a211-1c77-475a-ba33-b95374f4d86c" />

---

## 🚀 Key Capabilities

### 📊 All-in-One Resource Monitoring
METRICS reads directly from Linux `procfs` without external agents, rendering real-time sparkline graphs with minimal CPU overhead:

* **CPU Performance**: Real-time utilization calculated from cumulative CPU-time delta snapshots.
* **Memory Breakdown**: Instant visibility into Total, Free, Used, Available, and Cached RAM.
* **Disk I/O**: Tracks read and write operations converted into live IOPS rates.
* **Network Traffic**: Real-time monitoring of RX (received) and TX (transmitted) packet rates per second.

---

### 🛠️ Unified Systemd & Journalctl Inspector
Stop running `systemctl` and `journalctl` in separate terminal split-panes. The **Services** tab brings system control directly alongside performance metrics:

* **Categorized Views**: Instant status filtering for `Failed`, `Running`, and `Dead` services with live unit counts.
* **Unit Metadata**: View unit state, load state, substate, description, and job IDs.
* **Embedded Journal Logs**: Select any service to immediately inspect its tailing `journalctl -u` logs in the right panel.

---

### ⚙️ Process Explorer & Keyboard Navigation
* Navigate seamlessly using numerical shortcuts (`1`, `2`, `3`) or arrow keys.
* Instant sub-tab filtering with `f` (Failed), `r` (Running), and `d` (Dead).
* Lightweight, single binary built with Go, Bubble Tea, and Lip Gloss.

---

## 📦 Quick Start

```bash
# Clone the repository
git clone https://github.com/vak-rashu/metrics-go.git
cd metrics-go

# Build and run
(If not already installed, you would need this)
```sudo apt-get update
sudo apt-get install libsystemd-dev
```
./build
./metrics-go show
