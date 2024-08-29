import React, { useState, useEffect } from "react";
import { Button } from "@/components/ui/button";
import { Avatar, AvatarImage, AvatarFallback } from "@/components/ui/avatar";
import { Input } from "@/components/ui/input";
import EmojiPicker from 'emoji-picker-react';
import useStore from "@/app/store/useStore";
import { useWebSocket } from "@/app/actions/message";


export function ClickMessageApp({ socket }) {
    const [showMessageApp, setShowMessageApp] = useState(false);
    const { message, setMessage } = useStore()

    const [selectedUser, setSelectedUser] = useState({
        Id: "",
        Firstname: "",
        Lastname: ""
    });
    const handleNameClick = () => {
        if (socket && socket.readyState === WebSocket.OPEN) {
            socket.send(JSON.stringify({ Type: "clickOnUser" }));
        }
        setShowMessageApp(true);
    };
    const handleBackClick = () => {
        setShowMessageApp(false)
    }

    return showMessageApp ?
        (<MessageApp backClick={handleBackClick} user={selectedUser} socket={socket} />
        ) : (
            <MessageComponent onNameClick={handleNameClick} setSelectedUser={setSelectedUser} socket={socket} />
        );
}

export function MessageComponent({ onNameClick, setSelectedUser,socket }) {
    const { user, setUser } = useStore()

    const handleClick = (user) => {
        setSelectedUser({
            Id: user.Id,
            Firstname: user.Firstname,
            Lastname: user.Lastname
        });
        if (socket && socket.readyState === WebSocket.OPEN) {
            console.log("selected user", user)
            socket.send(JSON.stringify({ Type: "clickOnUser", ReceiverId: user.Id }));
        }
    };

    const [getMessage, setGetMessage] = useState([])

    return (
        <div className="w-full max-w-md mx-auto bg-background text-foreground rounded-lg shadow-lg">
            <div onClick={onNameClick} className="px-4 py-6">
                <div className="flex items-center justify-between mb-4">
                    <h2 className="text-xl font-bold">Messages</h2>
                </div>
                <div className="space-y-4">
                    {user.map(user => (
                        <div key={user.Id} className="flex items-center justify-between" onClick={() => handleClick(user)}>
                            <UserList id={user.Id} name={capitalize(user.Firstname) + " " + capitalize(user.Lastname)} lastMessage={`Salut ${user.Firstname}`} />
                            <div className="flex items-center space-x-2">
                                <time className="text-sm text-muted-foreground">2:34 PM</time>
                                <div className="w-2 h-2 bg-primary rounded-full" />
                            </div>
                        </div>
                    ))}
                </div>
            </div>
        </div>
    );
}

function capitalize(str) {
    if (!str) return '';
    return str.charAt(0).toUpperCase() + str.slice(1).toLowerCase();
}

function UserList({ id, name, lastMessage }) {
    return (
        <div className="flex items-center space-x-3">
            <Avatar className="bg-primary-foreground text-primary">
                <AvatarImage src="/placeholder-user.jpg" alt="John Doe" />
                <AvatarFallback>JD</AvatarFallback>
            </Avatar>
            <div className="cursor-pointer">
                <div className="font-bold">{name}</div>
                <div className="text-sm text-muted-foreground">{lastMessage}</div>
            </div>
        </div>
    )
}
export function MessageApp({ backClick, user, socket }) {
    const [message, setMessage] = useState('');

    const handleEmojiSelect = (emoji) => {
        setMessage((prevMessage) => prevMessage + emoji.emoji);
    };
    return (
        <div className="flex flex-col h-screen bg-white" style={{ height: "40%" }}>
            <header className="bg-[#f0f4f8] py-4 px-6 flex items-center justify-between">
                <div onClick={backClick} className="flex items-center gap-3">
                    <Button variant="ghost" size="icon">
                        <ArrowLeftIcon className="w-5 h-5" />
                    </Button>
                    <Avatar className="w-8 h-8">
                        <AvatarImage src="/placeholder-user.jpg" alt="User avatar" />
                        <AvatarFallback>MS</AvatarFallback>
                    </Avatar>
                    <div>
                        <p className="font-medium">{capitalize(user.Firstname) + " " + capitalize(user.Lastname)}</p>
                        <p className="text-sm text-muted-foreground">Online</p>
                    </div>
                </div>
            </header>
            <div className="flex-1 overflow-y-auto p-6">
                <div className="grid gap-4">
                    <ReceivedMessage />
                    <SendingMessage />
                </div>
            </div>
            <div className="bg-[#f0f4f8] py-4 px-6" style={{ position: "relative", bottom: 0, left: 0 }}>
                <form
                    className="flex items-center gap-3"
                    onSubmit={(e) => {
                        e.preventDefault();

                        if (message.trim() !== "") {
                            // console.log("Formulaire rempli avec:", message);
                            if (socket && socket.readyState === WebSocket.OPEN) {
                                console.log(user);

                                socket.send(JSON.stringify({ Type: "sendMessage", Content: message, ReceiverId: user.Id }));
                            }

                        } else {
                            console.log("Le formulaire n'est pas rempli");
                        }
                    }}
                >
                    <ShowEmoji onEmojiSelect={handleEmojiSelect} />
                    <Input
                        id="message"
                        placeholder="Type your message..."
                        className="flex-1"
                        autoComplete="off"
                        value={message}
                        onChange={(e) => setMessage(e.target.value)}
                    />
                    <Button type="submit" size="icon">
                        <SendIcon className="w-5 h-5" />
                    </Button>
                </form>

            </div>
        </div>
    );
}


function Emoji({ onEmojiClick }) {
    return (
        <div style={{ position: "absolute", bottom: 100, left: 0 }}>
            <EmojiPicker onEmojiClick={onEmojiClick} />
        </div>
    );
}

function ShowEmoji({ onEmojiSelect }) {
    const [showEmojiPicker, setEmojiPicker] = useState(false);

    const handleClick = () => {
        setEmojiPicker(!showEmojiPicker);
    };

    return (
        <div>
            <Button onClick={handleClick}>
                <SmileIcon className="w-5 h-5" />
            </Button>
            {showEmojiPicker && <Emoji onEmojiClick={onEmojiSelect} />}
        </div>
    );
}
function ReceivedMessage() {
    return (
        <div className="flex justify-end">
            <div className="bg-[#f0f4f8] text-sm rounded-lg px-4 py-3 max-w-[70%]">
                <p>Hey there! How are you doing today?</p>
            </div>
        </div>
    )
}

function SendingMessage() {
    return (
        <div className="flex">
            <div className="bg-[#e6f4ea] text-sm rounded-lg px-4 py-3 max-w-[70%]">
                <p>I'm doing great, thanks for asking! It's good to hear from you.</p>
            </div>
        </div>
    )
}

function SendIcon(props) {
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
            <path d="m22 2-7 20-4-9-9-4Z" />
            <path d="M22 2 11 13" />
        </svg>
    );
}

function ArrowLeftIcon(props) {
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
            <path d="m12 19-7-7 7-7" />
            <path d="M19 12H5" />
        </svg>
    )
}

function SmileIcon(props) {
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
            <circle cx="12" cy="12" r="10" />
            <path d="M8 14s1.5 2 4 2 4-2 4-2" />
            <line x1="9" x2="9.01" y1="9" y2="9" />
            <line x1="15" x2="15.01" y1="9" y2="9" />
        </svg>
    )
}