# NOTE(nic): this is a sample that shows the various and sundry
#  ways one can run `curl --insecure` in a Makefile.
#
#  The comments in this file should *NOT* match.

REMOTE_SERVER := https://remoteserver.com/
INSTALLER_URL := https://installer.com/installer.zip
POST_FILE := ./data.json
DOWNLOADED_FILE := ./installer.zip
CURL := curl
CURL_OPTIONS := -k
CURL_LONG_OPTIONS := --insecure

.PHONY: all post_results_short post_results_long download_installer_short download_installer_long plot_log update_readme

all: post_results_short post_results_long download_installer_short download_installer_long plot_log update_readme

# Using curl -k to post results to a remote server.  This should match
post_results_short:
	@echo "Posting results to remote server (short)..."
	/opt/widgetco/bin/curl -k -X POST -H "Content-Type: application/json" -d @$(POST_FILE) $(REMOTE_SERVER)

# Using curl --insecure to post results to a remote server.  This should match
post_results_long:
	@echo "Posting results to remote server (long)..."
	$(CURL) $(CURL_LONG_OPTIONS) -X POST -H "Content-Type: application/json" -d @$(POST_FILE) $(REMOTE_SERVER)

# Using curl -k to download installer files.  This should match
download_installer_short:
	@echo "Downloading installer files (short)..."
	$(CURL) $(CURL_OPTIONS) -O $(INSTALLER_URL)

# Using curl to download installer files.  This should *NOT* match
download_installer_short:
	@echo "Downloading installer files (short)..."
	$(CURL) -O $(INSTALLER_URL)

# Using curl with plaintext HTTP.  This should match
download_installer_short:
	$(CURL) http://example.com > /dev/null

# Using curl --insecure to download installer files.  This should match
download_installer_long:
	@echo "Downloading installer files (long)..."
	curl --insecure -O $(INSTALLER_URL)
	mv installer.zip $(DOWNLOADED_FILE)

# Using plot_log_semicurl -k to plot the log file.  This should *NOT* match
plot_log:
	@echo "Plotting log file..."
	plot_log_semicurl -k ./log.txt

# Append text to README.md if it does not already exist.  This should *NOT* match
update_readme:
	@echo "Updating README.md..."
	UPDATE_TEXT='We have banned the use of `curl --insecure` in our code'; \
	grep -qsF "$$UPDATE_TEXT" README.md || echo $$UPDATE_TEXT >> README.md

$(CURL):
	@echo "Pretending to install curl..."
	@touch $@

# Calling curl --insecure through plenty of indirection.  This should match
invoke_via_target: $(CURL)
	$^ $(CURL_LONG_OPTIONS) -X DELETE $(INSTALLER_URL)

