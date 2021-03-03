# toast-notification

bit overcomplicated cos copy-pasted from my other more complex project.

also ignore the docker-compose stuff.

Im not using https with my android app due to android 7.0 boringssl issue with elliptic curvey things.

Also if you want to try and get docker-compose working,

there is a chicken and egg scenario where,

you need nginx container running for certbot to succeed.
you need certbot container to have run for nginx container to succeed.

I resolved this by commenting out the https portion of nginx conf, letting certbot succeed (important to remove the --staging when its working), and then adding it back in.
