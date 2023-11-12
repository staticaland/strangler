# Strangler

## Introduction

`strangler` is a Go-based tool designed to incrementally replace functionality in a Bash script, inspired by the [Strangler Fig pattern](https://martinfowler.com/bliki/StranglerFigApplication.html). This approach is particularly useful for transitioning from simple Bash scripts to more robust, scalable Go programs while preserving existing functionality.

## Features

- **Subcommand Flexibility**: Easily define new subcommands in Go to extend script functionality.
- **Integrated Bash Script**: Leverage existing Bash scripts within the Go binary for a seamless transition.
- **Argument Passing**: Pass arguments from Go commands to the Bash script, maintaining script compatibility.

## Installation

### Prerequisites

- Go (version 1.16 or later)
- Bash
- Git (for cloning the repository)

### Steps

1. Clone the repository:

   ```bash
   git clone [repository-url]
   ```

2. Navigate to the project directory:

   ```bash
   cd strangler
   ```

3. Build the project:

   ```bash
   go build -o strangler
   ```

## Usage

Run the `strangler` binary with the desired subcommand. The two initial subcommands are:

- `hello`: Executes the Bash script logic associated with "hello".
- `dawg`: Executes the Bash script logic associated with "dawg".

### Examples

```bash
./strangler hello
./strangler dawg
```
