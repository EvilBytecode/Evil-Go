# Evil-Go
<a href="https://t.me/pulzetools"><img src="https://img.shields.io/badge/Join%20my%20Telegram%20group-2CA5E0?style=for-the-badge&logo=telegram&labelColor=db44ad&color=5e2775"></a>

- A malicious Golang Package

## Introduction

Evil-Go is a sample Golang package designed to illustrate potentially malicious behavior, akin to "evil-pip" in Python. It demonstrates how a package could execute commands that may harm or compromise a system.

---

## Functionality

The package includes the `utilities` module, which contains a function named `ExecuteCommand`. Depending on the operating system it detects, this function performs different actions:

- **Windows:** Initiates a new `cmd.exe` process to execute commands.
- **Linux:** Prints a message indicating command execution and then exits.
- **Unsupported Operating Systems:** Notifies that the system is not supported and then exits.

This package intentionally demonstrates unsafe practices and potential security risks associated with executing system commands based on system detection within a programming environment. It serves as a cautionary example rather than a recommended practice in software development.

Always exercise caution and adhere to best practices when handling command execution or system interactions in any software project.

### Why I Made This?
- Educational Purposes
- Encourages reviewing source code before execution, promoting best practices.

> :warning: THIS CAN BE VERY DANGEROUS IF USED FOR MALICIOUS PURPOESES


## License
This project is licensed under the MIT License. See the LICENSE file for details.