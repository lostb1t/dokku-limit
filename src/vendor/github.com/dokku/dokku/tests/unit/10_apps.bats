#!/usr/bin/env bats

load test_helper

setup () {
  global_setup
}

teardown () {
  global_teardown
}

@test "(apps) apps:create" {
  run dokku apps:create $TEST_APP
  echo "output: "$output
  echo "status: "$status
  assert_success
  run bash -c "dokku apps:list | grep $TEST_APP"
  echo "output: "$output
  echo "status: "$status
  assert_output $TEST_APP
  destroy_app

  run dokku apps:create 1994testapp
  echo "output: "$output
  echo "status: "$status
  assert_success

  run dokku apps:create testapp:latest
  echo "output: "$output
  echo "status: "$status
  assert_failure

  run dokku apps:create testApp:latest
  echo "output: "$output
  echo "status: "$status
  assert_failure

  run dokku apps:create TestApp
  echo "output: "$output
  echo "status: "$status
  assert_failure

  run bash -c "dokku --app $TEST_APP apps:create"
  echo "output: "$output
  echo "status: "$status
  assert_success
  run bash -c "dokku apps:list | grep $TEST_APP"
  echo "output: "$output
  echo "status: "$status
  assert_output $TEST_APP

  destroy_app
}

@test "(apps) apps:destroy" {
  create_app
  run bash -c "dokku --force apps:destroy $TEST_APP"
  echo "output: "$output
  echo "status: "$status
  assert_success

  create_app
  run bash -c "dokku --force --app $TEST_APP apps:destroy"
  echo "output: "$output
  echo "status: "$status
  assert_success
}

@test "(apps) apps:rename" {
  deploy_app
  run bash -c "dokku apps:rename $TEST_APP great-test-name"
  echo "output: "$output
  echo "status: "$status
  assert_success
  run bash -c "dokku apps:list | grep $TEST_APP"
  echo "output: "$output
  echo "status: "$status
  assert_output ""
  run bash -c "curl --silent --write-out '%{http_code}\n' `dokku url great-test-name` | grep 404"
  echo "output: "$output
  echo "status: "$status
  assert_output ""
  run bash -c "dokku --force apps:destroy great-test-name"
  echo "output: "$output
  echo "status: "$status
  assert_success

  run dokku apps:create $TEST_APP
  echo "output: "$output
  echo "status: "$status
  assert_success
  run bash -c "dokku apps:rename $TEST_APP great-test-name"
  echo "output: "$output
  echo "status: "$status
  assert_success
  run bash -c "dokku --force apps:destroy great-test-name"
  echo "output: "$output
  echo "status: "$status
  assert_success
}

@test "(apps) apps:clone" {
  deploy_app
  run bash -c "dokku apps:clone $TEST_APP great-test-name"
  echo "output: "$output
  echo "status: "$status
  assert_success
  run bash -c "dokku apps:list | grep $TEST_APP"
  echo "output: "$output
  echo "status: "$status
  assert_success
  run bash -c "curl --silent --write-out '%{http_code}\n' `dokku url great-test-name` | grep 404"
  echo "output: "$output
  echo "status: "$status
  assert_output ""
  run bash -c "dokku --force apps:destroy great-test-name"
  echo "output: "$output
  echo "status: "$status
  assert_success

  run bash -c "dokku apps:clone --skip-deploy $TEST_APP great-test-name"
  echo "output: "$output
  echo "status: "$status
  assert_success
  run bash -c "curl --silent --write-out '%{http_code}\n' `dokku url great-test-name` | grep 404"
  echo "output: "$output
  echo "status: "$status
  assert_failure
  run bash -c "dokku --force apps:destroy great-test-name"
  echo "output: "$output
  echo "status: "$status
  assert_success
}
