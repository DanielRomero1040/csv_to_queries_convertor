# CSV Convertor

## CSV to postgresql queries


## Description

This app translate a csv file to a posgrest sql queries file.

## Folders
* input
    - paste the csv file inside with the name sample.csv
* output
    - the output file is txt format and take the table name to show it in file name.
    ```
    create_and_inserts_table_table_name
    ```


## Tables csv input

| table_name        |               |
| -------------     |:-------------:|
| field_name        | field_name    |
| field_dataType    | field_dataType|
| field_default     | field_default |
| data              | data          |

## txt output

```
CREATE TABLE table_name (
	id serial NOT NULL,
	field_name field_dataType field_default,
	field_name field_dataType field_default
);

INSERT INTO table_name (field_name,field_name) VALUES ('data','data');
```

## Run it! 

To run it use this commands:
````
go run main.go

or

go build main.go
start main.exe
````