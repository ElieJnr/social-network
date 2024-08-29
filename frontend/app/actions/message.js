import { useEffect, useRef } from 'react';
import useStore from '../store/useStore';

export const useWebSocket = (url) => {
    const socket = useRef(null);
    const { user, setUser } = useStore()

    useEffect(() => {
        socket.current = new WebSocket(url);

        socket.current.addEventListener('open', (event) => {
            console.log('WebSocket connection opened');
            const message = { Type: "userSender" };
            socket.current.send(JSON.stringify(message));
        });

        socket.current.addEventListener('message', (event) => {
            let message = JSON.parse(event.data)
            if (message.Type === "sendUser") {
                setUser(message.Users)
            } else if (message.Type === "clickOnUser") {
                console.log('message: ',JSON.parse(event.data))
            } else if (message.Type === "messageService") {
                console.log('message: ', event.data)
            } else if ("simpleChat") {
                console.log(event.data);
            }
        });

        socket.current.addEventListener('close', (event) => {
            console.log('WebSocket connection closed');
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
