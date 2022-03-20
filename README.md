# shorturl

## About
This is something like a URL-shortener, but links aren't really short, but whatever, it works.

## How to use
1. Launch the app
2. Go to localhost:3000/s/{link that needs to get shorter}
3. You will get the "shortened" (not really short now, but whatever) link
4. Open that link, and you will be redirected to the needed page 

Currently, links are stored in a json file in C:\Go\Saved\data.json
The file loads into memory every time you launch the program and links are remembered.
Maybe I'll implement a database, but this works fine for now I guess.
