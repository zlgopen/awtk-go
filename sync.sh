#!/bin/bash

source env.sh
cd awtk
./gen.sh
go build
cd -

