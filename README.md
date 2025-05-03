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

