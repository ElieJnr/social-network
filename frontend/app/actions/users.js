import { makeGetFetch } from ".";

export async function userConnect() {
    const data = await makeGetFetch("http://localhost:8080/getUser?key=userConnect&id=null");
    return data;
}
  
  
  export async function GetAllInfoForUserById(userId){
    const infoUser = await makeGetFetch(`http://localhost:8080/getUser?key=user&id=${userId}`);
    //fetch des posts de l'utilisateur
    return infoUser;
  }