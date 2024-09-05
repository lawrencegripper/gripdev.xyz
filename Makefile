# Requires hugo and node

serve:
	hugo serve

build:
	hugo

publish: build 
	npx wrangler pages deploy