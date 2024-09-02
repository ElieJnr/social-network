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
            // console.log('WebSocket connection opened');
            const message = { Type: "userSender" };
            socket.current.send(JSON.stringify(message));
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
                } else {
                    console.log("no message between two users");
                }
                fetchNotifs(setMsgNotif, "msg")
            } 
            
            if (message.Type === "messageService") {
                console.log('message: ', event.data)
            } 
            
            if (message.Type === "notifications") {
                console.log("c bon");

                fetchNotifs(setNotifs)
            } 
            
            if (message.Type === "simpleChat") {
                setNewMessage(message.message)
                console.log(message.message);
            }
        });

        socket.current.addEventListener('close', (event) => {

        });

        socket.current.addEventListener('error', (event) => {
            console.error('WebSocket error:', event);
        });

        return () => {
            if (socket.current) {
                socket.current.close();
            }
        };
    }, [url]);

    return socket.current;
};
