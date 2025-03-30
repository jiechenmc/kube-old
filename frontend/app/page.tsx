"use client";
import Image from "next/image";
import React, { FC, useState } from "react";
import { ReactSortable } from "react-sortablejs";

interface ItemType {
  id: number;
  name: string;
}
export default function Home() {
  const [state, setState] = useState([
    { Runnable: [{ id: 1, name: "CSE534" }], Running: [], Sleep: [], Done: [] },
  ]);

  return (
    <div className="flex min-h-screen p-8 pb-20 gap-16 sm:p-20 font-[family-name:var(--font-geist-sans)]">

      {Object.keys(state).map(k=>{return <>
        {console.log(k)}
        <h1>{k}</h1>
 
      </>})}
    </div>
  );
}
