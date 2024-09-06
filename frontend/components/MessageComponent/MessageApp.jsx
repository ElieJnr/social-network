import { Button } from "@/components/ui/button";
import { Emoji, ReceivedMessage, SendIcon, SendingMessage, ShowEmoji } from "../ui/message";
import { Card, CardContent } from "../ui/card";
import { Input } from "../ui/input";
import { useState } from 'react';
import useStore from "@/app/store/useStore";
import { useWebSocket } from "@/app/actions/message";
import { useParams } from "next/navigation";




export default function GroupChat() {
    const [clicked, setClicked] = useState(false);
    const resocket = useWebSocket('ws://localhost:8080/ws');
    const { id } = useParams();
    const { groupMessage } = useStore();
    const { currentUser, setCurrenUser } = useStore()


    console.log(groupMessage);


    const handleClick = () => {
        resocket.send(JSON.stringify({
            Type: "enterGroupMessage",
            ReceiverId: id
        }));
        setClicked(true);
    };

    const handleBackClick = (event) => {
        event.stopPropagation();
        setClicked(false);
    };

    return (
        <div onClick={handleClick}>
            {clicked ? (
                <div className="fixed bottom-0 right-0 flex items-center justify-center">
                    <Card>
                        <CardContent>
                            <MessageHeader onBackClick={handleBackClick} />
                            <MessageBody message={groupMessage} id={id} currentUser={currentUser} />
                            <MessageForm socket={resocket} id={id} />
                        </CardContent>
                    </Card>
                </div>
            ) : (
                <Click />
            )}
        </div>
    )
}

function MessageHeader({ onBackClick }) {
    return (
        <header className="py-4 px-6 flex items-center justify-between">
            <div>
                <p className="font-medium">Messagerie du groupe</p>
            </div>
            <div>
                <button onClick={onBackClick}>
                    <ArrowDownIcon />
                </button>
            </div>
        </header>
    )
}


function MessageBody({ message, id, currentUser }) {
    return (
        <div className="h-[50vh] flex-1 overflow-y-auto p-6">
            <div className="grid gap-4">
                {message.map(msg => {
                    if (msg.sender_id === currentUser) {
                        return <ReceivedMessage key={msg.id} mess={msg.msg} />;
                    } else {
                        return <SendingMessage key={msg.id} mess={msg.msg} />;
                    }
                })}
            </div>
        </div>
    );
}



function MessageForm({ socket, id }) {
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
                <form onSubmit={(e) => handleSubmit(e, inputValue, setInputValue, socket, id)}>
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

function handleSubmit(e, message, setInputValue, socket, id) {
    e.preventDefault();


    console.log("Message envoyé :", message);
    if (message.trim() && socket?.readyState === WebSocket.OPEN) {
        socket.send(JSON.stringify({
            Type: "newGroupChat",
            Content: message,
            ReceiverId: id
        }));
        setInputValue('');
    } else {
        console.log("socket est ouvert: ", socket?.readyState === WebSocket.OPEN);
    }
}

function Click() {
    return (
        <div className="fixed bottom-0 right-0 flex items-center justify-center">
            <div className="w-full max-w-md flex items-center justify-between px-6 py-4">
                <h1 className="text-2xl font-bold" style={{ color: "black" }}>Messages</h1>
                <div className="flex items-center gap-4">
                    <Button variant="ghost" size="icon">
                        <MailPlusIcon className="w-5 h-5" />
                        <span className="sr-only">New message</span>
                    </Button>
                    <Button variant="ghost" size="icon">
                        <ArrowUpIcon className="w-5 h-5" />
                        <span className="sr-only">Expand/collapse</span>
                    </Button>
                </div>
            </div>
        </div>
    )
}

function ArrowUpIcon(props) {
    return (
        <svg
            {...props}
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            strokeWidth="2"
            strokeLinecap="round"
            strokeLinejoin="round"
        >
            <path d="m5 12 7-7 7 7" />
            <path d="M12 19V5" />
        </svg>
    )
}


function MailPlusIcon(props) {
    return (
        <svg
            {...props}
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            strokeWidth="2"
            strokeLinecap="round"
            strokeLinejoin="round"
        >
            <path d="M22 13V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v12c0 1.1.9 2 2 2h8" />
            <path d="m22 7-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7" />
            <path d="M19 16v6" />
            <path d="M16 19h6" />
        </svg>
    )
}
function ArrowDownIcon() {
    return (
        <div className="flex items-center justify-center">
            <svg
                xmlns="http://www.w3.org/2000/svg"
                width="24"
                height="24"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                strokeWidth="2"
                strokeLinecap="round"
                strokeLinejoin="round"
            >
                <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    d="M19 9l-7 7-7-7"
                />
            </svg>
        </div>
    );
}