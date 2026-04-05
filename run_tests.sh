#!/bin/sh

echo "Build app"
go mod download
CGO_ENABLED=0 GOOS=linux go build -o currencyAPI .
echo "Build done"

echo "Test 1: Check ENV"
export VERSION=1.0.0 AUTHOR=meow PORT=9997

./currencyAPI &
sleep 3

CHECK="{\"author\":\"$AUTHOR\",\"service\":\"currency\",\"version\":\"$VERSION\"}"
RESP=$(curl "http://localhost:$PORT/info")

if [ "$RESP" = "$CHECK" ]; then
    echo "Test 1 passed"
else
    echo "Test 1 failed"
    exit 1
fi


echo "Test 2: Check currency endpoint"
CHECK=$(curl http://localhost:$PORT/info/currency)
if echo "$CHECK" | grep "USD"; then
   echo "Test 2 passed"
else
   echo "Test 2 failed"
   exit 1
fi

echo "Test 3: Check currency query parameter"
CHECK=$(curl http://localhost:$PORT/info/currency\?currency=EUR)
if echo "$CHECK" | grep "USD"; then
   echo "Test 3 failed"
   exit 1
else
   if echo "$CHECK" | grep "EUR"; then
     echo "Test 3 passed"
   else
     echo "Test 3 failed"
     exit 1
   fi
fi

echo "Test 4: Check date query parameter"
CHECK=$(curl http://localhost:$PORT/info/currency\?date=2020-02-02\&currency="USD")
if echo $CHECK | grep "63.1385"; then
  echo "Test 4 passed"
else
  echo "Test 4 failed"
  exit 1
fi

echo "======== ALL TESTS PASSED ======="
exit 0
