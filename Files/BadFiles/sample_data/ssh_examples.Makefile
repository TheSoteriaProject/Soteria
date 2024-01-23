# NOTE(nic): a small example of an insecure `ssh` invocation.

SSH_OPTS := -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null
SSH_CMD := /usr/bin/ssh
SCP_CMD := /usr/bin/scp
TARGET := target_directory
SOURCE := source_directory
BUILD_CMD := build_command
COPY_FILES_CMD := copy_files_command

.PHONY: build copy

build:
	@echo 'Running remote build:'
	$(SSH_CMD) $(SSH_OPTS) user@example.com "cd $(TARGET); $(BUILD_CMD)"

copy:
	@echo 'Copying files back to local machine:'
	$(SCP_CMD) $(SSH_OPTS) user@example.com:$(TARGET)/$(COPY_FILES_CMD) $(SOURCE)

insecure_with_vars: build copy

insecure_plain:
	@echo 'Plain insecure SSH invocation:'
	ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null user@example.com "cd $(TARGET); $(BUILD_CMD)"
	scp -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null user@example.com:$(TARGET)/$(COPY_FILES_CMD) $(SOURCE)
