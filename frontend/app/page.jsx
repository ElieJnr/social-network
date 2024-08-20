'use client'

import { useEffect } from "react";

export default function Page() {

    useEffect(() => {
        const json = {
            EmailOrUsername: "adiane@",
            Password: "securepassword123"
        };
        const fetchData = async () => {
            
            const response = await fetch('http://localhost:8080/login', {
                method: 'POST',
                credentials: 'include', // Inclure les cookies dans la requête
                body: JSON.stringify(json)

            });
            
            const data = await response.json();
            console.log(data);
        }
        // const fetchDatas = async () => {
        //     const json = {

        //     }
        //     const response = await fetch('http://localhost:8080/', {
        //         method: 'GET',
        //         credentials: 'include', // Inclure les cookies dans la requête

        //     });

        //     const data = await response.json();
        //     console.log(data);
        // }
        fetchData()
        // fetchDatas()
    }, [])
    return <>
        <button>Post</button>
    </>
}


