// Copyright 2019 Dolthub, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package commands contains the command functions executed based on the dolt subcommand specified on the command line.
//
// The goal of the code within this package and sub packages is to be a thin client, handling user input and displaying
// formatted output to the user.  The hope is that someone could take the library code, and implement the
// same functionality contained within the command line with very little work beyond providing a new user interface.
package commands
