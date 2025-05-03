# How to use this program
1. edit the chat.ts prompt value
2. run command
```
deno --allow-net --allow-read --allow-env --allow-write chat.ts
```
3. run save command
```
./save.sh <filename>
```
- saves to archive folder, appends to existing file, else creates new

## Notes
- log.md saves every reponse, msg save last
    - rm log.md and msg.md periodically

---

I am currently writing a go webserver to handle the interaction and the file system writing.

I want to have the option to view the chat log, last chat and save either one to memory.
Maybe even save the api conversations? keep the json data? 

ok, save the json data as the log and last message, then have the option to save thee
to be used as context or as markdown file

