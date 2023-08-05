#!/bin/bash
build_publish() {
  if [ "$VITE_ENV" = "prod" ]; then
    npm run build-prod
  else
    npm run build-dev
  fi

  cd ../
  go build

  sudo mv vuegolith /usr/local/bin/vuegolith
  sudo cp server.key /etc/vuegolith/ssl/server.key
  sudo cp server.crt /etc/vuegolith/ssl/server.crt
}

build_publish
