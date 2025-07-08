#!/bin/bash

# EROS Backend Microservices Starter
# Her servisi arka planda başlatır ve loglarını ayrı dosyalara yazar
# Kullanım: bash start-backend.sh

SERVICES=(
  "backend/user-service"
  "backend/match-service"
  "backend/chat-service"
  "backend/api-gateway"
)

# Önce eski processleri öldür
for svc in "${SERVICES[@]}"; do
  pkill -f "$svc"
done
sleep 1

# Her servisi başlat
for svc in "${SERVICES[@]}"; do
  echo "[EROS] $svc başlatılıyor..."
  cd "$svc"
  if [ -f main.go ]; then
    nohup go run main.go > ../../${svc##*/}.log 2>&1 &
  else
    echo "main.go bulunamadı: $svc"
  fi
  cd - > /dev/null
  sleep 1
  echo "[EROS] $svc başlatıldı. Log: ${svc##*/}.log"
done

echo "Tüm servisler başlatıldı!" 