// Copyright 2020 PingCAP-QE libs Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package di

import (
    "database/sql"
    "time"
)

func ProcessCreatedDI(issueDB, diDB *sql.DB, repo, sig string, startTime, endTime time.Time) error {

}

func ProcessClosedDI(issueDB, diDB *sql.DB, repo, sig string, startTime, endTime time.Time) error {

}

func ProcessDI(issueDB, diDB *sql.DB, repo, sig string, time time.Time) error {

}

func ProcessCreatedDIs(issueDB, diDB *sql.DB, repo, sig string, startTime, endTime time.Time, frequency time.Duration) error {

}

func ProcessClosedDIs(issueDB, diDB *sql.DB, repo, sig string, startTime, endTime time.Time, frequency time.Duration) error {

}

func ProcessDIs(issueDB, diDB *sql.DB, repo, sig string, startTime, endTime time.Time, frequency time.Duration) error {

}
