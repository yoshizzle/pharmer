---
title: Pharmer Use Cluster
menu:
  product_pharmer_0.1.0-alpha.1:
    identifier: pharmer-use-cluster
    name: Pharmer Use Cluster
    parent: reference
product_name: pharmer
menu_name: product_pharmer_0.1.0-alpha.1
section_menu_id: reference
---
## pharmer use cluster

Sets `kubectl` context to given cluster

### Synopsis

Sets `kubectl` context to given cluster

```
pharmer use cluster [flags]
```

### Examples

```
pharmer use cluster <name>
```

### Options

```
  -h, --help        help for cluster
      --overwrite   Overwrite context if found. (default true)
```

### Options inherited from parent commands

```
      --alsologtostderr                  log to standard error as well as files
      --analytics                        Send analytical events to Google Guard (default true)
      --config-file string               Path to Pharmer config file
      --env string                       Environment used to enable debugging (default "prod")
      --log_backtrace_at traceLocation   when logging hits line file:N, emit a stack trace (default :0)
      --log_dir string                   If non-empty, write log files in this directory
      --logtostderr                      log to standard error instead of files (default true)
      --stderrthreshold severity         logs at or above this threshold go to stderr (default 2)
  -v, --v Level                          log level for V logs
      --vmodule moduleSpec               comma-separated list of pattern=N settings for file-filtered logging
```

### SEE ALSO

* [pharmer use](/docs/reference/pharmer_use.md)	 - 

