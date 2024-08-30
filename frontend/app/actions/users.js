import { makeGetFetchUsers, makeGetFetch } from ".";

export async function userConnect() {
    const data = await makeGetFetch("http://localhost:8080/getUser?key=userConnect&id=null");
    return data;
}


export async function GetAllInfoForUserById(userId) {
    const infoUser = await makeGetFetch(`http://localhost:8080/getUser?key=user&id=${userId}`);
    //fetch des posts de l'utilisateur
    return infoUser;
}

export async function allUsers() {
    const data = await makeGetFetchUsers("http://localhost:8080/users");
    return data
}


export function search(users, follows) {
    return users.filter(user => 
        !follows.some(follow => follow.followedUser === user.id)
    );
}



  
export async function searchUsers(users, statut) {
    if (users) {
        const userPromises = users.map(async (user) => {
            const userId = (statut === "userId") ? user.userId : user.followedUser;
            console.log(userId);
            const userInfo = await GetAllInfoForUserById(userId);
            return userInfo;
        });

        const tabResult = await Promise.all(userPromises);
        
        return tabResult;
    }
    return []; 
}

