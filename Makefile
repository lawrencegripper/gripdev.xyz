# Requires hugo and node

serve:
	hugo serve --buildDrafts

build:
	# Clean out directory
	rm -rf ./public 
	# Rebuild it
	hugo

publish: build 
	npx wrangler pages deploy