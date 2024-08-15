#!/bin/bash

# This exec script implements
# - the replacement of the original file with the mutation,
# - the execution of all tests originating from the current directory,
# - and the reporting if the mutation was killed.

if [ -z ${MUTATE_CHANGED+x} ]; then echo "MUTATE_CHANGED is not set"; exit 1; fi
if [ -z ${MUTATE_ORIGINAL+x} ]; then echo "MUTATE_ORIGINAL is not set"; exit 1; fi
if [ -z ${MUTATE_PACKAGE+x} ]; then echo "MUTATE_PACKAGE is not set"; exit 1; fi

function clean_up {
	if [ -f $MUTATE_ORIGINAL.tmp ];
	then
		mv $MUTATE_ORIGINAL.tmp $MUTATE_ORIGINAL
	fi
}

function sig_handler {
	clean_up

	exit $GOMUTEST_RESULT
}
trap sig_handler SIGHUP SIGINT SIGTERM

export GOMUTEST_DIFF=$(diff -u $MUTATE_ORIGINAL $MUTATE_CHANGED)

mv $MUTATE_ORIGINAL $MUTATE_ORIGINAL.tmp
cp $MUTATE_CHANGED $MUTATE_ORIGINAL

export MUTATE_TIMEOUT=${MUTATE_TIMEOUT:-10}

if [ -n "$TEST_RECURSIVE" ]; then
	TEST_RECURSIVE="/..."
fi

GOMUTEST_TEST=$(go test -timeout $(printf '%ds' $MUTATE_TIMEOUT) .$TEST_RECURSIVE 2>&1)
export GOMUTEST_RESULT=$?

if [ "$MUTATE_DEBUG" = true ] ; then
	echo "$GOMUTEST_TEST"
fi

clean_up

case $GOMUTEST_RESULT in
0) # tests passed -> FAIL
	echo "$GOMUTEST_DIFF"

	exit 1
	;;
1) # tests failed -> PASS
	if [ "$MUTATE_DEBUG" = true ] ; then
		echo "$GOMUTEST_DIFF"
	fi

	exit 0
	;;
2) # did not compile -> SKIP
	if [ "$MUTATE_VERBOSE" = true ] ; then
		echo "Mutation did not compile"
	fi

	if [ "$MUTATE_DEBUG" = true ] ; then
		echo "$GOMUTEST_DIFF"
	fi

	exit 2
	;;
*) # Unkown exit code -> SKIP
	echo "Unknown exit code"
	echo "$GOMUTEST_DIFF"

	exit $GOMUTEST_RESULT
	;;
esac
