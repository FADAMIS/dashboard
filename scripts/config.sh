#!/bin/bash

echo "Begin initial setup"

read -p "Continue? (y/n) " choice
case "$choice" in 
  y|Y ) echo "yes";;
  n|N ) exit 0;;
  * ) echo "Invalid option - exiting..." && exit 0;;
esac


# env file lines, env write file
function SetupEnvFiles {
    lines=("$@")

    # remove envBackendSample and envBackendFile from array
    unset "lines[0]"
    unset "lines[1]"

    local -n arr_ref=$1
    for line in "${lines[@]}"; do
        key=${line%%=*}
        value=${line#*=}

        arr_ref[$key]=$value
    done

    echo "" > "$2"

    for key in "${!arr_ref[@]}"; do
        read -p "Enter value for $key: " envVar
        arr_ref[$key]=$envVar

        if [ "$key" = "POSTGRES_USER" ] || [ "$key" = "POSTGRES_PASSWORD" ]; then
            envPostgresSample[$key]=$envVar
        fi

        echo "$key=$envVar" >> "$2"
    done
}

envBackendFile="$(dirname $(realpath $0))/../backend/.env"
envPostgresFile="$(dirname $(realpath $0))/../postgres/postgres.env"
envPgAdminFile="$(dirname $(realpath $0))/../postgres/pgadmin.env"

declare -A envBackendSample
declare -A envPostgresSample
declare -A envPgAdminSample

readarray -t envBackendLines < $(dirname $(realpath $0))/../backend/example.env
readarray -t envPostgresLines < $(dirname $(realpath $0))/../postgres/postgres.example.env
readarray -t envPgAdminLines < $(dirname $(realpath $0))/../postgres/pgadmin.example.env

for line in "${envPostgresLines[@]}"; do
    key=${line%%=*}
    value=${line#*=}

    envPostgresSample[$key]=$value
done

SetupEnvFiles envBackendSample "$envBackendFile" "${envBackendLines[@]}"
SetupEnvFiles envPgAdminSample "$envPgAdminFile" "${envPgAdminLines[@]}"

echo "" > "$envPostgresFile"

for key in "${!envPostgresSample[@]}"; do
    echo "$key=${envPostgresSample["$key"]}" >> "$envPostgresFile"
done