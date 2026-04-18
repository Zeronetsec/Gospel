<!-- Gospel Project -->

[![version](https://img.shields.io/badge/Gospel-Version%201.0-blue.svg?maxAge=259200)]()
[![gover](https://img.shields.io/badge/Go-Version%201.26.1-blue.svg)]()
[![os](https://img.shields.io/badge/Supported%20OS-Linux-blue.svg)]()
[![license](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![views](https://img.shields.io/endpoint?url=https%3A%2F%2Fhits.dwyl.com%2FZeronetsec%2FGospel.json&&label=Views&color=blue)](https://github.com/Zeronetsec/Gospel)

# Gospel
Gospel is an all-in-one CLI toolkit for system exploration, security auditing, and data inspection. <br>
From decoding obscure strings to uncovering system misconfigurations, Gospel gives you deeper insight into your system.

## Preview
<details>
<summary>Show Preview</summary>
<br>
<img src=".preview/preview_1.png" width="500">
<br><br>
<img src=".preview/preview_2.png" width="500">
<br><br>
<img src=".preview/preview_3.png" width="500">
<br>
</details>

## Features
- Inspect system and hardware information
- Analyze running processes and resource usage
- Detect misconfigurations and potential risks
- Decode and analyze encoded data automatically
- Extract insights from strings (entropy, hashes, metadata)
- And more

## Installation
```bash
git clone https://github.com/Zeronetsec/Gospel.git
cd Gospel
chmod +x install.sh
./install.sh

# for backup
./install.sh --backup
```

## Usage
```bash
gospel --misconfind <path>
gospel --dumpstring <string>
gospel --decode <string|file>
gospel --procinfo
gospel --sysinfo
```
And more commands.

## License
This project is licensed under the MIT License. <br>

<!-- Copyright (c) 2026 Zeronetsec -->