# NOTE(nic): this is a sample that shows the various and sundry
#  ways one can run `wget --no-check-certificate` in a Makefile.
#
#  The comments in this file should *NOT* match.

REMOTE_SERVER := https://remoteserver.com/
INSTALLER_URL := https://installer.com/installer.zip
POST_FILE := ./data.json
DOWNLOADED_FILE := ./installer.zip
WGET := wget
WGET_OPTIONS := --no-check-certificate

.PHONY: all post_results download_installer plot_log update_readme

all: post_results download_installer plot_log update_readme

# Using wget --no-check-certificate to post results to a remote server.  This should match
post_results:
	@echo "Posting results to remote server..."
	$(WGET) $(WGET_OPTIONS) --post-file=$(POST_FILE) --header="Content-Type: application/json" -O- $(REMOTE_SERVER)

# Using wget --no-check-certificate to download installer files.  This should *NOT* match
download_installer:
	@echo "Downloading installer files..."
	$(WGET) $(INSTALLER_URL)
	mv installer.zip $(DOWNLOADED_FILE)

# Using wget --no-check-certificate to download installer files.  This should match
download_installer:
	@echo "Downloading installer files..."
	$(WGET) --no-check-certificate $(INSTALLER_URL)
	mv installer.zip $(DOWNLOADED_FILE)

# Plaintext HTTP fetch.  This should match
check_exists:
	$(WGET) --spider http://example.com

# Append text to README.md if it does not already exist.  This should *NOT* match
update_readme:
	@echo "Updating README.md..."
	UPDATE_TEXT='We have banned the use of `wget --no-check-certificate` in our code'; \
	grep -qsF "$$UPDATE_TEXT" README.md || echo $$UPDATE_TEXT >> README.md

$(WGET):
	@echo "Pretending to install wget..."
	@touch $@

# Calling wget --no-check-certificate through plenty of indirection.  This should match
invoke_via_target: $(WGET)
	$^ $(WGET_OPTIONS) --delete-after $(INSTALLER_URL)
