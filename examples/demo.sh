#!/bin/bash

BASE="http://localhost:8080/game"

echo "=== Creating new game ==="
GAME=$(curl -s -X POST "$BASE/")
echo "$GAME"
echo

ID=$(echo "$GAME" | grep -o '"id":"[^"]*"' | cut -d'"' -f4)

echo "=== Player move: (0,0) ==="
curl -s -X POST "$BASE/$ID" \
  -H "Content-Type: application/json" \
  -d '{"board":[[1,0,0],[0,0,0],[0,0,0]]}'
echo
echo

echo "=== Player move: (2,2) ==="
curl -s -X POST "$BASE/$ID" \
  -H "Content-Type: application/json" \
  -d '{"board":[[1,0,0],[0,2,0],[0,0,1]]}'
echo
echo

echo "=== Invalid move (occupied cell) ==="
curl -s -X POST "$BASE/$ID" \
  -H "Content-Type: application/json" \
  -d '{"board":[[1,0,0],[0,1,0],[0,0,1]]}'
echo