// Copyright 2020 The Godror Authors
//
// SPDX-License-Identifier: UPL-1.0 OR Apache-2.0

package godror

//go:generate bash -c "echo 3.3.0>odpi-version; set -x; curl -L https://github.com/oracle/odpi/archive/v$(cat odpi-version).tar.gz | tar xzvf - odpi-$(cat odpi-version)/{embed,include,src,CONTRIBUTING.md,LICENSE.md,README.md} && rm -rf odpi && mv odpi-$(cat odpi-version) odpi; rm -f odpi-version; git checkout -- $(git status --porcelain -- odpi/*/*.go | sed -n -e '/ D / { s/^ D //;p;}')"

// Version of this driver
const Version = "v0.15.0"
