#!/bin/bash
# Wemebox commit script

git add --all
git commit -m "$1"
git push origin master

cd sites

git add --all
git commit -m "$1"
git push origin main

