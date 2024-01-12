#!/bin/bash

# Run initdb.go
go run initdb.go

# Check if initdb.go executed successfully
if [ $? -eq 0 ]; then
    echo "initdb.go executed successfully."

    # Run each seed file in the /seedFiles/setup folder
    for file in seedFiles/setup/*.go; do
        go run "$file"
        if [ $? -eq 0 ]; then
            echo "$file executed successfully."
        else
            echo "Error executing $file."
            exit 1
        fi
    done

    echo "All seed files executed successfully."

else
    echo "Error executing initdb.go."
fi
