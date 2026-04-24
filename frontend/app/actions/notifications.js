export async function fetchNotifs(setNotifications, type) {
    try {
        const response = await fetch("http://localhost:8080/notifications", {
            method: "GET",
            credentials: "include",
        });
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        let data = await response.json();

        if (data) {
            if (type) {
                data = data.filter((n) => n.Type == "msg")
            }
            data = data.filter((n) => !n.IsRead)
            console.log(data);
            setNotifications(data)
        }
    } catch (error) {
        console.error('Error fetching notif:', error);
    }
}

export function SendFollow(socket) {

}