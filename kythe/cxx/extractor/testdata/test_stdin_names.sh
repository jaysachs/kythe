#!/bin/bash -e
# Copyright 2015 The Kythe Authors. All rights reserved.
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
set -o pipefail
TEST_NAME="test_stdin_names"
. ./kythe/cxx/extractor/testdata/test_common.sh
. ./kythe/cxx/extractor/testdata/skip_functions.sh
rm -rf -- "${OUT_DIR}"
echo '#define STDIN_OK 1\n' | KYTHE_INDEX_PACK=1 \
    KYTHE_OUTPUT_DIRECTORY="${OUT_DIR}" \
    KYTHE_VNAMES="${BASE_DIR}/stdin.vnames" "${EXTRACTOR}" -x c -

# Before we chdir, we need the concrete path of the tool we're going to use.
readonly REAL_INDEXPACK="$PWD/$INDEXPACK"
pushd "${OUT_DIR}"
OUT_INDEX=$("${REAL_INDEXPACK}" --verbose --from_archive "${OUT_DIR}" | \
    awk '-F\t' '/Unpacked compilation/ { print $3 }')
popd

# Make sure that the indexer can handle <stdin:> paths.
"${INDEXER}" --ignore_unimplemented=true "${OUT_DIR}/${OUT_INDEX}" | \
  "${VERIFIER}" --nofile_vnames "${BASE_DIR}/test_stdin_names_verify.cc"
