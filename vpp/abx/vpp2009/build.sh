#!/bin/bash

# SPDX-License-Identifier: Apache-2.0

# Copyright 2021 PANTHEON.tech
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -e -o pipefail

if [ $# -lt 1 ]
then
    echo "usage: $0 /path/to/vpp/workspace"
    exit 1
fi

rm -rf build
mkdir -p build
pushd build
cmake -GNinja -DVPP_WORKSPACE=$1 ..
ninja
popd # build
