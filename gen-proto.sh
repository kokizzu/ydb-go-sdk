#!/usr/bin/env bash

##
# This script is intended to generate gRPC and protobuf, as well
# as other `*.go` dependencies to provide otherwise autogenerated
# with `ya make` code for linters, test debuggers and other
# development tools
#
# Has to be called time to time to update generated files with
# their newer versions
#
# Usage: just call this script in a folder of a project you want
# to generate dependencies for. This will also create dependencies
# well outside the project scope
# To undo the generation, call the newly created ungen-proto.sh
#
# Note: it will add all generated files to your `~/.arcignore`,
# creating it if it doesn't exist yet
##

script_path=$(realpath $0)

TMP_DIR=build-protos
ARC_IGNORE=$(realpath ~/.arcignore)

if [ "$ARCADIA_ROOT" = "" ] || ! [ -d "$ARCADIA_ROOT" ]; then
    ARCADIA_ROOT=`arc root`
fi

rm -fr ${TMP_DIR}
mkdir -p ${TMP_DIR}

trap "{ rm -fr ${TMP_DIR}; }" EXIT

ya make \
    --add-result .go \
    --add-protobuf-result \
    --no-src-links \
    --host-platform-flag USE_GCCFILTER="no" \
    --target-platform-flag USE_GCCFILTER="no" \
    --output ${TMP_DIR} --checkout

if ! [ -f "$ARC_IGNORE" ]; then
    echo "" >"$ARC_IGNORE"
fi

ADDED_LINES=0
add_ignore() {
    if grep -q "$1" ${ARC_IGNORE}; then
        return 0
    fi
    if [ $ADDED_LINES -eq 0 ]; then
        cat <<- EOF >>${ARC_IGNORE}

# lines generated by $script_path
EOF
        ADDED_LINES=1
    fi
    cat<<- EOF >>${ARC_IGNORE}
$1
EOF
}

cat <<- EOF >./ungen-proto.sh
#!/usr/bin/env bash

##
# This script is generated by $script_path
# it will remove all the files generated by it with no testing
##

EOF
add_ignore "**/ungen-proto.sh"

for i in `find ${TMP_DIR} -name "*.go"`; do
    real=${i//$TMP_DIR\/}
    add_ignore "$real"
    rm -rf "$ARCADIA_ROOT/$real"
    cp "$i" "$ARCADIA_ROOT/$real"
    cat <<- EOF >>./ungen-proto.sh
rm -rf "$ARCADIA_ROOT/$real"
EOF
done
cat <<- 'EOF' >>./ungen-proto.sh
rm -rf "$0"
EOF

