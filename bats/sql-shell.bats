#!/usr/bin/env bats
load $BATS_TEST_DIRNAME/helper/common.bash

setup() {
    setup_common
    dolt sql <<SQL
CREATE TABLE test (
  pk BIGINT NOT NULL COMMENT 'tag:0',
  c1 BIGINT COMMENT 'tag:1',
  c2 BIGINT COMMENT 'tag:2',
  c3 BIGINT COMMENT 'tag:3',
  c4 BIGINT COMMENT 'tag:4',
  c5 BIGINT COMMENT 'tag:5',
  PRIMARY KEY (pk)
);
SQL
}

teardown() {
    teardown_common
}

@test "run a query in sql shell" {
    skiponwindows "Works on Windows command prompt but not the WSL terminal used during bats"
    run bash -c "echo 'select * from test;' | dolt sql"
    [ $status -eq 0 ]
    [[ "$output" =~ "pk" ]] || false
}

@test "bad sql in sql shell should error" {
    run dolt sql <<< "This is bad sql"
    [ $status -eq 1 ]
    run dolt sql <<< "select * from test; This is bad sql; insert into test (pk) values (666); select * from test;"
    [ $status -eq 1 ]
    [[ ! "$output" =~ "666" ]] || false
}

@test "inline query with missing -q flag should error" {
    run dolt sql "SELECT * FROM test;"
    [ $status -eq 1 ]
    [[ "$output" =~ "Invalid Argument:" ]] || false
}