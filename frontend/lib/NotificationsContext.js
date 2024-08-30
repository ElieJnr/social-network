

import { createContext, useContext, useState } from 'react';

const NotificationsContext = createContext();

export function useNotifications() {
    return useContext(NotificationsContext);
}

export function NotificationsProvider({ children }) {
    const [notifications, setNotifications] = useState([]);

    return (
        <NotificationsContext.Provider value={{ notifications, setNotifications }}>
            {children}
        </NotificationsContext.Provider>
    );
}
