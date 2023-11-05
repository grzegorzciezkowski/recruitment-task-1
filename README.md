# Introduction
As a recruitment task, I developed a service with an endpoint that returns the index of a specified value, along with 
the number of occurrences of that value. 

If the exact value is not found, the service will return the closest match within a 10% tolerance, along with its index. 

If no match is found within the tolerance, the service will indicate that no matching value exists.

I added additionally endpoint to check is the service is working correctly.

# Code Structure
| *Folder name*           | *Description*                                 |
|-------------------------|-----------------------------------------------|
| `cmd`                   | Code with command                             |
| `exampleconfiguration ` | Example configuration                         |
| `internal`              | Internal package code                         |
| `testdata`              | Data with important file for the test purpose |

I didn't use the `pkg` folder from standard code structure because it was not necessary for that task.

# Search algorithm
During program initialization, a file containing numerical data is loaded. This data is stored in a slice. Additionally, 
a map is created to associate each numerical value with its corresponding index, enabling faster search operations when 
searching for specific values.

I implemented `divide and conquer` algorithm. So the average computational complexity is `O(log n)`.

The implementation of the algorithm is in `internal/repositories/numbers/repository.go`.

# Configuration file
As a configuration of the service I decided to use the configuration file, which should be named `config.yaml` and put 
into our working directory.

| *key name*     | *Example*     | *Description*                                                                                                                               |
|----------------|---------------|---------------------------------------------------------------------------------------------------------------------------------------------|
| `port`         | `8080`        | Server port                                                                                                                                 |
| `release_mode` | `true`        | Flag with value if we want to run our service in release mode. More information is available at https://pkg.go.dev/github.com/gin-gonic/gin |
| `input_file`   | `input.txt`   | Path to the input file                                                                                                                      |
| `log_file`     | `logfile.log` | Log file                                                                                                                                    |
| `log_level`    | `debug`       | Log level. Possible values is `debug`, `info`, `error`                                                                                      |

An example file You can find here: `exampleconfiguration/config.yaml`

# Logging
I used multi-writer to keep logs in a file and on console output. I decided to keep logs in JSON format because in that
format logs could be easily store in log storage. For example `Kibana`, `Splunk` etc.
