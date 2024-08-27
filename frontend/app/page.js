"use client";

import { useEffect } from "react";

export default function Page() {
  useEffect(() => {
    const json = {
      EmailOrName: "adiane@",
      Password: "securepassword12",
    };
    const fetchData = async () => {
      const response = await fetch("http://localhost:8080/login", {
        method: "POST",
        credentials: "include",
        body: JSON.stringify(json),
      });

      const data = await response.json();
      console.log(data);
    };
    fetchData();
  }, []);
  return (
    <>
      <button>Post</button>
    </>
  );
}
