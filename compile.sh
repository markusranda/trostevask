#!/bin/bash

echo "Building program.."
cd cmd/trostevask
go build -o ../../trostevask

echo "Finished building!"