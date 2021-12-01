# tokenproject

The service which gets assets and returns addresses of vaults and lptokens

## Requirements

* [Docker 20.10.6+](https://www.docker.com/get-started)
* [Compose 3.3+](https://docs.docker.com/compose/install/)
* [Go 1.16+](https://golang.org/)
* [Postgresql 12.6](https://www.postgresql.org/)

## Running the service
#### For development purposes
1. Modify *config.yaml* file with your needs:

    * Provide a database url where the information will be stored in the form as provided [here](https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING)

      ```sh
      db:
         url: "postgresql://[userspec]@[hostspec][/dbname][?paramspec]"
      ```


2.Add environment variable in the run configuration :
* KV_VIPER_FILE=config.yaml *(environment variable)*
5. Run service twice with the following command arguments:

   ```sh
   migrate up
   run service
   ```

# API
To change port, configure
```sh
listener:
  addr: :8090
```
where *8090* is a port to listen on.

#### Endpoints
```sh
/add # add new assets to the database
/list #get list of addresses in the database

```

