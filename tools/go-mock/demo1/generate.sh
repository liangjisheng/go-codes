#!/bin/bash

mockgen -source=db.go -destination=db_mock.go -package=main
