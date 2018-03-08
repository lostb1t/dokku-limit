#!/usr/bin/env bats

load test_helper

setup() {
  global_setup
  [[ -f "$DOKKU_ROOT/VHOST" ]] && cp -fp "$DOKKU_ROOT/VHOST" "$DOKKU_ROOT/VHOST.bak"
  [[ -f "$DOKKU_ROOT/HOSTNAME" ]] && cp -fp "$DOKKU_ROOT/HOSTNAME" "$DOKKU_ROOT/HOSTNAME.bak"
  create_app
}

teardown() {
  destroy_app 0 $TEST_APP
  [[ -f "$DOKKU_ROOT/VHOST.bak" ]] && mv "$DOKKU_ROOT/VHOST.bak" "$DOKKU_ROOT/VHOST" && chown dokku:dokku "$DOKKU_ROOT/VHOST"
  [[ -f "$DOKKU_ROOT/HOSTNAME.bak" ]] && mv "$DOKKU_ROOT/HOSTNAME.bak" "$DOKKU_ROOT/HOSTNAME" && chown dokku:dokku "$DOKKU_ROOT/HOSTNAME"
  global_teardown
}

assert_nonssl_domain() {
  local domain=$1
  assert_app_domain "${domain}"
  assert_http_success "http://${domain}"
}

assert_app_domain() {
  local domain=$1
  run /bin/bash -c "dokku domains $TEST_APP 2> /dev/null | grep -xF ${domain}"
  echo "output: "$output
  echo "status: "$status
  assert_output "${domain}"
}

assert_external_port() {
  local CID="$1"; local exit_status="$2"
  local EXTERNAL_PORT_COUNT=$(docker port $CID | wc -l)
  run /bin/bash -c "[[ $EXTERNAL_PORT_COUNT -gt 0 ]]"
  if [[ "$exit_status" == "success" ]]; then
    assert_success
  else
    assert_failure
  fi
}

@test "(proxy) proxy:enable/disable" {
  deploy_app
  assert_nonssl_domain "${TEST_APP}.dokku.me"

  run dokku proxy:disable $TEST_APP
  echo "output: "$output
  echo "status: "$status
  assert_success

  for CID_FILE in $DOKKU_ROOT/$TEST_APP/CONTAINER.web.*; do
    assert_external_port $(< $CID_FILE) failure
  done

  run dokku proxy:enable $TEST_APP
  echo "output: "$output
  echo "status: "$status
  assert_success
  assert_http_success "${TEST_APP}.dokku.me"

  for CID_FILE in $DOKKU_ROOT/$TEST_APP/CONTAINER.web.*; do
    assert_external_port $(< $CID_FILE) failure
  done
}

@test "(proxy) proxy:ports (list/add/remove/clear)" {
  run dokku proxy:ports-add $TEST_APP http:8080:5000 https:8443:5000 http:1234:5001
  echo "output: "$output
  echo "status: "$status
  assert_success

  run /bin/bash -c "dokku --quiet proxy:ports $TEST_APP | xargs"
  echo "output: "$output
  echo "status: "$status
  assert_output "http 1234 5001 http 8080 5000 https 8443 5000"

  run /bin/bash -c "dokku proxy:ports-remove $TEST_APP 8080"
  echo "output: "$output
  echo "status: "$status
  assert_success

  run /bin/bash -c "dokku --quiet proxy:ports $TEST_APP | xargs"
  echo "output: "$output
  echo "status: "$status
  assert_output "http 1234 5001 https 8443 5000"

  run /bin/bash -c "dokku proxy:ports-remove $TEST_APP http:1234:5001"
  echo "output: "$output
  echo "status: "$status
  assert_success

  run /bin/bash -c "dokku --quiet proxy:ports $TEST_APP | xargs"
  echo "output: "$output
  echo "status: "$status
  assert_output "https 8443 5000"

  run /bin/bash -c "dokku proxy:ports-clear $TEST_APP"
  echo "output: "$output
  echo "status: "$status
  assert_success

  run /bin/bash -c "dokku --quiet proxy:ports $TEST_APP | xargs"
  echo "output: "$output
  echo "status: "$status
  assert_output "http 80 5000"
}

@test "(proxy) proxy:ports (post-deploy add)" {
  deploy_app
  run dokku proxy:ports-add $TEST_APP http:8080:5000 http:8081:5000
  echo "output: "$output
  echo "status: "$status
  assert_success

  URLS="$(dokku --quiet urls "$TEST_APP")"
  for URL in $URLS; do
    assert_http_success $URL
  done
  assert_http_success "http://$TEST_APP.dokku.me:8080"
  assert_http_success "http://$TEST_APP.dokku.me:8081"
}
