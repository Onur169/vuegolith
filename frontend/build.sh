#!/bin/bash

function moveBuildOutput {
  rm -R ../ui/dist
  mv dist ../ui/dist
}

if [[ $VITE_ENV == "prod" ]]; then
  vite build --mode prod
  moveBuildOutput
else
  vite build --mode dev
  moveBuildOutput
fi