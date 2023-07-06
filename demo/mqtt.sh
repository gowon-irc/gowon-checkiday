#!/usr/bin/env bash

BROKER_HOST="${BROKER_HOST:-localhost}"
BROKER_PORT="${BROKER_PORT:-1883}"

ACTION="${1}"

days_mqtt_msg() {
	cat <<-EOF
	{"module":"gowon","msg":".days","nick":"tester","dest":"#gowon","command":"days","args":""}
	EOF
}

mdays_mqtt_msg() {
	cat <<-EOF
	{"module":"gowon","msg":".mdays","nick":"tester","dest":"#gowon","command":"mdays","args":""}
	EOF
}

pub() {
    days_mqtt_msg | mosquitto_pub -h "${BROKER_HOST}" -p "${BROKER_PORT}" -t "/gowon/input" -s
    mdays_mqtt_msg | mosquitto_pub -h "${BROKER_HOST}" -p "${BROKER_PORT}" -t "/gowon/input" -s
}

sub() {
    mosquitto_sub -h "${BROKER_HOST}" -p "${BROKER_PORT}" -t "/gowon/output" | jq -r '.msg'
}

case "${ACTION}" in
    pub)
        pub
        ;;
    sub)
        sub
        ;;
    *)
        echo "First argument must be either pub or sub" >&2
        ;;
esac
