#!/bin/bash
if [ -z ${CI_TAG} ] ; then
    export CUSTOM_VERSION="0-${CI_BRANCH}-${CI_SHORT_SHA}";
else
    export CUSTOM_VERSION=${CI_TAG#v}-${CI_SHORT_SHA};
fi
