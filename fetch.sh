#!/bin/bash
# 远程仓库名称
REMOTE_NAME="upstream"

# 远程仓库 URL
REMOTE_URL="git@github.com:adshao/go-binance.git"

# 检查远程仓库是否存在
if ! git remote show "$REMOTE_NAME" >/dev/null 2>&1; then
  # 添加远程仓库
  git remote add "$REMOTE_NAME" "$REMOTE_URL"
  echo "Remote '$REMOTE_NAME' added."
else
  echo "Remote '$REMOTE_NAME' already exists."
fi

git fetch upstream && git pull --no-ff upstream master
