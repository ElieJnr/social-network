"use client";

import SeeGroup from "@/components/seegroup"
import { useParams } from "next/navigation";

export default function SeeGroupPage() {
  const {id} = useParams()
  console.log(id);
  return (
    <div>
      <SeeGroup />
    </div>
  )
}