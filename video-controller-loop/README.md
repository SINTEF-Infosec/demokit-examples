# Video controller loop

This example demonstrates how to chain videos: once the VIDEO_PLAY event is 
received, the two videos (1) and (2) will alternate. The chain is created by 
binding a load video and play video actions to the internal event "MEDIA_ENDED". 