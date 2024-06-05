#!/bin/sh
for (( i=1; i <= 10000; i++ ))
do
grpc_cli call --json_input --json_output localhost:9091 github.com.mgrigoriev.chat_monorepo.chatmessages.ChatMessagesService/SaveChatMessage '{"info": {"content": "Hi dude", "recipient_id": "2", "recipient_type": "USER", "user_id": "3", "user_name": "Neo"}}'
done
