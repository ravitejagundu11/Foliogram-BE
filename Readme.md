# Foliogram-BE

This is the backend repository for the Foliogram project, built using the [Beego framework](https://beego.vip/).

---

## 🚀 Getting Started

Follow these instructions to get the project set up and running on your local machine for development and testing.

### 1. Prerequisites

Before you begin, you will need to install **Go** on your system.

* **Install Go:**
    1.  Go to the official download page: [https://go.dev/dl/](https://go.dev/dl/)
    2.  Download the installer for your operating system and follow the installation steps.

### 2. Project Installation

1.  **Clone the Repository:**
    ```sh
    git clone https://github.com/ravitejagundu11/Foliogram-BE.git
    cd Foliogram-BE
    ```

2.  **Install the `bee` CLI Tool:**
    The `bee` tool helps manage, build, and run Beego projects.
    ```sh
    go install github.com/beego/bee/v2@latest
    ```

3.  **Install Project Dependencies:**
    Navigate to the root directory of the project and run the following commands to download and verify the project's dependencies.
    ```sh
    go mod tidy
    go mod vendor -v
    ```

---

### 🔧 Troubleshooting: `bee command not found`

If you run `bee version` and get a "command not found" error (especially on macOS or Linux using Zsh), it means your shell's `PATH` is not configured correctly.

1.  **Add Go's bin directory to your Zsh profile:**
    ```sh
    echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshrc
    ```

2.  **Apply the changes to your current terminal session:**
    ```sh
    source ~/.zshrc
    ```

3.  **Verify the fix:**
    ```sh
    bee version
    ```
    You should now see the `bee` tool's version information.

---

## 🏃 Running the Application

### For Development (with Hot-Reload)

This command will start the server, watch for any file changes, and automatically rebuild and restart the application.

```sh
bee run