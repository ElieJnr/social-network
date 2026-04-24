"use client";

import SeeGroup from "@/components/seegroup";
import { useParams } from "next/navigation";

export default function SeeGroupPage() {
  const {id} = useParams()
  return (
    <div>
      <SeeGroup id={id}/>
    </div>
  )
}
