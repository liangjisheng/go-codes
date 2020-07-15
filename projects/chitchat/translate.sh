#!/bin/bash

# 安装包
# go get -u github.com/nicksnyder/go-i18n/v2/i18n
# go get -u github.com/nicksnyder/go-i18n/v2/goi18n

# 基于 goi18n 命令自动生成翻译文件到 locales 目录(执行前先创建 locales 目录)
goi18n extract -outdir=locales -format=json messages.go
