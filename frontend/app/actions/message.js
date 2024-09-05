import { useEffect, useRef } from 'react';
import useStore from '../store/useStore';
import { ClickMessageApp } from '@/components/ui/message';
import { fetchNotifs } from './notifications';
import { setMsgNotif } from '@/components/MsgNotig';
import { setNotifs } from '@/components/notifications';
export const useWebSocket = (url) => {
    let socket = useRef(null);
    let { user, setUser } = useStore()
    let { getMessage, setgetMessage } = useStore()
    let { newMessage, setNewMessage } = useStore()


    useEffect(() => {
        socket.current = new WebSocket(url);
        socket.current.addEventListener('open', (event) => {
            if (socket.current.readyState === WebSocket.OPEN) {
                const message = { Type: "userSender" };
                socket.current.send(JSON.stringify(message));
            }
        });


        socket.current.addEventListener('message', (event) => {
            let message = JSON.parse(event.data)
            console.log(message);

            if (message.Type === "sendUser") {
                setUser(message.Users)
            }

            if (message.Type === "clickOnUser") {

                if (message.Message) {
                    setgetMessage(message.Message)
                    console.log("parsing part: ", message);
                }
                fetchNotifs(setMsgNotif, "msg")
            }
            if (message.Type === "messageService") {
                console.log('message: ', event.data)
            }

            if (message.Type === "notifications") {
                fetchNotifs(setNotifs)
            }

            if (message.Type === "sendMessage") {
                setgetMessage(message.message)
            }
        });

        socket.current.addEventListener('close', (event) => {
            console.log('WebSocket closed:', event.code, event.reason);
        });


        socket.current.addEventListener('error', (event) => {
            console.error('WebSocket error:', event);
        });

        return () => {
            return () => {
                if (socket.current && socket.current.readyState === WebSocket.OPEN) {
                    socket.current.close();
                }
            };

        };
    }, [url]);

    return socket.current;
};

export function socketSend(socket, message) {
    if (socket && socket.readyState === WebSocket.OPEN) {
        console.log("the msg is send to your socket");

        socket.send(JSON.stringify(message))
    }
}
