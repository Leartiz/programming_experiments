#!/bin/bash

BUILD_DIR=build_manual

mkdir -p ${BUILD_DIR}
cd ${BUILD_DIR} || exit

cmake ..
make -j8