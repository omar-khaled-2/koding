#!/bin/bash

cd $(dirname $0)

NIGHTWATCH_BIN="../node_modules/.bin/nightwatch"
NIGHTWATCH_OPTIONS="$NIGHTWATCH_OPTIONS --config ../.nightwatch.json"
NIGHTWATCH_CMD="$NIGHTWATCH_BIN $NIGHTWATCH_OPTIONS"

BUILD_DIR=build/lib

REVISION=$(node -e "process.stdout.write(require('../.config.json').rev)")
export REVISION=${REVISION:0:7}

run() {
  SUITE=$1
  SUBSUITE=$2
  TESTNAME=$3

  if [ -z "$SUITE" ]; then
    for i in ./$BUILD_DIR/*;do
      if [ -d "$i" ];then
          if [ "$i" == "./$BUILD_DIR/helpers" ] || [ "$i" == "./$BUILD_DIR/utils" ]; then
            echo "skipping $i"
          else
            echo "running $i test suite"
            $NIGHTWATCH_CMD --group $i
          fi
      fi
    done
  else
    echo "running single test suite: $SUITE $SUBSUITE"

    if [ -z "$SUBSUITE" ]; then
      $NIGHTWATCH_CMD --group ./$BUILD_DIR/$SUITE
    elif [ -z "$TESTNAME" ]; then
      $NIGHTWATCH_CMD --group ./$BUILD_DIR/$SUITE/$SUBSUITE.js
    else
      $NIGHTWATCH_CMD --test ./$BUILD_DIR/$SUITE/$SUBSUITE.js --testcase $TESTNAME
    fi
  fi
}

make compile

run $1 $2 $3

CODE=$?

if [ $CODE -ne 0 ]; then
  mv users.json users-$1-$2.json
fi

exit $CODE
