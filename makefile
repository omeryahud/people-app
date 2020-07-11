OPERATOR_SDK_VERSION ?= v0.18.0

operator-sdk:
	curl -L https://github.com/operator-framework/operator-sdk/releases/download/${OPERATOR_SDK_VERSION}/operator-sdk-${OPERATOR_SDK_VERSION}-x86_64-linux-gnu -o bin/operator-sdk
	chmod 0755 bin/operator-sdk

.PHONY: opeator-sdk