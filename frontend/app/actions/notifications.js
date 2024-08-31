export async function fetchNotifs(setNotifications) {
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
            console.log(data);
            setNotifications(data)
        }
    } catch (error) {
        console.error('Error fetching posts:', error);
    }
}

