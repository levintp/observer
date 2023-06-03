# Observer

Monitoring system for cluster nodes.

## Configuration

Observer is configured via a configuration file and environment variables.

### Configuration file

The configuration file is located by default at `/etc/observer/observer.yaml`.

The file is written (as the name suggests) in YAML format, defined by the
scheme described below.

Scheme legend and generic placeholders are defined as follows:

- Parameters in `[ brackets ]` are optional.
- Parameters that start with ` - ` are lists.
- Parameters that ends with `...` are collections (nested objects / lists) of
arbitrary length.
- Default values are specified as `... | default = <default_value>`
- Parameters that contain only optional sub-parameters are also optional and
can be ommited.
- `<string>`: an arbitrary regular string.
- `<host>`: a valid string consisting of a hostname of a machine accessible by
the controller.
- `<port>`: an arbitrary integer number between `1-65535` representing a TCP
port.
- `<category>`: a string that matches the regular expression
`[a-zA-Z_][a-zA-Z0-9_]*` and represents a category of nodes.
- `<streamname>`: a string that matches the regular expression `[a-zA-Z_]+`
and represents a name of a stream.
- `<metricname>`: a string that matches the regular expression `[a-zA-Z_]+`
and represents a name of a metric.
- `<thresholdname>`: a string that matches the regular expression `[a-zA-Z_]+`
and represents a name of a threshold.
- `<modulename>`: a string that matches the regular expression `[a-zA-Z_]+`
and represents a name of a module.
- `<argumentname>`: a string that matches the regular expression
`[a-zA-Z_][a-zA-Z0-9_]*` and represents a name of an argument for a specific
module.
- `<argumentvalue>`: an arbitrary string.
- `<jsonpath>`: an arbitrary `JSONPath` expression.
- `<duration>`: an integer duration in seconds.
- `ALL`: magic value that is defined as all possible options of a specific
setting.
- `INHERIT`: magic value that is defined as the inheritted value of the same
option as is set on the parent item.

Any other placeholder is specified separately.

The scheme of the configuration file is as follows:

```yaml
# Definition of node categories.
categories:
  [ - <category> ... ]

# Configuration and specification for the controller.
controller:
  <controller_spec> 

# Configuration and specification for the database.
database:
  <database_spec> 

# Configuration and specification for the local agent.
agent:
  [ <agent_spec> ]

# Configuration and specification for the REST API.
api:
  <api_spec>

# Definition of stream specifications.
streams:
  [ <streamname>: <stream_spec> ... ]

# Definition of node specifications.
nodes:
  [ <host>: <node_spec> ... ]
```

#### `<controller_spec>`

A `controller_spec` section specifies a set of configuration options regarding
the controller component, specifically connection settings between the
controller and other components.

```yaml
# The hostname of the controller.
host: <host>

# The RPC endpoint port of the controller.
[ port: <port> | default = 1139 ]
```

#### `<database_spec>`

A `database_spec` section specifies a set of configuration options regarding
the elasticsearch database that the system depends on, specifically
connection settings and indices use by the Observer system.

```yaml
# The hostname of the database.
host: <host>

# The connection port of the database.
[ port: <port> | default = 9200 ]

# The user used to authenticate with the database.
user: <string>
 
# The password used to authenticate with the database.
# Note: Writing passwords in plaintext in configuration files is not secure.
#       Using the environment variable for this setting is advised.
pass: <string>

# The elasticsearch index to write metrics data to.
index: <string>
```

#### `<agent_spec>`

An `agent_spec` section specifies a set of configuration options regarding the
local agent on each node, specifically connection settings.

```yaml
# The RPC endpoint port of the local agent.
[ port: <port> | default = 1016]
```

#### `<api_spec>`

An `api_spec` section specifies a set of configuration options regarding the
REST API, specifically connection settings.

Note: The REST API component has 2 endpoints:
- The RPC endpoint, defined by the `port` option, serves as an internal
endpoint for communication between components of the system.
- The REST API endpoint, defiend by the `rest_port` option, serves as as an
exported API for users to control and query the system.

```yaml
# The hostname of the REST API instance.
host: <host>

# The RPC endpoint port of the REST API.
[ port: <port> | default = 1086]

# The REST API endpoint port of the REST API.
[ rest_port: <port> | default = 1291 ]
```

#### `<stream_spec>`

A `stream_spec` section specifies a set of configuration options that defines
a stream.

```yaml
# The node categories that are required to run the stream.
# Note: Each category that is specified in this option MUST match exaclty one
#       category as defined in the root `categories` field.
categories:
  [ - <category> ... | default = ALL ]

# The metrics that are collected by the stream.
metrics:
  - <metricname>: <metric_spec> ...
```

#### `<node_spec>`

A `node_spec` section specifies a set of configuration options that defines a
node.

```yaml
# The categories the node is a part of.
# Note: Each category that is specified in this option MUST match exaclty one
#       category as defined in the root `categories` field.
categories:
  [ - <category> ... | default = ALL]
```

#### `<metric_spec>`

A `metric_spec` section specifies a set of configuration options that defines
a metric in a stream.

```yaml
# The interval at which the metric is collected.
[ interval: <duration> | default = 60 ]

# The node categories that are required to collect the metric.
categories:
  [ - <category> ... | default = INHERIT ]

# The module used to collect data by the metric on collection time.
module:
  <module_spec>

# A set of thresholds on the metric to trigger actions on the local node.
thresholds:
  [ - <thresholdname>: <threshold_spec> ... ]
```

#### `<threshold_spec>`

A `threshold_spec` section specifies a set of configuration options that
defines a threshold on a metric.

```yaml
# The interval at which the threshold is checked.
[ interval: <duration> | default = 60 ]

# The node categories that are required to check the threshold.
categories:
  [ - <category> ... | default = INHERIT ]

# The threshold expressed as a JSONPath expression.
expression: <jsonpath>

# The module used as an action by the threshold when the threshold expression
# is met.
module:
  <module_spec>
```

#### `<module_spec>`

A `module_spec` section specifies a set of configuration options that defines
an arbitrary module, can be both a collection module and an action module.

```yaml
# The name of the module.
# Note: This name MUST match exactly one existing module.
name: <modulename>

# The timeout duration of the module execution. After this amount of time the
# execution of the module will halt and will be declared as failed.
[ timeout: <duration> | default = 59]

# The arguments given to the module.
arguments:
  [ - <argumentname>: <argumentvalue> ... ]
```
