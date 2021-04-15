# Webserver
This is just an attempt at making a really simple and really small webserver that runs in a docker container.  Intent is to use it for a Kubernetes demonstration at some point in the future.

`./run.sh` to build the image and launch the container

# Key Points 
- base `golang` image is about 800mb
- reduce the size by compiling our go program using a multi step container build
- need to compile in all dependencies
- need to add certificates

# Inspiration
- https://technology.riotgames.com/news/leveraging-golang-game-development-and-operations
- https://chemidy.medium.com/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324