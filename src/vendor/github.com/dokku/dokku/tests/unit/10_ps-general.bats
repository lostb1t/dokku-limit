#!/usr/bin/env bats

load test_helper

setup() {
  global_setup
  create_app
}

teardown() {
  destroy_app
  global_teardown
}

@test "(ps:scale) procfile commands extraction" {
  source "$PLUGIN_CORE_AVAILABLE_PATH/common/functions"
  source "$PLUGIN_CORE_AVAILABLE_PATH/ps/functions"
  cat <<EOF > "$DOKKU_ROOT/$TEST_APP/DOKKU_PROCFILE"
web: node web.js
worker: node worker.js
EOF
  run get_cmd_from_procfile "$TEST_APP" web
  echo "output: "$output
  echo "status: "$status
  assert_output "node web.js"

  run get_cmd_from_procfile "$TEST_APP" worker
  echo "output: "$output
  echo "status: "$status
  assert_output "node worker.js"
}

@test "(ps:restart-policy) default policy" {
  run bash -c "dokku --quiet ps:restart-policy $TEST_APP"
  echo "output: "$output
  echo "status: "$status
  assert_output "on-failure:10"
}


@test "(ps:restart-policy) ps:set-restart-policy, ps:restart-policy" {
  for policy in no unless-stopped always on-failure on-failure:20; do
    run bash -c "dokku ps:set-restart-policy $TEST_APP $policy"
    echo "output: "$output
    echo "status: "$status
    assert_success

    run bash -c "dokku --quiet ps:restart-policy $TEST_APP"
    echo "output: "$output
    echo "status: "$status
    assert_output "$policy"
  done
}

@test "(ps:restart-policy) deployed policy" {
  test_restart_policy="on-failure:20"
  run bash -c "dokku ps:set-restart-policy $TEST_APP $test_restart_policy"
  echo "output: "$output
  echo "status: "$status
  assert_success

  run bash -c "dokku --quiet ps:restart-policy $TEST_APP"
  echo "output: "$output
  echo "status: "$status
  assert_output "$test_restart_policy"

  deploy_app dockerfile

  CID=$(< $DOKKU_ROOT/$TEST_APP/CONTAINER.web.1)
  run bash -c "docker inspect -f '{{ .HostConfig.RestartPolicy.Name }}:{{ .HostConfig.RestartPolicy.MaximumRetryCount }}' $CID"
  echo "output: "$output
  echo "status: "$status
  assert_output "$test_restart_policy"
}
