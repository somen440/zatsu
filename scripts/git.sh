#!/bin/zsh

# 直近の commit のハッシュ値取得
git log --pretty=oneline -n 1 | awk '{print $1}'
