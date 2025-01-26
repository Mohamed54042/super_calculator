#!/bin/bash

echo "Running end-to-end tests..."

./supercalculator add 1 2
if [ $? -ne 0 ]; then
  echo "Addition failed!"
  exit 1
fi

./supercalculator div 10 0
if [ $? -eq 0 ]; then
  echo "Division by zero did not fail as expected!"
  exit 1
fi

echo "All tests passed!"
