# CLI Chainer :chains:
# :sparkles: Overview
CLIChainer is a versatile command-line interface tool designed to execute predefined chains of commands. It simplifies the automation of a series of related CLI commands across various tools like Azure CLI, Terraform, and more, offering a customizable approach to command management.

# :rocket: Features
- Automate Command Sequences: Execute interconnected CLI commands with ease.
- Customizable Configurations: Define command chains and variables in config.yaml for various scenarios.
- User-Friendly: Designed for ease of use, regardless of the complexity of the task.
- Extensible: Easily extendable to include new command chains for different tools and services.
# :wrench: Installation
(Provide detailed steps on how to install CLIChainer here.)

# :computer: Usage
To use CLIChainer, run the chain command followed by a specific alias and optional arguments, as defined in the config.yaml file.

Basic Command Structure
```
chain run <alias> [additional arguments...]
```
Example
```
chain run aks_login mySubscription
```
Configuration
CLIChainer uses a config.yaml file to define command chains (aliases) and their respective CLI commands.

# :gear: Sample config.yaml
```
variables:
  subscriptions:
    dev: "subscription-id-dev"
    prod: "subscription-id-prod"

cli_aliases:
  aks_login:
    commands:
      - command: "az"
        args: ["account", "set", "--subscription", "{subscriptions.$arg}"]
      - command: "az"
        args: ["account", "show"]
  Add more aliases as needed...
```

:page_facing_up: License
Distributed under the MIT License. See LICENSE for more information.

Project Link: https://github.com/your_username/CliChainer

