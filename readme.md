## Setup

Go to `transip.nl/cp/` log in and go to the API tab.
create a key-pair there

save the output there in `key.private` 

## Build the image
If you cant to build this image yourself, simply run 
`docker build -t dyndns .` 

## Run the image 
If you have build it yourself use the tag you used in previous step

`docker run -it -v ./key.private:/key.private -e DOMAIN=example.com -e USERNAME=example_account -e SUBDOMAIN=dyndns xantios/dyndns`

this wil setup dynamic DNS to your WAN ip on `dyndns.example.com` assuming you own that domain in the account named `example_account` 