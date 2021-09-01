const TelegramBot = require('node-telegram-bot-api');
const fetch = require('request');

const BOT_TOKEN = '1921970606:AAFvOb2DLn58gQqaBGXy2R4a5PFewMcP5NE'
const API_KEY = '560eb9ae5c3d4fcea61d6c85ece0317a'
const GET_URL = `https://newsapi.org/v2/top-headlines?sources=techcrunch&apiKey=${API_KEY}`;

const bot = new TelegramBot(BOT_TOKEN, {polling: true});
const GREETINGS = ['Hi', 'hey', 'Hello']
const CONVO_PHRASES = ['Sup','Sup?', 'Wassup', 'What\'s up?']


bot.onText(/\/getnews(.*)/, (msg) => {
    let chatId = msg.chat.id;
    fetch(GET_URL, (err, res, body)=> {
        if(!err && res.statusCode == 200) {
            let responseBody = JSON.parse(body)
            for(let i = 0 ; i < 5; i++) {
                let responseString = responseBody.articles[i].url;
                bot.sendMessage(chatId, responseString).then(r => r);;
            }
        }
    })
})

bot.onText(/\/intro(.*)/, (msg) => {
    let chatId = msg.chat.id;
    bot.sendMessage(chatId, "I am Munin, one of Odin's crows. I fly around the world everyday and gather news for my master.").then(r => r);;
})

bot.on('message', (msg)=> {
    let chatId = msg.chat.id;
    if(msg.text ==='Hello there'){
        bot.sendMessage(chatId, "General Kenobi...").then(r => r);;
    }
    if(CONVO_PHRASES.includes(msg.text)){
        bot.sendMessage(chatId, "I'm alright.. how bout you??").then(r => r);
    }
    if(GREETINGS.includes(msg.text)){
        bot.sendMessage(chatId, "Heyy.. How you doin' ;)").then(r => r);;
    }
})

