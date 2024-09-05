import { Button } from "@/components/ui/button";
import { Emoji, ReceivedMessage, SendIcon, SendingMessage, ShowEmoji } from "../ui/message";
import { Card, CardContent } from "../ui/card";
import { Input } from "../ui/input";
import { useState } from 'react';


export default function GroupChat({ socket }) {
    return (
        <Card>
            <CardContent>
                <MessageHeader groupName={"Test Group"} />
                <MessageBody />
                <MessageForm socket={socket} />
            </CardContent>
        </Card>
    )
}

function MessageHeader({ groupName }) {
    return (
        <header className="py-4 px-6 flex items-center justify-center">
            <div className="flex items-center justify-center">
                <div>
                    <p className="font-medium">{groupName}</p>
                </div>
            </div>
        </header>
    )
}

function MessageBody() {
    return (
        <div className="h-[50vh] flex-1 overflow-y-auto p-6">
            <div className="grid gap-4">
                <ReceivedMessage mess={"je teste juste"} />
                <SendingMessage mess={"je teste aussi"} />
            </div>
        </div>
    )
}


function MessageForm({ socket }) {
    const [inputValue, setInputValue] = useState('');

    const handleEmojiSelect = (emoji) => {
        setInputValue((prevMessage) => prevMessage + emoji.emoji);
    };

    return (
        <div className="bg-[#f0f4f8] py-4 px-6 relative bottom-0 left-0">
            <div className="flex items-center gap-3 justify-center">
                <div variant="ghost" size="icon" type="button">
                    <ShowEmoji onEmojiSelect={handleEmojiSelect} />
                </div>
                <form onSubmit={(e) => handleSubmit(e, inputValue, setInputValue, socket)}>
                    <FormInput inputValue={inputValue} setInputValue={setInputValue} />
                </form>
            </div>
        </div>
    );
}

function FormInput({ inputValue, setInputValue }) {
    return (
        <div className="w-[20vw] flex items-center gap-3">
            <Input
                id="message"
                placeholder="Type your message..."
                className="flex-1"
                autoComplete="off"
                value={inputValue}
                onChange={(e) => setInputValue(e.target.value)}
            />
            <Button type="submit" size="icon">
                <SendIcon className="w-5 h-5" />
            </Button>
        </div>
    );
}

function handleSubmit(e, message, setInputValue, socket) {
    e.preventDefault();
    console.log("Message envoyé :", message);
    if (message.trim() && socket?.readyState === WebSocket.OPEN) {
        socket.send(JSON.stringify({
            Type: "newGroupChat",
            Content: message,
            ReceiverId: "1e8a778e-9965-42e3-a2c0-e4263c287d5e"
        }));
        setInputValue('');
    } else {
        console.log("socket est ouvert: ", socket?.readyState === WebSocket.OPEN);
    }
}
