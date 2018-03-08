#!/usr/bin/env bats

load test_helper
TEST_PLUGIN_NAME=smoke-test-plugin
TEST_PLUGIN_GIT_REPO=https://github.com/dokku/${TEST_PLUGIN_NAME}.git

setup() {
  global_setup
}

remove_test_plugin() {
  rm -rf $PLUGIN_ENABLED_PATH/$TEST_PLUGIN_NAME $PLUGIN_AVAILABLE_PATH/$TEST_PLUGIN_NAME
}

teardown() {
  remove_test_plugin || true
  global_teardown
}

@test "(plugin) plugin:install, plugin:disable, plugin:update plugin:uninstall" {
  run bash -c "dokku plugin:install $TEST_PLUGIN_GIT_REPO --name $TEST_PLUGIN_NAME"
  echo "output: "$output
  echo "status: "$status
  assert_success

  run bash -c "dokku plugin:update $TEST_PLUGIN_NAME"
  echo "output: "$output
  echo "status: "$status
  assert_success

  run bash -c "dokku plugin | grep enabled | grep $TEST_PLUGIN_NAME"
  echo "output: "$output
  echo "status: "$status
  assert_success

  run bash -c "sudo -E -u nobody dokku plugin:uninstall $TEST_PLUGIN_NAME"
  echo "output: "$output
  echo "status: "$status
  assert_failure

  run bash -c "dokku plugin:disable $TEST_PLUGIN_NAME"
  echo "output: "$output
  echo "status: "$status
  assert_success

  run bash -c "dokku plugin | grep disabled | grep $TEST_PLUGIN_NAME"
  echo "output: "$output
  echo "status: "$status
  assert_success

  run bash -c "dokku plugin:uninstall $TEST_PLUGIN_NAME"
  echo "output: "$output
  echo "status: "$status
  assert_success

  run bash -c "dokku plugin | grep $TEST_PLUGIN_NAME"
  echo "output: "$output
  echo "status: "$status
  assert_failure
}

@test "(plugin) plugin:install plugin:update (with tag)" {
  run bash -c "dokku plugin:install $TEST_PLUGIN_GIT_REPO --committish v0.2.0 --name $TEST_PLUGIN_NAME"
  echo "output: "$output
  echo "status: "$status
  assert_success

  run bash -c "dokku plugin | grep enabled | grep $TEST_PLUGIN_NAME | grep 0.2.0"
  echo "output: "$output
  echo "status: "$status
  assert_success

  run bash -c "dokku plugin:update $TEST_PLUGIN_NAME v0.3.0"
  echo "output: "$output
  echo "status: "$status
  assert_success

  run bash -c "dokku plugin | grep enabled | grep $TEST_PLUGIN_NAME | grep 0.2.0"
  echo "output: "$output
  echo "status: "$status
  assert_failure

  run bash -c "dokku plugin | grep enabled | grep $TEST_PLUGIN_NAME | grep 0.3.0"
  echo "output: "$output
  echo "status: "$status
  assert_success
}

@test "(plugin) plugin:install, plugin:disable, plugin:uninstall as non-root user failure" {
  run bash -c "sudo -E -u nobody dokku plugin:install $TEST_PLUGIN_GIT_REPO"
  echo "output: "$output
  echo "status: "$status
  assert_failure

  run bash -c "dokku plugin:install $TEST_PLUGIN_GIT_REPO"
  echo "output: "$output
  echo "status: "$status
  assert_success

  run bash -c "dokku plugin | grep enabled | grep $TEST_PLUGIN_NAME"
  echo "output: "$output
  echo "status: "$status
  assert_success

  run bash -c "sudo -E -u nobody dokku plugin:disable $TEST_PLUGIN_NAME"
  echo "output: "$output
  echo "status: "$status
  assert_failure
}
