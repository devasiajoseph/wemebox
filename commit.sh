#!/bin/bash
# Toptech commit script

git add --all
git commit -m "$1"
git push origin main
