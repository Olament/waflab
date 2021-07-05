# WAFLab-CLI

WAFLab-ClI offers an command line interface for generating testcase and testing Web Application Firwall

## waflab

The entrypoint of the command line interface

## waflab generate

```bash
waflab generate [WAF rule files directory] [output directory]
```

waflab generate will recurisively read ModSecurity configuration file from a directory and then write the generated ftw-compatible yaml testcase to another directory  

### Avaliable flags

#### Count

Specify the number of testcase you would like to generate for each variable in WAF rules

For example, for following rules with 2 variable:

```
SecRule REQUEST_URI|REQUEST_HEADERS "@validateByteRange 32-36,38-126" \
    "id:920272,\
    phase:2,\
    block,\
    t:none,t:urlDecodeUni,\
    setvar:'tx.anomaly_score_pl3=+%{tx.critical_anomaly_score}'"
```

```bash
waflab generate [WAF rule directory] [output directory] -c 1
```

will generate one testcase for each variable

```yaml
- test_title: 920272-0
  desc: REQUEST_URI
  stages:
  - stage:
      input:
        stop_magic: true
        dest_addr: 127.0.0.1
        method: GET
        port: 80
        protocol: http
        uri: /%259A%25C6%251C%25D1%25B2%2596%2506%258A%2580%2587
        version: HTTP/1.0
      output:
        status:
        - 403
- test_title: 920272-1
  desc: REQUEST_HEADERS
  stages:
  - stage:
      input:
        stop_magic: true
        dest_addr: 127.0.0.1
        method: GET
        port: 80
        protocol: http
        uri: /
        version: HTTP/1.0
        headers:
          fzNLAcpn7w: '%9A%C6%1C%D1%B2%96%06%8A%80%87'
      output:
        status:
        - 403
```

#### Seed

You can use seed flags to specify the random seed used for each testcase. The seed flag is useful when debugging or you want your generated testcase to be deterministic each time

Example usage

```bash
waflab generate [WAF rule directory] [output directory] -s 41
```

## waflab test

```bash
waflab test [target WAF address]
```

waflab test use the ftw-compatible tool to send the existing testcase from data source or generated testcase from the ModSecurity configs to the target WAF server. 

Example usage

```bash
waflab test mywaf.com -y [YAML testcase directory]

944250-1 | 403 | 920320 944250 949110
944250-2 | 403 | 920320 944250 949110
944250-3 | 403 | 920320 944250 949110
944250-4 | 403 | 920320 944250 949110
944250-5 | 403 | 920320 944250 949110
944250-6 | 403 | 920320 944250 949110
944250-7 | 403 | 920320 944250 949110
```

Notice that you must specify a data source for testcase either using ```-g``` or ```-y```

### Avaliable flags

#### Config 

Use ModSecurity configuration files as the data source for the test

```bash
waflab test [target WAF address] --config [ModSecurity configuration directory]
waflab test [target WAF address] -g [ModSecurity configuration directory]
```

#### Yaml

Use existing yaml testcase as the data source for the test

```bash
waflab test [target WAF address] --yaml [YAML testcase directory]
waflab test [target WAF address] -y [YAML testcase directory]
```

#### Filter 

Pass a regular expression to filter out hitrules. By default, the filter is set to .*, meaning that we do not filter out anything at all.

For example, for hit-rule output like this

```bash
944250-1 | 403 | PROTOCOL-ENFORCEMENT-920320, JAVA-944250, BLOCKING-EVALUATION-949110
```

we can filter out necessary information with regular expression ```\d{2,}``` and get

```bash
944250-1 | 403 | 920320 944250 949110
```

By default, each matched group will separate by a single space.

#### Format 

format flags allows you to reorganize the ouput result and even omit certain entry you do not interested. 

Reorder the output entry

```bash
waflab test MyWAF.com -y [YAML testcase directory] --format "%NAME | %HIT | %STATUS"

944250-7 | 920320 944250 949110 | 403 
```

Omit certain output entry

```bash
waflab test MyWAF.com -y [YAML testcase directory] --format "%NAME | %STATUS"

944250-7 | 403 
```

Change the entry separator

```bash
waflab test MyWAF.com -y [YAML testcase directory] --format "%NAME,%STATUS,%HIT"

944250-7,403,920320 944250 949110  
```

#### Json

Json flags allows you to read enabled rules in a specified json file and only these rules will be tested

```bash
waflab test [target WAF address] --json [JSON path]
waflab test [target WAF address] -j [JSON path]
```

By default, go "repos/wafrules-drs-2.0.json" to find the json file


#### LOG

LOG flags specify the file path for WAFLab-CLI to read the log files when you supplied your own yaml files that contains the ```log_contains``` and ```no_log_contains``` in the ```output``` section. Notice that since we do not show the result of ```log_contains``` and ```no_log_contains``` by default, you need to specify you own output format 

```bash
waflab test [Target WAF address] -y [YAML testcase directory] --format "%NAME | %HIT | %STATUS | %LOG_MATCH | %NOLG_MATCH" 
```