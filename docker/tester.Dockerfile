# SPDX-License-Identifier: Apache-2.0

# Copyright 2021 PANTHEON.tech
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

FROM ubuntu:20.04

ARG DEBIAN_FRONTEND=noninteractive

RUN set -ex; \
    apt-get update && apt-get install -y --no-install-recommends \
		bridge-utils \
		ca-certificates \
		curl \
		ethtool \
		expect-dev \
		iperf \
		iproute2 \
		iptables \
		iputils-ping \
		netcat-openbsd \
		net-tools \
		tcl8.6 \
		tcpdump \
		wget \
    # TODO remove
    python3-scapy \
	; \
	rm -rf /var/lib/apt/lists/*
