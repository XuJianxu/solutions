# sql-crc32: A data crc32 checksum tool used for mysql compatible databases

## Installation

```bash
$ go install
```

## Usage

### compare

> Notice: You may need to use `-f` for fast fail.

```bash
$ sql-crc32 compare -h
compare returns the different tables between 2 mysql compatible databases

Usage:
  sql-crc32 compare [flags]

Flags:
  -d, --databases stringArray   crc32 databases
  -D, --downstream string       downstream DSN
  -f, --fast-fail               fast fail when crc32 comparison mismatch
  -h, --help                    help for compare
  -r, --result-file string      result file (default "./result.json")
  -T, --threads int             threads number (default 10)
  -U, --upstream string         upstream DSN
```

### crc32

```bash
$ sql-crc32 crc32 -h
crc32 returns crc32 checksum of databases or tables

Usage:
  sql-crc32 crc32 [flags]

Flags:
  -a, --all                     crc32 all user databases
  -d, --databases stringArray   crc32 databases
  -h, --help                    help for crc32
  -H, --host string             hostname (default "127.0.0.1")
  -p, --password string         password
  -P, --port int                port (default 4000)
  -t, --table string            crc32 tables, example: db.table
  -T, --threads int             threads number (default 10)
  -u, --username string         username (default "root")
```

Example

```bash
$ sql-crc32 compare -U "root:@tcp(127.0.0.1:4000)/" -D "root:@tcp(127.0.0.1:4100)/" -f -d betting_0 | tee compare.log
```

## Performance

* Data in TiDB
    * Size: 1.4TB
    * Table Number: 60,000
    * Data Size Per Table: 24.47GB
* `sync-diff-inspector` verified 11% in a weekend
* `sql-crc32 compare` finished in 44m51s with 10 threads. The performance of `sql-crc32 compare` depends on the upstream
  and downstream databases.

## `sql-crc32 compare` vs `sql-diff-inspector`

| spec                 | sql-crc32   | sync-diff-inspector |
| ---                  | ---         | ---                 |
| speed                | much faster | slow                |
| locate mismatch data | table       | batch rows          | 

## Recommendation

* If the data size is small enough (e.g. <=100MB) and the table number is less than 10, go with `sync-diff-inspector`
* If the data size is large enough or the table number is greater than 10, go with `sql-crc32 compare`
    * If the result is not OK, use `sync-diff-inspector` to check the mismatch tables to locate the mismatch data

## ToDo

* compare
    * Add batch rows crc32 check to improve the capability to locate mismatch data
* crc32
    * improve result output
