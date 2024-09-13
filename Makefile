################################################################################
# Project: coolqr                                                              #
# Filename: /Makefile                                                          #
# Created Date: Friday September 13th 2024 11:28:13 +0800                      #
# Author: Sallehuddin Abdul Latif (salleh@sallehuddin.com)                     #
# Company: Sallehuddin Industries                                              #
# --------------------------------------                                       #
# Last Modified: Friday September 13th 2024 11:32:20 +0800                     #
# Modified By: Sallehuddin Abdul Latif (salleh@sallehuddin.com)                #
# --------------------------------------                                       #
# Copyright (c) 2024 Sallehuddin Industries                                    #
################################################################################

clean:
	rm -rf bin/*

build_cli:
	go build -o bin/qoolqr main.go

build: build_cli

all: clean build