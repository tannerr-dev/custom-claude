let prompt = 'Job for caddy.service failed because the control process exited with error code. See "systemctl status caddy.service" and "journalctl -xeu caddy.service" for details. ';
let test = '';
let test2 = {};

import { load } from "https://deno.land/std/dotenv/mod.ts";
import Anthropic from 'npm:@anthropic-ai/sdk';

const env = await load();
// console.log(env)

const anthropic = new Anthropic({
  apiKey: env["CLAUDEKEY"]// defaults to process.env["ANTHROPIC_API_KEY"]
});

console.log('asking claude...');
const msg = await anthropic.messages.create({
  model: "claude-3-7-sonnet-20250219",
  max_tokens: 2024,
  messages: [{ role: "user", content: prompt }],
});
console.log('"' + prompt + '"' + '\n' +"---"+ '\n' + msg.content[0].text +"---")

const response = 
 "\n"
 +prompt
 +"\n"
 +`**${Date.now()}**`
 +"\n"
 +"\n---"
 +"\n"
 +msg.content[0].text
 +"\n"
 +"---\n"
 +"\n";

await Deno.writeFile("msg.md", new TextEncoder().encode(response));

const file = await Deno.open("log.md", { append: true, create: true });
await file.write(new TextEncoder().encode(response));
file.close();
console.log("claude responded...");
