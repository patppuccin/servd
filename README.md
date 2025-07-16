# servd
Cross-platform service management tool

## What this should do?
- allow to register any long running process as a service (also include bat files, sh files, invoked commands like java -jar ...)
- allow to start/stop/restart/uninstall services
- allow to list all services (by status)
- allow to export the list of services
- get the status of the services monitored
- a way to maintain a list of services managed by servd

## Command Structure

servd [command] [options]
├── --help
├── --version
├── --no-interactive
├── register (alias: add) <service-name>
│   # Identity
│   ├── --display-name <display-name>      
│   ├── --description <description>               
│   ├── --tags [list of <tag>]             
│   # Execution
│   ├── --exec <command>                   
│   ├── --dir <working-dir>                
│   ├── --env [list of <key=value>]        
│   ├── --env-files [list of <path>]       
│   # Service Behavior
│   ├── --startup-type (auto|manual)
│   ├── --restart-policy (always|on-failure|unless-stopped)
│   ├── --restart-config <composite-string>
│   └── --start-now (true|false)
├── start <service-name>
├── stop <service-name>
├── restart <service-name>
├── uninstall <service-name>
│   └── --confirm
├── edit <service-name>
│   # Identity
│   ├── --display-name <display-name>      
│   ├── --description <description>               
│   ├── --tags [list of <tag>]             
│   # Execution
│   ├── --exec <command>                   
│   ├── --dir <working-dir>                
│   ├── --env [list of <key=value>]        
│   ├── --env-files [list of <path>]       
│   # Service Behavior
│   ├── --startup-type (auto|manual)
│   ├── --restart-policy (always|on-failure|unless-stopped)
│   ├── --restart-config <composite-string>
│   └── --restart (true|false)
├── list <service-name-pattern|optional>
│   ├── --status (running|stopped|all[default])
│   ├── --format (json|yaml|csv|table[default])
│   ├── --mode (simple[default]|full|detailed)
│   └── --export-path <path>
└── doctor
    ├── --format (json|yaml|csv|list[default])
    ├── --mode (quick[default]|full|detailed)
    └── --export-path <path>

Notes:
- In the register command, if any  mandatory value is omitted, asks interactively
- For start, stop, restart, uninstall, edit commands, if a service name is omitted, asks interactively
- For list command, if a service name pattern is omitted, every service is listed